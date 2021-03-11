package ozzo_validators

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/xplorfin/ozzo-validators/rules"
)

// str rules

// wether or not a string has spaces
var (
	// wether or a not a string contains spaces
	IsSpaceless = validation.NewStringRuleWithError(rules.IsSpaceless, rules.ErrIsSpaceless)
	// wether or not a string is numeric
	IsNumeric = validation.NewStringRuleWithError(rules.IsNumeric, rules.ErrIsNumeric)
	// wether or not a string is an int
	IsInt = validation.NewStringRuleWithError(rules.IsInt, rules.ErrIsInt)
	// wether or not a date is a valid iso8601 date
	IsValidIso8601Date = validation.NewStringRuleWithError(rules.IsISO8601Date, rules.ErrIsIso8601Date)
	// wether a date is a valid js millisecond date
	IsJavascriptMillisecondDate = validation.NewStringRuleWithError(rules.IsJavascriptMillisecondDate, rules.ErrIsMillisecondDate)
)

// path rules
var (
	// wether or not a path exists on the disk
	IsValidPath = validation.NewStringRuleWithError(rules.IsValidPath, rules.ErrIsPath)
	// wether or not a path is a valid directory
	IsDir = validation.NewStringRuleWithError(rules.IsDir, rules.ErrIsDir)
	// wether or not a  path is a file
	IsFile = validation.NewStringRuleWithError(rules.IsFile, rules.ErrIsFile)
)

// int rules
var (
	// wether a port is in the valid range
	IsValidPort = rules.NewIntRuleWithError(rules.IsValidPort, rules.ErrIsInvalidPort)
	// wether or not a port is taken
	IsPortAvailable = rules.NewIntRuleWithError(rules.IsAvailablePort, rules.ErrPortUnavailable)
)
