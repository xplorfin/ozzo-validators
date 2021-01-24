package rules

import (
	"github.com/asaskevich/govalidator"
	"github.com/brianvoe/gofakeit/v5"
	"math"
	"strconv"
	"testing"
)

// note: this is for special cases (e.g. form fuzzing) where the custom_validators_test
// in the parent directory is not appropriate. Most methods _can_ and *should* be tested
// exlcusively in the parent dir

func TestIsValidPort(t *testing.T) {
	// validate all valid ports come back correct
	for validPortInt := MinPort; validPortInt <= MaxPort; validPortInt++ {
		validPort := strconv.Itoa(validPortInt)
		isValid := IsValidPort(int64(validPortInt))
		isParity := govalidator.IsPort(validPort)
		if !isValid || !isParity {
			t.Errorf("expected port %d to be valid and match govalidator. Got %t from ozzo and %t from isvalidator", validPortInt, isValid, isParity)
		}
	}
	// validate results for invalid are correct
	for i := 0; i < 5000; i++ {
		invalidPortInt := gofakeit.Number(MaxPort+1, math.MaxInt64)
		if gofakeit.Bool() {
			invalidPortInt = gofakeit.Number(math.MinInt32, MinPort-1)
		}
		invalidPort := strconv.Itoa(invalidPortInt)
		isValid := IsValidPort(int64(invalidPortInt))
		isParity := govalidator.IsPort(invalidPort)
		if isValid || isParity {
			t.Errorf("expected port %d to be valid and match govalidator. Got %t from ozzo and %t from isvalidator", i, isValid, isParity)
		}
	}
}
