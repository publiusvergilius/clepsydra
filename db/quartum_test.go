package db

import "testing"

func TestCreateDate(t *testing.T) {

	q := Quartum{}
	test := "16:40:00"
	time, err := q.CreateHourFromString(test)
	if err != nil {
		t.Errorf("was not told to error: %s", err.Error())
	}
	q.SetHora(time)

	got := q.GetHora()

	if q.GetHora() != "16:40:00" {
		t.Errorf("want %q, got %q", test, got)
	}

}
