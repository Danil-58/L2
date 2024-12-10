package main

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
}

type ComputerBuilder interface {
	BuildCore() *ComputerBuilder
	BuildBrand() *ComputerBuilder
	BuildMemory() *ComputerBuilder
	BuildGraphicCard() *ComputerBuilder
	GetComputer() *Computer
}

type ConcreteComputerBuilder struct {
	computer *Computer
}

func NewConcreteComputerBuilder() *ConcreteComputerBuilder {
	return &ConcreteComputerBuilder{computer: &Computer{}}
}

func (b *ConcreteComputerBuilder) BuildCore() *ConcreteComputerBuilder {
	b.computer.Core = 8
	return b
}

func (b *ConcreteComputerBuilder) BuildBrand() *ConcreteComputerBuilder {
	b.computer.Brand = "CustomBrand"
	return b
}

func (b *ConcreteComputerBuilder) BuildMemory() *ConcreteComputerBuilder {
	b.computer.Memory = 16
	return b
}

func (b *ConcreteComputerBuilder) BuildGraphicCard() *ConcreteComputerBuilder {
	b.computer.GraphicCard = 4
	return b
}

func (b *ConcreteComputerBuilder) GetComputer() *Computer {
	return b.computer
}

func main() {
	builder := NewConcreteComputerBuilder()
	computer := builder.BuildCore().BuildBrand().BuildMemory().BuildGraphicCard().GetComputer()

	fmt.Println("Built computer with the following specifications:")
	fmt.Printf("Core: %d\n", computer.Core)
	fmt.Printf("Brand: %s\n", computer.Brand)
	fmt.Printf("Memory: %dGB\n", computer.Memory)
	fmt.Printf("Graphic Card: %dGB\n", computer.GraphicCard)
}
