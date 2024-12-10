package main

import "fmt"

type Lights struct{}

func (l *Lights) turnOn() {
	fmt.Println("Lights are turned on")
}

func (l *Lights) turnOff() {
	fmt.Println("Lights are turned off")
}

type Thermostat struct{}

func (t *Thermostat) setTemperature(temp int) {
	fmt.Printf("Thermostat set to %d°C\n", temp)
}

type SecuritySystem struct{}

func (s *SecuritySystem) arm() {
	fmt.Println("Security system armed")
}

func (s *SecuritySystem) disarm() {
	fmt.Println("Security system disarmed")
}

type SmartHomeFacade struct {
	lights         *Lights
	thermostat     *Thermostat
	securitySystem *SecuritySystem
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lights:         &Lights{},
		thermostat:     &Thermostat{},
		securitySystem: &SecuritySystem{},
	}
}

func (sh *SmartHomeFacade) leaveHome() {
	fmt.Println("Preparing the house for leaving...")
	sh.lights.turnOff()
	sh.thermostat.setTemperature(18)
	sh.securitySystem.arm()
}

func (sh *SmartHomeFacade) arriveHome() {
	fmt.Println("Preparing the house for arrival...")
	sh.lights.turnOn()
	sh.thermostat.setTemperature(22)
	sh.securitySystem.disarm()
}

func main() {
	smartHome := NewSmartHomeFacade()
	smartHome.leaveHome()
	fmt.Println()
	smartHome.arriveHome()
}

