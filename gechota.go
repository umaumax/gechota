package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	port     string
	network  string
	echoFlag bool
	verbose  bool
)

func init() {
	flag.StringVar(&port, "p", "3939", "port")
	flag.StringVar(&network, "network", "tcp", "network")
	flag.BoolVar(&echoFlag, "echo", true, "echo flag")
	flag.BoolVar(&verbose, "v", false, "verbose flag")
}

func main() {
	flag.Parse()

	if verbose {
		log.Printf("network:%s\n", network)
		log.Println("[Listening]", port)
	}
	if strings.HasPrefix(network, "tcp") {
		listener, err := net.Listen(network, ":"+port)
		if err != nil {
			log.Fatalln(err)
		}
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatalln(err)
			}
			if verbose {
				log.Printf("listen from %s to %s\n", conn.RemoteAddr(), conn.LocalAddr())
			}
			go func(conn net.Conn) {
				if echoFlag {
					duplicatedWriter := io.MultiWriter(os.Stdout, conn)
					io.Copy(duplicatedWriter, conn)
				} else {
					io.Copy(os.Stdout, conn)
				}
				conn.Close()
			}(conn)
		}
	}
	if strings.HasPrefix(network, "udp") {
		ra, err := net.ResolveUDPAddr(network, ":"+port)
		if err != nil {
			log.Fatalln(err)
		}
		conn, err := net.ListenUDP(network, ra)
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		for {
			const bufSize = 1024
			buf := make([]byte, bufSize)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				log.Fatalln(err)
			}
			buf = buf[:n]
			if echoFlag {
				os.Stdout.Write(buf)
				conn.WriteToUDP(buf, addr)
			} else {
				conn.WriteToUDP(buf, addr)
			}
		}
	}
}
