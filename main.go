package main

import (
	"dancho/testing/anotherpackage"
	"dancho/testing/privateinterfacepackage"
)

func main() {

	b := privateinterfacepackage.Base{}

	println(anotherpackage.OperatingBase(&b))
}
