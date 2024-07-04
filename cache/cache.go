// Description: This file contains the GorgonCache struct and its methods. It is
// profoundly uninteresting at this point in time. Cache operations will be added
// later on.
package cache

import (
	"os"
	"time"

	"github.com/fuzzy/gorgon/utils"
)

// GorgonCache is a struct that holds the cache directory and the last sync time.
type GorgonCache struct {
	Dir      string
	LastSync time.Time
	TempDir  string
}

// NewGorgonCache creates a new GorgonCache struct.
func NewGorgonCache(dir string) *GorgonCache {
	retv := &GorgonCache{
		Dir:     dir,
		TempDir: "temp",
	}

	retv.initCheck()

	return retv
}

func (gc *GorgonCache) initCheck() {
	for _, dir := range []string{gc.Dir, gc.TempDir} {
		if !utils.Exists(dir) {
			os.MkdirAll(dir, 0755)
		}
	}
}
