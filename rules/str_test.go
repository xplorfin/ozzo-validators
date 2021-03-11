package rules

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
)

// we fuzz a number of these using a random seed from gofakeit
// these are NOT meant to be deterministic

// helper to generate int strings
func generateIntStrings() (intStrings []string) {
	return []string{
		fmt.Sprintf("%d", gofakeit.Int8()),
		fmt.Sprintf("%d", gofakeit.Int16()),
		fmt.Sprintf("%d", gofakeit.Int32()),
		fmt.Sprintf("%d", gofakeit.Int64()),
		// uints,
		fmt.Sprintf("%d", gofakeit.Uint8()),
		fmt.Sprintf("%d", gofakeit.Uint16()),
		fmt.Sprintf("%d", gofakeit.Uint32()),
		fmt.Sprintf("%d", gofakeit.Uint64()),
	}
}

func generateFloatStrings() []string {
	return []string{
		fmt.Sprintf("%f", gofakeit.Float32()),
		fmt.Sprintf("%f", gofakeit.Float64()),
	}
}

func TestIsSpaceless(t *testing.T) {
	gofakeit.Seed(int64(rand.Int()))
	for i := 0; i < 50; i++ {
		word := gofakeit.Word()
		if !IsSpaceless(word) {
			t.Errorf("expected string %s to be spaceless", word)
		}

		sentence := gofakeit.Sentence(50)
		if IsSpaceless(sentence) {
			t.Errorf("expected string '%s' to not be spaceless", sentence)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	gofakeit.Seed(int64(rand.Int()))
	for i := 0; i < 50; i++ {
		validNumbers := append(generateIntStrings(), generateFloatStrings()...)

		for _, number := range validNumbers {
			if !IsNumeric(number) {
				t.Errorf("expected number %s to be numeric", number)
			}
		}

		invalidNumbers := strings.Split(gofakeit.LoremIpsumSentence(len(validNumbers)), " ")
		for _, invalidNumber := range invalidNumbers {
			if IsNumeric(invalidNumber) {
				t.Errorf("expected invalid number %s to be invalid", invalidNumber)
			}
		}
	}
}

func TestIsInt(t *testing.T) {
	gofakeit.Seed(int64(rand.Int()))
	for i := 0; i < 50; i++ {
		for _, validInt := range generateIntStrings() {
			if !IsInt(validInt) {
				t.Errorf("expected valid int '%s' to be valid", validInt)
			}
		}
		for _, validFloat := range generateFloatStrings() {
			if IsInt(validFloat) {
				t.Errorf("expected float '%s' to be invalid int", validFloat)
			}
		}
	}
}
