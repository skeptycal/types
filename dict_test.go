package types

import (
	"reflect"
	"testing"
)

var (
	mapTests = []struct {
		name      string
		protected bool
		want      *dict
		wantErr   bool
	}{
		// no error (using make(AnyMap))
		{"protected", true, &dict{name: "protected", protected: true, m: make(AnyMap)}, false},
		{"unprotected", false, &dict{name: "unprotected", protected: false, m: make(AnyMap)}, false},

		// empty AnyMap the same as make(AnyMap)
		{"empty map true", true, &dict{name: "empty map true", protected: true, m: AnyMap{}}, false},
		{"empty map false", false, &dict{name: "empty map false", protected: false, m: AnyMap{}}, false},

		// name is incorrect
		{"name error", false, &dict{name: "fake", protected: false, m: make(AnyMap)}, true},
		{"protected error", true, &dict{name: "protected error", protected: false, m: make(AnyMap)}, true},

		// protected is incorrect
		{"protected", false, &dict{name: "protected", protected: true, m: make(AnyMap)}, true},
		{"unprotected", true, &dict{name: "unprotected", protected: false, m: make(AnyMap)}, true},

		// no map in struct leaves out zero value of initialized map
		{"no map true", true, &dict{name: "no map true", protected: true}, true},
		{"no map false", false, &dict{name: "no map false", protected: false}, true},

		// nil map is not the same as empty, initialized AnyMap
		{"nil map true", true, &dict{name: "nil map true", protected: true, m: nil}, true},
		{"nil map false", false, &dict{name: "nil map false", protected: false, m: nil}, true},
	}

	sampleAnyMap = AnyMap{
		"int 42":    42,
		"string 42": "42",
		"byte 42":   []byte("42"),
		"rune 42":   []rune("42"),
		"AnyMap":    AnyMap{},
		"nil":       nil,
	}

	sampleKeys = []Any{
		"int 42",
		"string 42",
		"byte 42",
		"rune 42",
		"AnyMap",
		"nil",
	}

	sampleValues = []Any{
		42,
		"42",
		[]byte("42"),
		[]rune("42"),
		AnyMap{},
		nil,
	}
)

func TestNewDict(t *testing.T) {
	for _, tt := range mapTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDict(tt.name, tt.protected); !reflect.DeepEqual(got, tt.want) != tt.wantErr {
				t.Errorf("NewDict(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_dict_Keys(t *testing.T) {
	name := "sampleAnyMap"
	d := &dict{m: sampleAnyMap}
	want := sampleKeys
	for k, v := range sampleAnyMap {
		err := d.Set(k, v)
		if err != nil {
			t.Errorf("Set(%v, %v) = %v", k, v, err)
		}
	}

	t.Run(name, func(t *testing.T) {

		got := d.Keys()
		gotLen := len(got)
		wantLen := len(want)
		if gotLen != wantLen {
			t.Errorf("dict.Keys().Len() = %v, want %v", gotLen, wantLen)
		}

		for k := range sampleAnyMap {
			if !Contains(k, want) {
				t.Errorf("dict.Keys() does not contain %v ", k)
				break
			}
			if Count(k, want) != Count(k, got) {
				t.Errorf("dict.Keys() incorrect count: %v ", k)
				break
			}
		}

	})
}

func Test_dict_Values(t *testing.T) {
	name := "sampleAnyMap"
	d := &dict{m: sampleAnyMap}
	want := sampleValues
	for k, v := range sampleAnyMap {
		err := d.Set(k, v)
		if err != nil {
			t.Errorf("Set(%v, %v) = %v", k, v, err)
		}
	}

	t.Run(name, func(t *testing.T) {

		got := d.Values()
		gotLen := len(got)
		wantLen := len(want)
		if gotLen != wantLen {
			t.Errorf("dict.Values().Len() = %v, want %v", gotLen, wantLen)
		}

		for _, v := range sampleAnyMap {
			if !Contains(v, want) {
				t.Errorf("dict.Values() does not contain %v ", v)
				break
			}
			if Count(v, want) != Count(v, got) {
				t.Errorf("dict.Values() incorrect count: %v ", v)
				break
			}
		}

	})
}

func Test_dict_Less(t *testing.T) {
	type fields struct {
		name      string
		protected bool
		m         AnyMap
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dict{
				name:      tt.fields.name,
				protected: tt.fields.protected,
				m:         tt.fields.m,
			}
			if got := d.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("dict.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
