package beat

import (
    //"fmt"
    "net"
    "github.com/christiangalsterer/execbeat/config"
    "github.com/elastic/beats/libbeat/logp"
    "gopkg.in/yaml.v2"
)

func CheckError(err error) {
    if err  != nil {
        logp.Err("Error: " , err)
    }
}

func CmdMonitorServerLoop(exexBeat *Execbeat) {
    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
    CheckError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
	var poller *Executor

        n,addr,err := ServerConn.ReadFromUDP(buf)
	if err != nil {
		//fmt.Println(err);
		logp.Err("execbeat", err)
	}
        logp.Debug("execbeat", "Received from %s", addr)
	logp.Debug("execbeat", "%s", string(buf[0:n]))
	//fmt.Println(string(buf[0:n]))

	//fmt.Println(buf)
	data := config.ExecConfig{}
	//FIXME: Unmarshalling will fail if stream/data is segmented
	if err := yaml.Unmarshal([]byte(string(buf[0:n])), &data); err != nil {
		logp.Err("execbeat", err)
		continue
	}

	logp.Debug("execbeat", "Creating poller with command: %v", data.Command)

	//Insert the command received from client in ExecConfig slice
	exexBeat.ExecConfig.Execbeat.Commands = append(exexBeat.ExecConfig.Execbeat.Commands, data)
	poller = NewExecutor(exexBeat, data)
	go poller.Run()
    }
}
