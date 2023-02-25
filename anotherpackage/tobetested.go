package anotherpackage

// instead of using:
//
// import "dancho/testing/privateinterfacepackage"
//
// func OperatingBase(base *privateinterfacepackage.Base) string {
// 	return base.DoSomething()
// }

// the local method is changed to use a local interface and so satisfy both the code and the test

func OperatingBase(base myOwnInterface) string {
	return base.DoSomething()
}
