package service

import "GoLangTricks/src/wire_example/mapper"

type helloService struct {
	helloMapper mapper.HelloMapper
}

func NewHelloService(helloMapper mapper.HelloMapper) HelloService {
	return &helloService{
		helloMapper: helloMapper,
	}
}

func (h helloService) HelloWorldService() string {
	return h.helloMapper.HelloWorldMapper()
}
