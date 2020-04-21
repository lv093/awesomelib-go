package hello

type HelloAliceService struct {
}

func (this HelloAliceService) SaidHello(request string, reply *string) error {
	*reply = "hello, alice said:" + request
	println("said Hello:"+ request)
	return nil
}

func (this HelloAliceService) NotSaidHello(request string, reply *string) error {
	*reply = "hello, alice didnot said:" + request
	return nil
}