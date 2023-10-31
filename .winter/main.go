package main

import (
	"go-clean-monolith/.winter/validator"
)

func main() {
	validator.StructureValidation()
	validator.LayersValidation()
	validator.EndpointsValidation()
}
