package cyberpowerclient

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type CyberPwerPdu interface {
	SetOutletPowerState(outletIndex int, action string) error
	Close()
}

type cyberPower struct {
	conn     net.Conn
	attempts int
	timeout  int
}

//var conn net.Conn

func NewCyberPowerPdu(hostname string, port, attempts, timeout int) (CyberPwerPdu, error) {
	connection, err := connect(hostname, port)
	if err != nil {
		return nil, err
	}
	return &cyberPower{
		conn:     connection,
		attempts: attempts,
		timeout:  timeout,
	}, nil
}

func (cp *cyberPower) SetOutletPowerState(outletIndex int, action string) error {
	// CONNECT := cp.host + ":" + strconv.Itoa(cp.port)
	// fmt.Printf("connection string = %s", CONNECT)
	// conn, err := net.Dial("tcp", CONNECT)
	// if err != nil {
	// 	return err
	// }

	// for {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print(">> ")
	//text, _ := reader.ReadString('\n')
	text := strconv.Itoa(outletIndex) + ", " + action
	fmt.Fprintf(cp.conn, text+"\n")

	message, _ := bufio.NewReader(cp.conn).ReadString('\n')
	fmt.Print("->: " + message)
	if strings.TrimSpace(string(text)) == "CANCEL" ||
		strings.TrimSpace(string(text)) == "Cancel" {
		fmt.Println("TCP client exiting...")
		return nil
	}
	// }

	return nil
}

func connect(hostname string, port int) (net.Conn, error) {
	CONNECT := hostname + ":" + strconv.Itoa(port)
	fmt.Printf("connection string = %s", CONNECT)
	conn, err := net.Dial("tcp", CONNECT)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (cp *cyberPower) Close() {
	if cp.conn != nil {
		defer cp.conn.Close()
	}
}
