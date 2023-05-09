//go:build wireinject
// +build wireinject

package main

import (
	"GoLangTricks/src/wire_example/controller"
	"GoLangTricks/src/wire_example/mapper"
	"GoLangTricks/src/wire_example/service"
	"github.com/google/wire"
)

func InitWire() (c *controller.HelloController, cancel func(), err error) {
	panic(wire.Build(
		mapper.Set,
		service.Set,
		controller.Set))
	return
}
