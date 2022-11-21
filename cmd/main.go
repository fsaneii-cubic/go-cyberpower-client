package main

import (
	"fmt"
	"os"
	"strconv"

	. "cyberpowerclient"
)

const (
	defaultAttempts = 5
	defaultTimeout  = 5
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	// CONNECT := arguments[1] + ":" + arguments[2]
	// fmt.Printf("connection string = %s", CONNECT)
	port, _ := strconv.Atoi(arguments[2])

	cp, err := NewCyberPowerPdu(arguments[1], port, defaultAttempts, defaultTimeout)
	if err != nil {
		panic(err)
	}

	outletIndex := 5
	outletState := "Off"
	err = cp.SetOutletPowerState(outletIndex, outletState)
	if err != nil {
		fmt.Printf("Error setting outlet#%d to state %s, %s", outletIndex, outletState, err.Error())
	}
	outletIndex = 6
	cp.SetOutletPowerState(outletIndex, outletState)
	if err != nil {
		fmt.Printf("Error setting outlet#%d to state %s, %s", outletIndex, outletState, err.Error())
	}

	outletIndex = 5
	outletState = "On"
	cp.SetOutletPowerState(outletIndex, outletState)
	if err != nil {
		fmt.Printf("Error setting outlet#%d to state %s, %s", outletIndex, outletState, err.Error())
	}

	outletState = "Cancel"
	cp.SetOutletPowerState(outletIndex, outletState)

	cp.Close()
	// c, err := net.Dial("tcp", CONNECT)
	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print(">> ")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Fprintf(c, text+"\n")

	// 	message, _ := bufio.NewReader(c).ReadString('\n')
	// 	fmt.Print("->: " + message)
	// 	if strings.TrimSpace(string(text)) == "STOP" {
	// 		fmt.Println("TCP client exiting...")
	// 		return
	// 	}
	// }
}
