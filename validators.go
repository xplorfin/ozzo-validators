package ozzo_validators

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/xplorfin/ozzo-validators/rules"
)

// string rules
var (
	// wether or not a string has spaces
	IsSpaceless = validation.NewStringRuleWithError(rules.IsSpaceless, rules.ErrIsSpaceless)
	// wether or not a path exists on the disk
	IsValidPath = validation.NewStringRuleWithError(rules.IsValidPath, rules.ErrIsPath)
	// wether or not a date is a valid iso8601 date
	IsValidIso8601Date = validation.NewStringRuleWithError(rules.IsISO8601Date, rules.ErrIsIso8601Date)
	// wether a date is a valid js millisecond date
	IsJavascriptMillisecondDate = validation.NewStringRuleWithError(rules.IsJavascriptMillisecondDate, rules.ErrIsMillisecondDate)
)

// int rules
var (
	// wether a port is in the valid range
	IsValidPort = rules.NewIntRuleWithError(rules.IsValidPort, rules.ErrIsInvalidPort)
	// wether or not a port is taken
	IsPortAvailable = rules.NewIntRuleWithError(rules.IsAvailablePort, rules.ErrPortUnavailable)
)
