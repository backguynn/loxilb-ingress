package pkg

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	LoxiLBImg = "/root/loxilb-io/loxilb/loxilb"
)

func SpawnLoxiLB() {
	for {

		command := fmt.Sprintf("%s --blacklist=eth0", LoxiLBImg)
		cmd := exec.Command("bash", "-c", command)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(3000 * time.Millisecond)
	}
}