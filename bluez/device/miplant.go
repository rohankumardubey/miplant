package device

//Based on :
// - https://wiki.hackerspace.pl/projects:xiaomi-flora
// - https://github.com/open-homeautomation/miflora/blob/master/miflora/miflora_poller.py
// - https://github.com/muka/go-bluetooth/blob/master/devices/SensorTag.go

import (
	"github.com/godbus/dbus"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/emitter"

	"github.com/op/go-logging"
	"github.com/tj/go-debug"
)

var log = logging.MustGetLogger("main")
var dbg = debug.Debug("bluez:device")

var miplantTagUUIDs = map[string]string{
	"name": "0x03",
}

//MiPlant a MiPlant object representation
type MiPlant struct {
	*api.Device
	dataChannel chan dbus.Signal
}

// NewMiPlant creates a new miplant instance
func NewMiPlant(d *api.Device) (*MiPlant, error) {
	m := &MiPlant{
		Device:      d,
		dataChannel: make(chan dbus.Signal),
	}

	var connect = func(dev *api.Device) error {
		if !dev.IsConnected() {
			err := dev.Connect()
			if err != nil {
				return err
			}
		}
		return nil
	}

	d.On("changed", emitter.NewCallback(func(ev emitter.Event) {
		changed := ev.GetData().(api.PropertyChangedEvent)
		if changed.Field == "Connected" {
			if !changed.Value.(bool) {
				dbg("Device disconnected")
				if m.dataChannel != nil {
					close(m.dataChannel)
				}
			}
		}
	}))

	err := connect(d)
	if err != nil {
		log.Warning("MiPlant connection failed: %v", err)
		return nil, err
	}

	return m, nil
}

// GetName return the sensor name
func (m MiPlant) GetName() string {
	return "TODO"
}
