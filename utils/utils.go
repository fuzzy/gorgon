package utils

import "os"

func Exists(fname string) bool {
	_, err := os.Stat(fname)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
