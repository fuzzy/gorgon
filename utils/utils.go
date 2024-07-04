package main

import "os"

func exists(fname string) bool {
	_, err := os.Stat(fname)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
