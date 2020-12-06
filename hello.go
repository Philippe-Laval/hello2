package main

import (
	"fmt"

	"github.com/Philippe-Laval/stringutil"
)

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/sqlite

func main() {
	str := stringutil.Reverse("Hello Marina and Philippe.")
	fmt.Println(str)
	fmt.Println(stringutil.Reverse(str))
}
