package handler

import "net"

type HandlerFunc func(chan net.Conn)
