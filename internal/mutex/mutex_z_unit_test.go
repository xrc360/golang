package mutex_test

import (
	"testing"
	"time"

	"github.com/xrc360/golang/container/garray"
	"github.com/xrc360/golang/internal/mutex"
	"github.com/xrc360/golang/test/gtest"
)

func TestMutexIsSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		lock := mutex.New()
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(false)
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(false, false)
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(true, false)
		t.Assert(lock.IsSafe(), true)

		lock = mutex.New(true, true)
		t.Assert(lock.IsSafe(), true)

		lock = mutex.New(true)
		t.Assert(lock.IsSafe(), true)
	})
}

func TestSafeMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		safeLock := mutex.New(true)
		array := garray.New(true)

		go func() {
			safeLock.Lock()
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			safeLock.Unlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			safeLock.Lock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			safeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 4)
	})
}

func TestUnsafeMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			unsafeLock = mutex.New()
			array      = garray.New(true)
		)

		go func() {
			unsafeLock.Lock()
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			unsafeLock.Unlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			unsafeLock.Lock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			unsafeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 4)
	})
}
