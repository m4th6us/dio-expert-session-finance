package util

import "testing"

func TestStringToTime(t *testing.T) {

	var convertedTime = StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que o ano seja 2019")
	}

	if convertedTime.Month() != 2 {
		t.Errorf("Espera que o mês seja 2")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Espera que a hora seja 10")
	}

}
