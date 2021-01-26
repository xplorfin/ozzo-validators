package rules

import "os"

// IsFile checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// check if a path exists/is a dir
func IsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// check wether or not a path is valid
func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
