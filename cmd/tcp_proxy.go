package cmd

import (
	"io"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var (
	listenAddress string
	targetAddress string
)

var tcpProxy = &cobra.Command{
	Use:   "tproxy",
	Short: "tcp proxy",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ln, err := net.Listen("tcp", listenAddress)
		if err != nil {
			panic(err)
		}

		log.Printf("listen success! listent address:%s, target address:%s\n", listenAddress, targetAddress)
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("get conn from %s", conn.RemoteAddr())
			target, err := net.Dial("tcp", targetAddress)
			if err != nil {
				log.Println(err)
				conn.Close()
				continue
			}
			log.Printf("connect to %s success!", target.RemoteAddr())
			go transMessage(conn, target)
			go transMessage(target, conn)
		}
	},
}

func init() {
	rootCmd.AddCommand(tcpProxy)
	tcpProxy.Flags().StringVarP(&listenAddress, "listen_address", "l", "0.0.0.0:60000", "监听地址")
	tcpProxy.Flags().StringVarP(&targetAddress, "target_address", "t", "127.0.0.1:63438", "转发地址")
}

func transMessage(from net.Conn, to net.Conn) {
	defer from.Close()
	defer log.Printf("close connect %s", from.RemoteAddr())
	buf := make([]byte, 4096)
	for {
		n, err := from.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			return
		}

		n, err = to.Write(buf[:n])
		if err != nil && err != io.EOF {
			if err != io.EOF {
				log.Println(err)
			}
			return
		}
		log.Printf("%s ===> %s\n", from.RemoteAddr(), to.RemoteAddr())
	}
}
