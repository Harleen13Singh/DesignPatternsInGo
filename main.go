package main

import (
	"main/designpatterns/abstractfactory"
	"main/designpatterns/builder"
	"main/designpatterns/factory"
	"main/designpatterns/prototype"
)

// main is our client code
func main() {
  factory.ClientCode()
  abstractfactory.ClientCode()
  builder.ClientCode()
  prototype.ClientCode()
}
