package main 

import "fmt"
import "flag"

var CTRL_PORT = 64321

func main() {
	modePtr := flag.String("mode", "server", "server or client")
	portPtr := flag.Int("port", CTRL_PORT, "port for interacting with control channel")
	protoPtr := flag.String("protocol", "udp_uc", "protocol for discovering mapago servers: udp_uc, udp_mc or tcp")
	addrPtr := flag.String("addr", "ipv4_uc", "target addr: ipv4_uc, ipv4_mc, ipv6_uc, ipv6_mc") 

	flag.Parse()

	fmt.Println("mapago-light(c) - 2018")
	fmt.Println("Mode:", *modePtr)
	fmt.Println("Port:", *portPtr)
	fmt.Println("Protocol:", *protoPtr)
	fmt.Println("Address:", *addrPtr)

	if *modePtr == "server" {
		run_server(*portPtr)
	} else if *modePtr == "client" {
		run_client(*portPtr, *addrPtr, *protoPtr)
	} else {
		panic("mode not supported! chose server or client")
	}
}

func run_server(port int) {
	protos := supported_disco_protos()
	fmt.Println("supported protos: ", protos)

	for _, disco_proto := range protos {
		fmt.Println("next protocol to verify: ", disco_proto)
	}
}

func supported_disco_protos() []string {
	var disco_protos []string
	disco_protos = append(disco_protos, "udp4_uc", "udp6_uc")
	disco_protos = append(disco_protos, "udp4_mc", "udp6_mc")
	disco_protos = append(disco_protos, "tcp4_uc", "tcp6_uc")
	// placeholder for possible QUIC support

	return disco_protos
}

func run_client(port int, addr string, proto string) {
	fmt.Println("client handler dummy func")
}