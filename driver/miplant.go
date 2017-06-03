package driver

import (
	"bytes"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
)

//Based on :
// - https://wiki.hackerspace.pl/projects:xiaomi-flora
// - https://github.com/open-homeautomation/miflora/blob/master/miflora/miflora_poller.py
// - https://github.com/muka/go-bluetooth/blob/master/devices/SensorTag.go

var miplantTagUUIDs = map[string]string{
	"live":     "0x033",
	"name":     "0x03",
	"firmware": "0x038",
}

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

// adaptor returns BLE adaptor
func (m *MiPlantDriver) adaptor() ble.BLEConnector {
	return m.Connection().(ble.BLEConnector)
}

// Name returns the Driver name
func (m *MiPlantDriver) Name() string { return m.name }

// SetName sets the Driver name
func (m *MiPlantDriver) SetName(n string) { m.name = n }

// Connection returns the Driver's Connection to the associated Adaptor
func (m *MiPlantDriver) Connection() gobot.Connection { return m.connection }

// Start tells driver to get ready to do work
func (m *MiPlantDriver) Start() (err error) {

	//m.adaptor().Subscribe("0033", func(data []byte, err error) {
	//	fmt.Println("Event ...", data, err)
	//})
	m.adaptor().WriteCharacteristic(miplantTagUUIDs["live"], []byte{0xa0, 0x1f})
	return
}

// Halt stops battery driver (void)
func (m *MiPlantDriver) Halt() (err error) { return }

// GetFirmware returns the device firmware version
func (m *MiPlantDriver) GetFirmware() (string, error) {
	c, err := m.adaptor().ReadCharacteristic(miplantTagUUIDs["firmware"])
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(c)
	buf.ReadByte() //Skip battery
	return buf.String(), nil
}

// GetName returns the device name
func (m *MiPlantDriver) GetName() (string, error) {
	c, err := m.adaptor().ReadCharacteristic(miplantTagUUIDs["name"])
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(c)
	return buf.String(), nil
}

// GetBatteryLevel reads and returns the current battery level
func (m *MiPlantDriver) GetBatteryLevel() (level uint8, err error) {
	//*
	c, err := m.adaptor().ReadCharacteristic(miplantTagUUIDs["firmware"])
	if err != nil {
		return 0, err
	}
	buf := bytes.NewBuffer(c)
	val, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint8(val), nil
	//*/
	//return 0, nil
}
