package gcache

import (
	"github.com/xrcn/cg/os/gtime"
)

// IsExpired checks whether `item` is expired.
func (item *adapterMemoryItem) IsExpired() bool {
	// Note that it should use greater than or equal judgement here
	// imagining that the cache time is only 1 millisecond.

	return item.e < gtime.TimestampMilli()
}
