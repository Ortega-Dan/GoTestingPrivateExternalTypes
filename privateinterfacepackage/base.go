package privateinterfacepackage

type privateInterface interface {
	DoSomething() string
}

type Base struct {
	privateInterface
}

func (b *Base) DoSomething() string {
	return "base is doing"
}
