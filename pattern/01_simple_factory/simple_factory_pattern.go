package simple_factory

type Animal interface {
	Sing(name string)
}

func NewAnimal(t string) Animal {
	if t == "bird" {
		return &bird{}
	} else if t == "fish" {
		return &fish{}
	} else {
		return nil
	}
}

type bird struct {
}

func (this *bird) Sing(name string) {
	println("bird:" + name)
}

type fish struct {
}
func (this *fish) Sing(name string) {
	println("fish:" + name)
}