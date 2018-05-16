package main

import (
    "log"
    "flag"

    "github.com/tarm/serial"
)


func main() {
    var port = flag.Arg(0)
    //var port = flag.String("port", "COM45", "--port PORTNAME")
    flag.Parse()


    c := &serial.Config { Name: port, Baud: 115200 }

    s, err := serial.OpenPort(c)
    if err != nil {
            log.Fatal(err)
    }

    n, err := s.Write([]byte("test"))
    if err != nil {
            log.Fatal(err)
    }

    buf := make([]byte, 128)
    n, err = s.Read(buf)
    if err != nil {
            log.Fatal(err)
    }

    log.Printf("%q", buf[:n])
}

