// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package empty_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gcron/empty"
)

type TestInt int

type TestString string

type TestPerson interface {
	Say() string
}

func TestIsEmpty(t *testing.T) {

	tmpT1 := "0"
	tmpT2 := func() {}
	tmpT2 = nil
	tmpT3 := make(chan int, 0)
	var (
		tmpT4 TestPerson  = nil
		tmpT5 *TestPerson = nil
		tmpT7 TestInt     = 0
		tmpT8 TestString  = ""
	)
	tmpF1 := "1"
	tmpF2 := func(a string) string { return "1" }
	tmpF3 := make(chan int, 1)
	tmpF3 <- 1
	var (
		tmpF5 TestInt    = 1
		tmpF6 TestString = "1"
	)

	// true
	assert.Equal(t, empty.IsEmpty(nil), true)
	assert.Equal(t, empty.IsEmpty(0), true)
	assert.Equal(t, empty.IsEmpty(gconv.Int(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Int8(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Int16(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Int32(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Int64(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Uint64(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Uint(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Uint16(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Uint32(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Uint64(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Float32(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(gconv.Float64(tmpT1)), true)
	assert.Equal(t, empty.IsEmpty(false), true)
	assert.Equal(t, empty.IsEmpty([]byte("")), true)
	assert.Equal(t, empty.IsEmpty(""), true)
	assert.Equal(t, empty.IsEmpty(tmpT2), true)
	assert.Equal(t, empty.IsEmpty(tmpT3), true)
	assert.Equal(t, empty.IsEmpty(tmpT3), true)
	assert.Equal(t, empty.IsEmpty(tmpT4), true)
	assert.Equal(t, empty.IsEmpty(tmpT5), true)
	assert.Equal(t, empty.IsEmpty(tmpT7), true)
	assert.Equal(t, empty.IsEmpty(tmpT8), true)

	// false
	assert.Equal(t, empty.IsEmpty(gconv.Int(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Int8(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Int16(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Int32(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Int64(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Uint(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Uint8(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Uint16(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Uint32(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Uint64(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Float32(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(gconv.Float64(tmpF1)), false)
	assert.Equal(t, empty.IsEmpty(true), false)
	assert.Equal(t, empty.IsEmpty(tmpT1), false)
	assert.Equal(t, empty.IsEmpty([]byte("1")), false)
	assert.Equal(t, empty.IsEmpty(tmpF2), false)
	assert.Equal(t, empty.IsEmpty(tmpF3), false)
	assert.Equal(t, empty.IsEmpty(tmpF5), false)
	assert.Equal(t, empty.IsEmpty(tmpF6), false)

}

func TestIsNil(t *testing.T) {

	assert.Equal(t, empty.IsNil(nil), true)

	var ii int
	assert.Equal(t, empty.IsNil(ii), false)

	var i *int
	assert.Equal(t, empty.IsNil(i), true)

	assert.Equal(t, empty.IsNil(&i), false)
	assert.Equal(t, empty.IsNil(&i, true), true)

}
