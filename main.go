package main

import (
	"fmt"

	"github.com/rockisch/appigo/driver"
)

func main() {
	caps := map[string]string{}
	caps["one"] = "ONE"
	caps["two"] = "TWO"

	fmt.Println("vim-go")
	driver := driver.CreateDriver("https://0.0.0.0", caps)
	fmt.Print(driver)
}
