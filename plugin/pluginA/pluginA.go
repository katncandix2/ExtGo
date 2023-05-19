package main

import "fmt"

var V int64

// go build -buildmode=plugin
func F() {
	fmt.Println("pluginA---->", V)
}
