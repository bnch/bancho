package main

import (
	"fmt"
	"git.zxq.co/ripple/go-bancho/web"
)

func main() {
	fmt.Println("== Welcome to go-bancho. ==")
	web.Start(":3000", "10443")
}
