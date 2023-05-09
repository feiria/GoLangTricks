package mapper

type helloMapper struct {
}

func NewHelloMapper() HelloMapper {
	return &helloMapper{}
}

func (h helloMapper) HelloWorldMapper() string {
	return "hello, world"
}
