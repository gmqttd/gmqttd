package main
import (
    "fmt"
    "net"
)
func main() {
    starServer()
}

func starServer() {
    listener, err := net.Listen("tcp", "localhost:9399")
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }else {
        fmt.Printf("Listen on: localhost:9399\n")
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("err:%v\n", err)
        }else {
            fmt.Printf("remote conn req:%v\n", conn.RemoteAddr())
        }
        go handConn(conn)
    }
}

func handConn(conn net.Conn) {
    readBuf := make([]byte, 512)

    for {
        len, err := conn.Read(readBuf)
        if err != nil {
            fmt.Printf("read err:%v\n", err)
        }
        fmt.Printf("readBuf:%v\n", readBuf[:len])
    }
}
