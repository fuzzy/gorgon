package utils

import "os"

// Exists checks if a file or directory exists and returns true if it does, false otherwise.
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
