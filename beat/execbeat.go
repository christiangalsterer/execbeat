package beat

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
)

type Execbeat struct {
	done chan struct{}
	ExecConfig ConfigSettings
	events publisher.Client
}

func New() *Execbeat {
	return &Execbeat{}
}

func (execBeat *Execbeat) Config(b *beat.Beat) error {

	err := cfgfile.Read(&execBeat.ExecConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	logp.Info("execbeat", "Init execbeat")

	return nil
}

func (execBeat *Execbeat) Setup(b *beat.Beat) error {
	execBeat.events = b.Events

	return nil
}

func (exexBeat *Execbeat) Run(b *beat.Beat) error {
	var err error

	var poller *Executor

	for i, exitConfig := range exexBeat.ExecConfig.Execbeat.Execs {
		logp.Debug("execbeat", "Creating poller # %v with command: %v", i, exitConfig.Command)
		poller = NewExecutor(exexBeat, exitConfig)
		go poller.Run()
	}

	for {
	}

	return err
}

func (execBeat *Execbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (execBeat *Execbeat) Stop() {
	close(execBeat.done)
}
