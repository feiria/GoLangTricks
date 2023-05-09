package controller

import "GoLangTricks/src/wire_example/service"

type HelloController struct {
	helloService service.HelloService
}

func NewHelloController(helloService service.HelloService) *HelloController {
	return &HelloController{
		helloService: helloService,
	}
}

func (h HelloController) HelloWorldController() string {
	return h.helloService.HelloWorldService()
}
