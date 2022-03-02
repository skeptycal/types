package types

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type (
	gocode struct {
		name     string
		filepath string
		numLines int
		lines    []string
	}

	arg struct {
		v reflect.Value
		k reflect.Kind
	}

	Arg interface{}

	Function interface {
		Name() string
		Func() Any
		Args() []Arg
	}

	GoCode interface {
		Functions() []Function
	}
)

func GetFileList(pattern string) (list []os.FileInfo, err error) {
	if pattern == "" {
		pattern = "*"
	}

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !strings.HasSuffix(info.Name(), "*.go") {
			list = append(list, info)
		}

		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})

	return
}
