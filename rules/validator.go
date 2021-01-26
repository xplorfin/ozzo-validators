package rules

import (
	"github.com/xplorfin/netutils/testutils"
	"os"
)

// check wether or not a path is valid
func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func IsISO8601Date(date string) bool {
	return iSO8601DateRegex.MatchString(date)
}

func IsJavascriptMillisecondDate(date string) bool {
	return millisecondRegex.MatchString(date)
}

// check wether or not a port is valid as per rfc 6056
func IsValidPort(port int64) bool {
	if port >= int64(MinPort) && port <= int64(MaxPort) {
		return true
	}
	return false
}

// check wether or not anything is running on a given port
func IsAvailablePort(rawPort int64) bool {
	if !IsValidPort(rawPort) {
		return false
	}
	if testutils.PortIsAvailable(int(rawPort)) {
		return true
	}
	return false
}
