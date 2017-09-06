package handler

import "net"

type HandlerFunc func(connection net.Conn)
