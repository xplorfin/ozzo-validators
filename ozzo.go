package ozzo_validators

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/xplorfin/ozzo-validators/rules"
)

var (
	IsSpaceless                 = validation.NewStringRuleWithError(rules.IsSpaceless, rules.ErrIsSpaceless)
	IsValidPath                 = validation.NewStringRuleWithError(rules.IsValidPath, rules.ErrIsPath)
	IsValidIso8601Date          = validation.NewStringRuleWithError(rules.IsISO8601Date, rules.ErrIsIso8601Date)
	IsJavascriptMillisecondDate = validation.NewStringRuleWithError(rules.IsJavascriptMillisecondDate, rules.ErrIsMillisecondDate)
	IsValidPort                 = rules.NewIntRuleWithError(rules.IsValidPort, rules.ErrIsInvalidPort)
	IsPortAvailable             = rules.NewIntRuleWithError(rules.IsAvailablePort, rules.ErrPortUnavailable)
)
