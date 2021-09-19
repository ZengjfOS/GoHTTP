package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"

	"log"
	"os/exec"
	"strconv"
)

func main() {
	info := accessory.Info{
		Name:         "Lamp",
		SerialNumber: "051AC-23AAM1",
		Manufacturer: "Apple",
		Model:        "AB",
	}
	acc := accessory.NewSwitch(info)

	var action int = 0

	cmd := exec.Command("bash", "-c", "gpio export 9 out")
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Printf("output: %s", data)

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Println("Client changed switch to on")
			action = 0
		} else {
			log.Println("Client changed switch to off")
			action = 1
		}

		cmd = exec.Command("bash", "-c", "echo " + strconv.Itoa(action) + " > /sys/class/gpio/gpio9/value")

		data, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
		}

		log.Printf("output: %s", data)
	})

	config := hc.Config{Pin: "00102003"}
	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
