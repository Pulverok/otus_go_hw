package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	revString := stringutil.Reverse("Hello, OTUS!")

	fmt.Println(revString)
}
