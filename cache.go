package main

import (
	"os"
	"time"
)

type GorgonCache struct {
	Dir         string
	LastSync    time.Time
	LastSyncDir string
	TempDir     string
}

func NewGorgonCache(dir string) *GorgonCache {
	retv := &GorgonCache{
		Dir:         dir,
		LastSyncDir: "lastsync",
		TempDir:     "temp",
	}

	retv.initCheck()

	return retv
}

func (gc *GorgonCache) initCheck() {
	for _, dir := range []string{gc.Dir, gc.LastSyncDir, gc.TempDir} {
		if !gc.exists(dir) {
			os.MkdirAll(dir, 0755)
		}
	}
}

func (gc *GorgonCache) exists(fname string) bool {
	_, err := os.Stat(fname)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
