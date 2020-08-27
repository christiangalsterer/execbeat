package beat

import (
	"fmt"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/sonnylaskar/execbeat/config"
)

type Execbeat struct {
	done       chan bool
	ExecConfig config.ConfigSettings
	client     publisher.Client
}

func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	execbeat := &Execbeat{
		done: make(chan bool),
	}

	err := cfgfile.Read(&execbeat.ExecConfig, "")
	if err != nil {
		logp.Err("Error reading config file: %v", err)
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	return execbeat, nil
}

func (exexBeat *Execbeat) Run(b *beat.Beat) error {
	var poller *Executor
	logp.Info("execbeat is running! Hit CTRL-C to stop it.")
	exexBeat.client = b.Publisher.Connect()

	for i, exitConfig := range exexBeat.ExecConfig.Execbeat.Commands {
		logp.Debug("execbeat", "Creating poller # %v with command: %v", i, exitConfig.Command)
		poller = NewExecutor(exexBeat, exitConfig)
		go poller.Run()
	}

	for {
		select {
		case <-exexBeat.done:
			return nil
		}
	}
}

func (execBeat *Execbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (execBeat *Execbeat) Stop() {
	close(execBeat.done)
}
