package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
)

func main() {
	done := make(chan bool)
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	miplant := NewMiPlantDriver(bleAdaptor)

	work := func() {
		gobot.Every(5*time.Second, func() {
			fmt.Println("Getting battery level ...")
			fmt.Println("Battery level:", miplant.GetBatteryLevel())
		})
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{miplant},
		work,
	)

	robot.Start()
	<-done
}
