package main

import (
	"testing"
)

func TestShouldConvert50PercentageInBoldAndNotJumpWords(t *testing.T) {
	percentage := 50
	jumps := 0
	text := "maturidade maturidade"

	expectedText := "**matur**idade **matur**idade"

	convertResult := ConvertTextToBold(percentage, text, jumps)

	if convertResult != expectedText {
		t.Errorf("Expected '%s' but was '%s'", expectedText, convertResult)
	}
}

func TestShouldConvert70PercentageInBoldAndNotJumpWords(t *testing.T) {
	percentage := 70
	jumps := 0
	text := "maturidade maturidade"

	expectedText := "**maturid**ade **maturid**ade"

	convertResult := ConvertTextToBold(percentage, text, jumps)

	if convertResult != expectedText {
		t.Errorf("Expected '%s' but was '%s'", expectedText, convertResult)
	}
}

func TestShouldConvert70PercentageInBoldAndJumpOneWord(t *testing.T) {
	percentage := 70
	jumps := 1
	text := "maturidade maturidade"

	expectedText := "maturidade **maturid**ade"

	convertResult := ConvertTextToBold(percentage, text, jumps)

	if convertResult != expectedText {
		t.Errorf("Expected '%s' but was '%s'", expectedText, convertResult)
	}
}

func TestShouldConvert70PercentageInBoldAndJumpTwoWords(t *testing.T) {
	percentage := 70
	jumps := 2
	text := "maturidade maturidade maturidade"

	expectedText := "maturidade maturidade **maturid**ade"

	convertResult := ConvertTextToBold(percentage, text, jumps)

	if convertResult != expectedText {
		t.Errorf("Expected '%s' but was '%s'", expectedText, convertResult)
	}
}
