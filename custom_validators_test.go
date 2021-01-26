package ozzo_validators

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xplorfin/netutils/testutils"
	"github.com/xplorfin/ozzo-validators/rules"
	"reflect"
	"testing"
)

const (
	Iso8601String = "2018-01-01T06:43:14Z"
	// millisecond iso string format
	MilisecondIsoString = "2019-04-01T00:00:00.000Z"
)

func TestAll(t *testing.T) {
	tests := []struct {
		tag            string
		rule           validation.Rule
		valid, invalid interface{}
		err            string
	}{
		// string tests
		{"IsSpaceless", IsSpaceless, "jake_was_here", "jake was here", rules.ErrIsSpaceless.Message()},
		{"IsNumeric", IsNumeric, fmt.Sprintf("%f", gofakeit.Float64()), "not a number", rules.ErrIsNumeric.Message()},
		{"IsInt", IsInt, fmt.Sprintf("%d", gofakeit.Int64()), fmt.Sprintf("%f", gofakeit.Float64()), rules.ErrIsInt.Message()},
		{"IsIso8601", IsValidIso8601Date, Iso8601String, gofakeit.Date().String(), rules.ErrIsIso8601Date.Message()},
		{"IsMillisecondDate", IsJavascriptMillisecondDate, MilisecondIsoString, gofakeit.Date().String(), rules.ErrIsMillisecondDate.Message()},
		// file tests
		{"IsPath", IsValidPath, filet.TmpFile(t, "", gofakeit.Sentence(50)).Name(), gofakeit.CreditCard().Number, rules.ErrIsPath.Message()},
		{"IsDir", IsDir, filet.TmpDir(t, ""), filet.TmpFile(t, "", gofakeit.Sentence(10)).Name(), rules.ErrIsDir.Message()},
		{"IsFile", IsFile, filet.TmpFile(t, "", gofakeit.Sentence(10)).Name(), filet.TmpDir(t, ""), rules.ErrIsFile.Message()},
		// net tests
		{"IsValidPort", IsValidPort, rules.MinPort, rules.MaxPort + 1, rules.ErrIsInvalidPort.Message()},
		{"IsAvailablePort", IsPortAvailable, testutils.GetFreePort(t), testutils.GetUnfreePort(t), rules.ErrPortUnavailable.Message()},
	}

	for _, test := range tests {
		if reflect.TypeOf(test.valid).Kind() == reflect.String && reflect.TypeOf(test.invalid).Kind() == reflect.String {
			err := test.rule.Validate("")
			assert.Nil(t, err, test.tag)
		}
		err := test.rule.Validate(test.valid)
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(&test.valid)
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(test.invalid)
		assertError(t, test.err, err, test.tag)
		err = test.rule.Validate(&test.invalid)
		assertError(t, test.err, err, test.tag)
	}
}

func assertError(t *testing.T, expected string, err error, tag string) {
	if expected == "" {
		assert.Nil(t, err, tag)
	} else if assert.NotNil(t, err, tag) {
		assert.Equal(t, expected, err.Error(), tag)
	}
}
