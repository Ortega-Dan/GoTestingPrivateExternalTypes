package anotherpackage

import (
	"testing"
)

type baseMock struct {
	whatToReturn string
}

func (b *baseMock) DoSomething() string {
	return b.whatToReturn
}

func TestSomething(t *testing.T) {

	mock := baseMock{whatToReturn: "here Im mocking"}

	if OperatingBase(&mock) != mock.whatToReturn {
		t.Error("failed return")
	} else {
		println("successful test")
	}

}
