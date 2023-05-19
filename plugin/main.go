package main

import (
	"flag"
	"plugin"
)

// go run main.go --file_name=pluginA/pluginA.so
// go run main.go --file_name=pluginB/pluginB.so
func main() {

	fileName := flag.String("file_name", "plugin/pluginA/pluginA.so", "plug_file")
	flag.Parse()

	p, err := plugin.Open(*fileName)
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int64) = 7
	f.(func())() // prints "Hello, number 7"
}
