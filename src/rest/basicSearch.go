package rest

import ()

type BasicSearchHandler struct {
	RestHandlerBase
}

func NewBasicSearchHandler() *BasicSearchHandler {
	bsh := new(BasicSearchHandler)
	bsh.rh = bsh
	bsh.context = "/basic"
	return bsh
}

func (rhb *BasicSearchHandler) ServeRest(params map[string]string) (interface{}, *RestError) {
	//re := new(RestError)
	var values []string
	values = append(values, "one")
	values = append(values, "two")
	values = append(values, "red")
	values = append(values, "blue")
	values = append(values, "green")
	return values, nil
}
