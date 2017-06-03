package main

import (
	"github.com/muka/go-bluetooth/api"
	"github.com/sapk/miplant/device"

	"github.com/op/go-logging"
	"github.com/tj/go-debug"
)

var log = logging.MustGetLogger("main")
var dbg = debug.Debug("bluez:main")

var adapterID = "hci0"               //TODO arg
var tagAddress = "C4:7C:8D:61:F7:81" //TODO arg

func main() {
	devList, err := api.GetDevices()
	for _, d := range devList {
		log.Info(d.Properties.Address)
	}
	if err != nil {
		log.Fatal(err)
	}
	dev, err := api.GetDeviceByAddress(tagAddress)
	//dev, err := api.GetDevices()
	if err != nil {
		log.Fatal(err)
	}
	if dev == nil {
		log.Fatal("Device not found")
	}
	/*
		err = dev.Connect()
		if err != nil {
			log.Fatal(err)
		}
	*/
	miplant, err := device.NewMiPlant(dev) //Use 0
	if err != nil {
		log.Fatal(err)
	}
	log.Info(miplant, err)
}
