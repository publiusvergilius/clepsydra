package main

import (
	// "clepsydra/db"
	"clepsydra/db"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestConversionSliceToJson(t *testing.T) {

	emptyQuarta := []db.Quartum{}

	var quartum db.Quartum
	quartum.SetTitulum("Test")
	quartum.SetDiesId(1)
	quartum.SetHora(time.Now())
	quartum.SetPars(1)
	mockedQuarta := []db.Quartum{quartum}

	type User struct {
		Nome  string
		Idade int
	}

	type Cases struct {
		Name     string
		Source   any
		Expected string
	}

	fmt.Println(mockedQuarta)

	cases := []Cases{
		{
			Name: "test user struct slice to json",
			Source: []User{
				{
					Nome:  "Vinícius",
					Idade: 20,
				},
			},
			Expected: `[{"Nome":"Vinícius","Idade":20}]`,
		},
		{
			Name:     "test empty Quartum struct slice to json",
			Source:   emptyQuarta,
			Expected: `[]`,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			got, err := json.Marshal(test.Source)
			if err != nil {
				t.Errorf("was not told to error: %q ", err)
			}

			if string(got) != strings.TrimSpace(test.Expected) {
				t.Errorf("got %q, want %q", got, test.Expected)

			}

		})
	}

}

func TestToString(t *testing.T) {
	emptyQuarta := []db.Quartum{}

	var quartum db.Quartum
	quartum.SetTitulum("Test1")
	quartum.SetDiesId(1)
	quartum.SetHora(time.Now())
	quartum.SetPars(1)
	mockedQuarta := []db.Quartum{quartum}

	type Cases[T db.Entity] struct {
		Name     string
		Source   []db.Quartum
		Expected string
	}

	fmt.Println(mockedQuarta)

	cases := []Cases[db.Entity]{
		{
			Name:     "test empty Quartum struct slice to json",
			Source:   emptyQuarta,
			Expected: `[]`,
		},
		{
			Name:     "test mocked Quartum to json",
			Source:   mockedQuarta,
			Expected: fmt.Sprintf(`[{"id":0,"titulum":"Test1","pars":1,"hora":"%s","dies_id":1}]`, quartum.GetHora()),
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			for _, q := range test.Source {
				var listJson []string
				str, err := q.ToString()
				if err != nil {
					t.Errorf("was not told to error: %q ", err)
				}

				listJson = append(listJson, str)
				got := Stringfy(listJson)
				if err != nil {
					t.Errorf("was not told to error: %q ", err)
				}

				if string(got) != strings.TrimSpace(test.Expected) {
					t.Errorf("got %q, want %q", got, test.Expected)

				}
			}

		})
	}
}
