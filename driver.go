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
		name:       gobot.DefaultName("MiPLant"),
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
	return
}

// Halt stops battery driver (void)
func (m *MiPlantDriver) Halt() (err error) { return }

// GetBatteryLevel reads and returns the current battery level
func (m *MiPlantDriver) GetBatteryLevel() (level uint8) {
	var l uint8
	c, _ := m.adaptor().ReadCharacteristic("0038")
	buf := bytes.NewBuffer(c)
	val, _ := buf.ReadByte()
	l = uint8(val)
	return l
}
