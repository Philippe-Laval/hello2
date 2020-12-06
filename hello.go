package main

import (
	"fmt"

	"github.com/Philippe-Laval/stringutil"
)

func main() {
	str := stringutil.Reverse("Hello Marina and Philippe.")
	fmt.Println(str)
	fmt.Println(stringutil.Reverse(str))
}
