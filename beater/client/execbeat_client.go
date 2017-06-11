package main

import (
    "fmt"
    "net"
//    "time"
//    "strconv"
//    "encoding/json"
    "os"
    "github.com/christiangalsterer/execbeat/config"
    "github.com/elastic/beats/libbeat/logp"
    "gopkg.in/yaml.v2"
)

func CheckError(err error) {
    if err  != nil {
        logp.Err("execbeat", "Error: " , err)
    }
}

func main() {
    if (len(os.Args) == 1) {
	logp.Err("execbeat", "Not enough arguments")
	logp.Err("execbeat", os.Args)
	return
    }
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

    argsWithProg := os.Args

    req := &config.ExecConfig{
	Schedule: "",
	Command: "",
	Args: "",
	DocumentType: "",
	Fields: nil}


    if (len(os.Args[1:]) > 1){
	req.Command = argsWithProg[1]

	out := ""
	for _, str := range os.Args[2:] {
		out = fmt.Sprintf("%s%s%s", out, str, " ")
	}
	req.Args = out
    }else{
	req.Command = argsWithProg[1]
    }

    reqMar, _ := yaml.Marshal(req)
    logp.Debug("execbeat", "client %s", string(reqMar))

    _,err = Conn.Write(reqMar)
    CheckError(err)

    defer Conn.Close()
}
