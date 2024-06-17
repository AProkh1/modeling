package main

import (
	"fmt"
	"time"
)

type Light struct {
	on bool
}

func (l *Light) TurnOn() {
	l.on = true
	fmt.Println("Свет включен")
}

func (l *Light) TurnOff() {
	l.on = false
	fmt.Println("Свет выключен")
}

type Heater struct {
	on bool
}

func (h *Heater) TurnOn() {
	h.on = true
	fmt.Println("Обогреватель включен")
}

func (h *Heater) TurnOff() {
	h.on = false
	fmt.Println("Обогреватель выключен")
}

type Humidifier struct {
	on bool
}

func (h *Humidifier) TurnOn() {
	h.on = true
	fmt.Println("Увлажнитель воздуха включен")
}

func (h *Humidifier) TurnOff() {
	h.on = false
	fmt.Println("Увлажнитель воздуха выключен")
}

type Door struct {
	closed bool
}

func (d *Door) Open() {
	d.closed = false
	fmt.Println("Дверь открыта")
}

func (d *Door) Close() {
	d.closed = true
	fmt.Println("Дверь закрыта")
}

type Window struct {
	open bool
}

func (w *Window) Open() {
	w.open = true
	fmt.Println("Окно открыто")
}

func (w *Window) Close() {
	w.open = false
	fmt.Println("Окно закрыто")
}

type TemperatureSensor struct {
	temperature float32
}

func (s *TemperatureSensor) GetTemperature() float32 {
	return s.temperature
}

func (s *TemperatureSensor) SetTemperature(temp float32) {
	s.temperature = temp
}

type LightSensor struct {
	light bool
}

func (s *LightSensor) IsLightOn() bool {
	return s.light
}

func (s *LightSensor) SetLightStatus(status bool) {
	s.light = status
}

type HumiditySensor struct {
	humidity float32
}

func (s *HumiditySensor) GetHumidity() float32 {
	return s.humidity
}

func (s *HumiditySensor) SetHumidity(hum float32) {
	s.humidity = hum
}

type SmartHome struct {
	lightSensor       *LightSensor
	temperatureSensor *TemperatureSensor
	humiditySensor    *HumiditySensor

	light      *Light
	heater     *Heater
	humidifier *Humidifier
	door       *Door
	window     *Window

	isWorking bool
}

func (sh *SmartHome) TurnOn() {
	sh.isWorking = true
	fmt.Println("Умный дом включен")
}

func (sh *SmartHome) TurnOff() {
	sh.isWorking = false
	fmt.Println("Умный дом выключен")
}

func (sh *SmartHome) Run() {
	for sh.isWorking {
		// Пример работы системы
		if sh.lightSensor.IsLightOn() {
			sh.light.TurnOn()
		} else {
			sh.light.TurnOff()
		}

		if sh.temperatureSensor.GetTemperature() < 20 {
			sh.heater.TurnOn()
		} else {
			sh.heater.TurnOff()
		}

		if sh.humiditySensor.GetHumidity() < 30 {
			sh.humidifier.TurnOn()
		} else {
			sh.humidifier.TurnOff()
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	lightSensor := &LightSensor{}
	temperatureSensor := &TemperatureSensor{}
	humiditySensor := &HumiditySensor{}

	light := &Light{}
	heater := &Heater{}
	humidifier := &Humidifier{}
	door := &Door{}
	window := &Window{}

	smartHome := &SmartHome{
		lightSensor:       lightSensor,
		temperatureSensor: temperatureSensor,
		humiditySensor:    humiditySensor,
		light:             light,
		heater:            heater,
		humidifier:        humidifier,
		door:              door,
		window:            window,
		isWorking:         true,
	}

	smartHome.TurnOn()

	lightSensor.SetLightStatus(true)
	temperatureSensor.SetTemperature(18.0)
	humiditySensor.SetHumidity(25.0)

	go smartHome.Run()

	time.Sleep(10 * time.Second)

	smartHome.TurnOff()
}
