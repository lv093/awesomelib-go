package hello

type HelloJsonService struct {

}

func (this HelloJsonService) SaidJsonHello(request string, reply *string) error {
	*reply = "said:" + request
	return nil
}