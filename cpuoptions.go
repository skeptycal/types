package types

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/skeptycal/errorlogger"
	"golang.org/x/sys/cpu"
)

var log = errorlogger.Log

type (
	field struct {
		Name string
		Type reflect.Type
	}
	fields map[string]field
)

var ValueOf = reflect.ValueOf

func Kind(a Any) reflect.Kind {
	return ValueOf(a).Kind()
}

func Fields(a Any) map[string]reflect.Value {

	v := ValueOf(a)

	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()

	// structName := t.Name()

	n := t.NumField()

	m := make(map[string]reflect.Value, n)

	// name1 := v.Type().Field(0).Name

	for i := 0; i < n; i++ {
		f := t.Field(i)
		// value := fmt.Sprintf("%v", f)
		value := ValueOf(f)
		m[f.Name] = value
	}

	return m
}

func cpuFields() map[string]reflect.Value {

	m := make(map[string]reflect.Value)

	list := []Any{
		cpu.ARM,
		cpu.ARM64,
		cpu.MIPS64X,
		cpu.PPC64,
		cpu.S390X,
		cpu.X86,
	}

	for _, item := range list {
		for k, v := range Fields(item) {
			for _, vv := range Fields(v) {
				if vv.Kind() == reflect.Bool {
					// if vv.Bool() {
					key := fmt.Sprintf("%v_%v", item, k)
					m[key] = v
					// }
				}
			}
		}
	}
	return m
}

func CPUOptions() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for k, v := range cpuFields() {
		sb.WriteString(fmt.Sprintf("%v = %v\n", k, v))
	}

	return sb.String()
}

func HasAVX2() bool {
	return cpu.X86.HasAVX2
}
