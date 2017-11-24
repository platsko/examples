package types

import (
	"reflect"
	"testing"
)

func TestBR(t *testing.T) {
	tests := []struct {
		name     string
		want     *brLine
		wantJSON []byte
	}{
		{
			name:     "TestBRLine",
			want:     &brLine{Type: "br"},
			wantJSON: []byte(`{"type":"br"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			br := Br()
			if !reflect.DeepEqual(br, test.want) {
				t.Errorf("BR(): %v, want: %v", br, test.want)
			}

			json, err := test.want.MarshalJSON()
			if err != nil {
				t.Errorf("MarshalJSON() error: %v", err)
				return
			}
			if !reflect.DeepEqual(json, test.wantJSON) {
				t.Errorf("MarshalJSON() got: %s, want: %s", json, test.wantJSON)
			}
		})
	}
}

func TestBarCode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *barCode
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := BarCode(test.args.data); !reflect.DeepEqual(got, test.want) {
				t.Errorf("BarCode() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPairLine(t *testing.T) {
	tests := []struct {
		name string
		want *pairLine
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := PairLine(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("PairLine() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	type args struct {
		s *Slip
	}
	tests := []struct {
		name string
		i    Print
		args args
		want *Print
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.i.AddSlip(test.args.s); !reflect.DeepEqual(got, test.want) {
				t.Errorf("AddSlip() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestQrCode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *qrCode
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := QrCode(test.args.data); !reflect.DeepEqual(got, test.want) {
				t.Errorf("QrCode() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestRowLine(t *testing.T) {
	tests := []struct {
		name string
		want *rowLine
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := RowLine(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("RowLine() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSlip(t *testing.T) {
	type args struct {
		line SlipLiner
	}
	tests := []struct {
		name string
		sb   Slip
		args args
		want *Slip
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.sb.AddLine(test.args.line); !reflect.DeepEqual(got, test.want) {
				t.Errorf("AddLine() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestTextLine(t *testing.T) {
	tests := []struct {
		name string
		want *textLine
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := TextLine(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("TextLine() = %v, want %v", got, test.want)
			}
		})
	}
}
