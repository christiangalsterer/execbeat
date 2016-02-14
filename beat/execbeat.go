package beat

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"time"
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
	execBeat.done = make(chan struct{})

	return nil
}

func (execBeat *Execbeat) Run(b *beat.Beat) error {
	var err error

	var poller *Executor

	for i, exitConfig := range execBeat.ExecConfig.Execbeat.Execs {
		logp.Debug("execbeat", "Creating poller # %v with command: %v", i, exitConfig.Command)
		poller = NewExecutor(execBeat, exitConfig)
		go poller.Run()
	}

	ticker := time.NewTicker(3600 * time.Second)
	defer ticker.Stop()
	for {
		select {
			case <-execBeat.done:
				return nil
			case <-ticker.C:
		}
	}

	return err
}

func (execBeat *Execbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (execBeat *Execbeat) Stop() {
	close(execBeat.done)
}
