package main

import (
	"fmt"
	dependency "github.com/britishrum/test-renovated-repo-dep"
)

func main() {
	fmt.Println(dependency.GetVersion())
}
