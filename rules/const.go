package rules

import "regexp"

var (
	// see: https://www.regexpal.com/94925
	iSO8601DateRegex = regexp.MustCompile(`^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])(?:T|\\s)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])?(Z)?$`)
	// see: https://stackoverflow.com/a/43931246/1011803
	millisecondRegex = regexp.MustCompile(`\d{4}-[01]\d-[0-3]\dT[0-2]\d:[0-5]\d:[0-5]\d\.\d+`)
	// see: https://www.speedguide.net/port.php?port=65353
	MinPort = 1
	MaxPort = 65353
)
