package g

import (
	"github.com/xrcn/cg/internal/utils"
)

// SetDebug enables/disables the GoXrc internal logging manually.
// Note that this function is not concurrent safe, be aware of the DATA RACE,
// which means you should call this function in your boot but not the runtime.
func SetDebug(enabled bool) {
	utils.SetDebugEnabled(enabled)
}
