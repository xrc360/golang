package rwmutex_test

import (
	"testing"
	"time"

	"github.com/xrc360/golang/container/garray"
	"github.com/xrc360/golang/internal/rwmutex"
	"github.com/xrc360/golang/test/gtest"
)

func TestRWMutexIsSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		lock := rwmutex.New()
		t.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(false)
		t.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(false, false)
		t.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(true, false)
		t.Assert(lock.IsSafe(), true)

		lock = rwmutex.New(true, true)
		t.Assert(lock.IsSafe(), true)

		lock = rwmutex.New(true)
		t.Assert(lock.IsSafe(), true)
	})
}

func TestSafeRWMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			localSafeLock = rwmutex.New(true)
			array         = garray.New(true)
		)

		go func() {
			localSafeLock.Lock()
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			localSafeLock.Unlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			localSafeLock.Lock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			localSafeLock.Unlock()
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

func TestSafeReaderRWMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			localSafeLock = rwmutex.New(true)
			array         = garray.New(true)
		)
		go func() {
			localSafeLock.RLock()
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			localSafeLock.RUnlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			localSafeLock.RLock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			localSafeLock.RUnlock()
		}()
		go func() {
			time.Sleep(500 * time.Millisecond)
			localSafeLock.Lock()
			array.Append(1)
			localSafeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 4)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 6)
	})
}

func TestUnsafeRWMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			localUnsafeLock = rwmutex.New()
			array           = garray.New(true)
		)
		go func() {
			localUnsafeLock.Lock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			localUnsafeLock.Unlock()
		}()
		go func() {
			time.Sleep(500 * time.Millisecond)
			localUnsafeLock.Lock()
			array.Append(1)
			time.Sleep(500 * time.Millisecond)
			array.Append(1)
			localUnsafeLock.Unlock()
		}()
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 4)
	})
}
