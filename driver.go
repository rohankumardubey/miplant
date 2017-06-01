package main

import (
	"bytes"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
)

//Based on https://wiki.hackerspace.pl/projects:xiaomi-flora

// MiPlantDriver represents the MiPlantDriver Service for a BLE Peripheral
type MiPlantDriver struct {
	name       string
	connection gobot.Connection
	gobot.Eventer
}

// NewMiPlantDriver creates a BatteryDriver
func NewMiPlantDriver(a ble.BLEConnector) *MiPlantDriver {
	n := &MiPlantDriver{
		name:       gobot.DefaultName("Flower care"),
		connection: a,
		Eventer:    gobot.NewEventer(),
	}
	return n
}

// Connection returns the Driver's Connection to the associated Adaptor
func (m *MiPlantDriver) Connection() gobot.Connection { return m.connection }

// Name returns the Driver name
func (m *MiPlantDriver) Name() string { return m.name }

// SetName sets the Driver name
func (m *MiPlantDriver) SetName(n string) { m.name = n }

// adaptor returns BLE adaptor
func (m *MiPlantDriver) adaptor() ble.BLEConnector {
	return m.Connection().(ble.BLEConnector)
}

// Start tells driver to get ready to do work
func (m *MiPlantDriver) Start() (err error) {

	//m.adaptor().Subscribe("0033", func(data []byte, err error) {
	//	fmt.Println("Event ...", data, err)
	//})
	m.adaptor().WriteCharacteristic("0033", []byte{0x1f, 0xa0})
	return
}

// Halt stops battery driver (void)
func (m *MiPlantDriver) Halt() (err error) { return }

// GetBatteryLevel reads and returns the current battery level
func (m *MiPlantDriver) GetBatteryLevel() (level uint8, err error) {
	//*
	var l uint8
	c, err := m.adaptor().ReadCharacteristic("0038") //"0038"
	if err != nil {
		return 0, err
	}
	buf := bytes.NewBuffer(c)
	val, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	l = uint8(val)
	return l, nil
	//*/
	//return 0, nil
}
