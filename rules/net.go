package rules

import "github.com/xplorfin/netutils/testutils"

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

// check wether or not a port is valid as per rfc 6056
func IsValidPort(port int64) bool {
	if port >= int64(MinPort) && port <= int64(MaxPort) {
		return true
	}
	return false
}
