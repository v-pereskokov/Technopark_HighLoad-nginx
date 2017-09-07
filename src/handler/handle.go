package handler

import "net"

type HandlerFunc func(connection chan net.Conn)
