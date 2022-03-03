package main

import (
	"testing"
)

func TestMorseConversion(t *testing.T) {
	expected := ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."
	result := ConvertTo("hello world")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
	expected = "-- ..- .... .- -- -- . -.. / ... .- .. -.. / --. .-.. - . -.- .. -."
	result = ConvertTo("Muhammed Said Gültekin")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
	expected = ". .-. -.. . -- / .- -.- -.-- .-.. -.. --.."
	result = ConvertTo("Erdem Akyıldız")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
	expected = "- . ... -"
	result = ConvertTo("test")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
	expected = "----- ----. ---.. --... -.... ..... ....- ...-- ..--- .----"
	result = ConvertTo("0987654321")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
	expected = ""
	result = ConvertTo("*^%")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
	expected = ""
	result = ConvertTo("?(/")
	if result != expected {
		t.Errorf("we got %s result should be %s", result, expected)
	}
}


