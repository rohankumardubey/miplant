package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"

	"github.com/sapk/miplant/driver"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	//generic := ble.NewGenericAccessDriver(bleAdaptor)
	//information := ble.NewDeviceInformationDriver(bleAdaptor)
	//battery := ble.NewBatteryDriver(bleAdaptor)
	miplant := driver.NewMiPlantDriver(bleAdaptor)
	work := func() {
		gobot.Every(60*time.Second, func() {
			fmt.Println("Loop ...")
			/*
				fmt.Println("Device Name:", generic.GetDeviceName())
				fmt.Println("Appearance:", generic.GetAppearance())
			//*/
			/*
				fmt.Println("Battery level:", battery.GetBatteryLevel())
			//*/
			/*
				fmt.Println("Manufacturer:", information.GetManufacturerName())
				fmt.Println("ModelNumber:", information.GetModelNumber())
				fmt.Println("HardwareRevision:", information.GetHardwareRevision())
				fmt.Println("FirmwareRevision:", information.GetFirmwareRevision())
				fmt.Println("PnPId:", information.GetPnPId())
			//*/
			//*
			firm, err := miplant.GetFirmware()
			if err != nil {
				fmt.Println("Failed Firmware level:", err)
			} else {
				fmt.Println("Firmware level:", firm)
			}
			//*/
			name, err := miplant.GetName()
			if err != nil {
				fmt.Println("Failed GetName: ", err)
			} else {
				fmt.Println("GetName: ", name)
			}
			//*
			batt, err := miplant.GetBatteryLevel()
			if err != nil {
				fmt.Println("Failed Custom Battery level:", err)
			} else {
				fmt.Println("Custom Battery level:", batt)
			}
			//*/
		})
	}

	robot := gobot.NewRobot("miplantBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{miplant},
		work,
	)

	robot.Start()
}
