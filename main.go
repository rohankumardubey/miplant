package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/examples/lib/dev"

	"golang.org/x/net/context"
)

var (
	device   = flag.String("device", "default", "implementation of ble")
	addr     = flag.String("addr", "", "Address of remote device")
	timeout  = flag.Duration("t", 5*time.Second, "timeout for search of addr")
	interval = flag.Duration("i", time.Minute, "interval between collection")
)

var battery uint8
var firmware string

func exec() {
	client, err := connect(*addr)
	if err != nil {
		log.Fatalf("connect failed : %s", err)
	}
	log.Printf("Connected to %s", client.Address())

	if battery == 0 {
		battery, cached, err := getBattery(client)
		log.Printf("Battery %v %v %v", battery, cached, err)

		firmware, cached, err := getFirmware(client)
		log.Printf("Firmware %v %v %v", firmware, cached, err)
	}

	err = activateRealtimeData(client)
	//log.Printf("Set realtime %v", err)
	data, cached, err := getData(client)
	//log.Printf("Data %v %v %v", data, cached, err)
	temp, ligth, fert, moist := parseData(data)

	log.Printf("Temperature: %v °C, Moisture: %v %%, Light: %v lux, Fertility: %v uS/cm", temp, moist, ligth, fert)

	client.CancelConnection()
}
func main() {
	flag.Parse()

	d, err := dev.NewDevice(*device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	ticker := time.NewTicker(*interval)
	quit := make(chan struct{})
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				exec()
			case <-quit:
				ticker.Stop()
				done <- true //Trigger at end
				return
			}
		}
	}()

	exec()
	<-done
}

func activateRealtimeData(client ble.Client) error {
	return client.WriteCharacteristic(&ble.Characteristic{
		ValueHandle: 0x33,
	}, []byte{0xa0, 0x1f}, false)
}

func parseData(data []byte) (temp float64, ligth, fert, moist uint) {
	temp = (float64(data[1])*256 + float64(data[0])) / 10
	moist = uint(data[7])
	ligth = uint(data[4])*256 + uint(data[3])
	fert = uint(data[9])*256 + uint(data[8])
	return temp, ligth, fert, moist
}

func getData(client ble.Client) ([]byte, bool, error) {
	b, err := client.ReadCharacteristic(&ble.Characteristic{
		ValueHandle: 0x35,
	})
	return b, false, err
}
func getFirmware(client ble.Client) (string, bool, error) {
	b, err := client.ReadCharacteristic(&ble.Characteristic{
		ValueHandle: 0x38,
	})
	return string(b[2:]), false, err
}

func getBattery(client ble.Client) (uint8, bool, error) {
	b, err := client.ReadCharacteristic(&ble.Characteristic{
		ValueHandle: 0x38,
	})
	return uint8(b[0]), false, err
}

func connect(addr string) (ble.Client, error) {
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), *interval))
	if addr != "" {
		bleAddr := ble.NewAddr(addr)
		fmt.Printf("Dialing to specified address: %s\n", bleAddr)
		return ble.Dial(ctx, bleAddr)
	}

	return nil, fmt.Errorf("no addr specified")
}
