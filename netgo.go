package main

import (
	"flag"
	"fmt"
	"github.com/grt1st/netgo/netextends"
	//"github.com/grt1st/netgo/utils"
	"log"
	"net"
	"os"
	"os/signal"
)

const versionNumber = "1.1.0#20180610"

func main() {
	version := flag.Bool("version", false, "Show program's version number and exit")
	listen := flag.Bool("l", false, "Listen on the server")
	addr := flag.String("a", "", "Address to use")
	port := flag.Int("p", 0, "Port to use ")
	port1 := flag.Int("rp", 0, "Another port to use ")
	htmlFlag := flag.Bool("html", false, "Send http request of GET")
	exeCmd := flag.String("e", "", "Use pipe to transform command")
	//max := flag.Int("m", 1, "Max client num")
	rhost := flag.String("rhost", "", "The remote address to connect")
	help := flag.Bool("h", false, "Show usage")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  %s [Options]\n\nOptions\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Println(versionNumber)
		return
	}
	if *help || *port == 0 || (*addr == "" && *listen == false && *rhost == "") {
		flag.Usage()
		return
	}

	if *listen {
		if *port1 == 0 {
			listenE(*addr, *port, *exeCmd)
		} else {
			netextends.ServerAndServer(*addr, *port, *port1)
		}
	} else {
		if *rhost != "" {
			netextends.RemotePortForward(*addr, *port, *rhost)
		} else {
			connectS(*addr, *port, *htmlFlag, *exeCmd)
		}
	}

}

func connectS(addr string, port int, htmlFlag bool, exeCmd string) {

	stopChan := make(chan os.Signal) // 接收系统中断信号
	signal.Notify(stopChan, os.Interrupt)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		<-stopChan
		if err = conn.Close(); err != nil {
			log.Print(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	if htmlFlag {
		netextends.ConnectHtmlMode(conn)
	} else if exeCmd != "" {
		netextends.ConnectExecMode(conn, exeCmd)
	} else {
		netextends.ConnectNormalMode(conn)
	}
}

func listenE(addr string, port int, exeCmd string) {
	stopChan := make(chan os.Signal) // 接收系统中断信号
	signal.Notify(stopChan, os.Interrupt)

	if addr == "" {
		addr = "localhost"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//listener = utils.LimitListener(listener, 1)

	go func() {
		<-stopChan
		if err = listener.Close(); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	if exeCmd != "" {
		netextends.ListenExecMode(listener, exeCmd)
	} else {
		netextends.ListenNormalMode(listener)
	}
}
