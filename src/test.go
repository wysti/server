package main

import (
	"fmt"
	"rest"
)

type WystiAPI interface {
	Login() rest.APIProvider
	GetArguments() rest.APIProvider
	CreateWord() rest.APIProvider
}

type Wysti struct {
}

func (w *Wysti) Login() rest.HandlerSpec {
	var hs rest.HandlerSpec
	hs.Context = "/login"
	hs.ServeRest = func(map[string][]string) (interface{}, *rest.RestError) {
		//var re RestError
		testing := "123"
		fmt.Println("?", testing)
		return testing, nil
	}
	fmt.Println("logging in...")
	return hs
}

func (w *Wysti) GetArguments() rest.HandlerSpec {
	var hs rest.HandlerSpec
	hs.Context = "/getArgs"
	hs.ServeRest = func(map[string][]string) (interface{}, *rest.RestError) {
		//var re RestError
		testing := "345"
		fmt.Println("?", testing)
		return testing, nil
	}
	fmt.Println("Getting args...")
	return hs
}

func (w *Wysti) CreateWord() rest.HandlerSpec {
	var hs rest.HandlerSpec
	hs.Context = "/createWord"
	hs.ServeRest = func(map[string][]string) (interface{}, *rest.RestError) {
		//var re RestError
		testing := [...]string{"678", "hwllo", "no", "hello"}
		fmt.Println("?", testing)
		return testing, nil
	}
	fmt.Println("Creating...")
	return hs
}

func main() {
	rest.InitializeErrors()
	w := new(Wysti)
	rest.AcceptRequests(w)
}
