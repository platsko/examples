package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPhone_Parse(t *testing.T) {
	phone := Phone("+12223334455")

	tests := []struct {
		name    string
		phone   *Phone
		vals    []string
		want    *Phone
		wantErr bool
	}{
		{
			"TestPhone_Parse_Success",
			new(Phone),
			[]string{"+12223334455"},
			&phone,
			false,
		},
		{
			"TestPhone_Parse_Error",
			new(Phone),
			[]string{"", "12223334455", "012223334455"},
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, val := range test.vals {
				got, err := test.phone.Parse(val)
				if (err != nil) != test.wantErr {
					t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
					return
				}
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("Parse() got = %v, want %v", got, test.want)
				}
			}
		})
	}
}

func TestPhone_String(t *testing.T) {
	number := "+12223334455"
	phone := Phone(number)

	tests := []struct {
		name  string
		phone *Phone
		want  string
	}{
		{
			"TestPhone_String",
			&phone,
			number,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.phone.String(); got != test.want {
				t.Errorf("String() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPhone_Value(t *testing.T) {
	value := "12223334455"
	number := fmt.Sprintf("+%s", value)
	phone := Phone(number)

	tests := []struct {
		name  string
		phone *Phone
		want  string
	}{
		{
			"TestPhone_Value",
			&phone,
			value,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.phone.Value(); got != test.want {
				t.Errorf("Value() = %v, want %v", got, test.want)
			}
		})
	}
}
