// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcron

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gcron/gtimer"
	"github.com/zhwei820/gcron/gtype"
	"github.com/zhwei820/log"
	"go.uber.org/zap"
)

// JobFunc is the timing called job function in cron.
type JobFunc = gtimer.JobFunc

// Entry is timing task entry.
type Entry struct {
	cron       *Cron         // Cron object belonged to.
	timerEntry *gtimer.Entry // Associated timer Entry.
	schedule   *cronSchedule // Timed schedule object.
	jobName    string        // Callback function name(address info).
	times      *gtype.Int    // Running times limit.
	infinite   *gtype.Bool   // No times limit.
	Name       string        // Entry name.
	Job        JobFunc       `json:"-"` // Callback function.
	Time       time.Time     // Registered time.
}

type doAddEntryInput struct {
	Name        string          // Name names this entry for manual control.
	Job         JobFunc         // Job is the callback function for timed task execution.
	Ctx         context.Context // The context for the job.
	Times       int             // Times specifies the running limit times for the entry.
	Pattern     string          // Pattern is the crontab style string for scheduler.
	IsSingleton bool            // Singleton specifies whether timed task executing in singleton mode.
	Infinite    bool            // Infinite specifies whether this entry is running with no times limit.
}

// doAddEntry creates and returns a new Entry object.
func (c *Cron) doAddEntry(in doAddEntryInput) (*Entry, error) {
	if in.Name != "" {
		if c.Search(in.Name) != nil {
			return nil, fmt.Errorf(`cron job "%s" already exists`, in.Name)
		}
	}
	schedule, err := newSchedule(in.Pattern)
	if err != nil {
		return nil, err
	}
	// No limit for `times`, for timer checking scheduling every second.
	entry := &Entry{
		cron:     c,
		schedule: schedule,
		jobName:  runtime.FuncForPC(reflect.ValueOf(in.Job).Pointer()).Name(),
		times:    gtype.NewInt(in.Times),
		infinite: gtype.NewBool(in.Infinite),
		Job:      in.Job,
		Time:     time.Now(),
	}
	if in.Name != "" {
		entry.Name = in.Name
	} else {
		entry.Name = "cron-" + gconv.String(c.idGen.Add(1))
	}
	// When you add a scheduled task, you cannot allow it to run.
	// It cannot start running when added to timer.
	// It should start running after the entry is added to the Cron entries map, to avoid the task
	// from running during adding where the entries do not have the entry information, which might cause panic.
	entry.timerEntry = gtimer.AddEntry(
		in.Ctx,
		time.Second,
		entry.checkAndRun,
		in.IsSingleton,
		-1,
		gtimer.StatusStopped,
	)
	c.entries.Set(entry.Name, entry)
	entry.timerEntry.Start()
	return entry, nil
}

// IsSingleton return whether this entry is a singleton timed task.
func (entry *Entry) IsSingleton() bool {
	return entry.timerEntry.IsSingleton()
}

// SetSingleton sets the entry running in singleton mode.
func (entry *Entry) SetSingleton(enabled bool) {
	entry.timerEntry.SetSingleton(enabled)
}

// SetTimes sets the times which the entry can run.
func (entry *Entry) SetTimes(times int) {
	entry.times.Set(times)
	entry.infinite.Set(false)
}

// Status returns the status of entry.
func (entry *Entry) Status() int {
	return entry.timerEntry.Status()
}

// SetStatus sets the status of the entry.
func (entry *Entry) SetStatus(status int) int {
	return entry.timerEntry.SetStatus(status)
}

// Start starts running the entry.
func (entry *Entry) Start() {
	entry.timerEntry.Start()
}

// Stop stops running the entry.
func (entry *Entry) Stop() {
	entry.timerEntry.Stop()
}

// Close stops and removes the entry from cron.
func (entry *Entry) Close() {
	entry.cron.entries.Remove(entry.Name)
	entry.timerEntry.Close()
}

// checkAndRun is the core timing task check logic.
// The running times limits feature is implemented by gcron.Entry and cannot be implemented by gtimer.Entry.
// gcron.Entry relies on gtimer to implement a scheduled task check for gcron.Entry per second.
func (entry *Entry) checkAndRun(ctx context.Context) {
	currentTime := time.Now()
	if !entry.schedule.checkMeetAndUpdateLastSeconds(ctx, currentTime) {
		return
	}
	var err error
	switch entry.cron.status.Val() {
	case StatusStopped:
		log.WarnZ(ctx, "job stopped", zap.String("name", entry.getJobNameWithPattern()))
		return

	case StatusClosed:
		log.WarnZ(ctx, "job removed", zap.String("name", entry.getJobNameWithPattern()))
		entry.Close()

	case StatusReady, StatusRunning:
		defer func() {
			if exception := recover(); exception != nil {
				log.ErrorZ(ctx, "job failed", zap.String("name", entry.getJobNameWithPattern()))
			} else if err == nil {
				log.ErrorZ(ctx, "job end", zap.String("name", entry.getJobNameWithPattern()))
			}

			if entry.timerEntry.Status() == StatusClosed {
				entry.Close()
			}
		}()

		// Running times check.
		if !entry.infinite.Val() {
			times := entry.times.Add(-1)
			if times <= 0 {
				if entry.timerEntry.SetStatus(StatusClosed) == StatusClosed || times < 0 {
					return
				}
			}
		}
		if entry.cron.etcdclient != nil {
			m, errEtcd := entry.cron.etcdclient.NewMutex(fmt.Sprintf("etcd_gcron/%s/%d", entry.jobName, currentTime.Unix()))
			if errEtcd != nil {
				log.ErrorZ(ctx, "fail to create etcd mutex for job", zap.Reflect("entry.jobName", entry.jobName))
				err = errEtcd
				return
			}
			lockCtx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			errEtcd = m.Lock(lockCtx)
			if errEtcd == context.DeadlineExceeded {
				err = errEtcd
				return
			} else if errEtcd != nil {
				log.ErrorZ(ctx, "fail to lock mutex", zap.String("m.Key()", m.Key()))
				err = errEtcd
				return
			}
		}

		log.DebugZ(ctx, `cron job starts`, zap.Any("name", entry.getJobNameWithPattern()))

		entry.Job(ctx)
	}
}

func (entry *Entry) getJobNameWithPattern() string {
	return fmt.Sprintf(`%s(%s)`, entry.jobName, entry.schedule.pattern)
}
