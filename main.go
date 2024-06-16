package main

import (
	"fmt"
	"github.com/jaloren/go-enumexp/enumcar"
)

func main() {
	c := enumcar.New(enumcar.Honda)
	fmt.Println(c)
}
