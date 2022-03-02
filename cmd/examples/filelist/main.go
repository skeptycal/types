package main

import (
	"fmt"

	"github.com/skeptycal/types"
)

func main() {
	list := types.GetFileList("")
	for _, fi := range list {
		fmt.Printf("%v\n", fi.Name())
	}
}
