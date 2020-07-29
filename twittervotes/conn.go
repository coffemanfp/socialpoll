package main

import (
	"io"
	"net"
	"time"
)

// Conn and data reader
var conn net.Conn
var reader io.ReadCloser

func dial(netw, addr string) (netc net.Conn, err error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}

	netc, err = net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return
	}

	conn = netc
	return
}

func closeConn() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		reader.Close()
	}
}
