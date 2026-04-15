package shared

import "net"


func StartGRPCServer(port string) (net.Listener, error) {
    return net.Listen("tcp", ":" + port)
}