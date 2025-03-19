package db

import (
	"fmt"
	"testing"
	"time"
)

func TestDies(t *testing.T) {

	newDies := Dies{id: 1}
	newDies.SetDate(time.Now())

	want := fmt.Sprintf(`{"id":1,"date":"%s"}`, newDies.GetDate())

	got, err := newDies.ToString()

	if err != nil {
		t.Errorf("was not told to error: %q", err)
	}
	if want != got {
		t.Errorf("wanted %q, got %q", want, got)
	}
}
