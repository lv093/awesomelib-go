package factory_method

import "testing"


func TestFactoryMethod(t *testing.T) {
	var factory OperatorFactory
	var operator Operator

	//factory = MinusOperatorFactory{}
	factory = PlusOperatorFactory{}

	operator = factory.Create()
	operator.SetA(1)
	operator.SetB(2)
	println(operator.Result())

}