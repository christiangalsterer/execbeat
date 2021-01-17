package beat

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/christiangalsterer/execbeat/config"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/robfig/cron/v3"
)

type Executor struct {
	execbeat     *Execbeat
	config       config.ExecConfig
	schedule     string
	documentType string
}

func NewExecutor(execbeat *Execbeat, config config.ExecConfig) *Executor {
	executor := &Executor{
		execbeat: execbeat,
		config:   config,
	}

	return executor
}

func (e *Executor) Run() {

	// setup default config
	e.documentType = config.DefaultDocumentType
	e.schedule = config.DefaultSchedule

	// setup document type
	if e.config.DocumentType != "" {
		e.documentType = e.config.DocumentType
	}

	// setup cron schedule
	if e.config.Schedule != "" {
		logp.Debug("Execbeat", "Use schedule: [%w]", e.config.Schedule)
		e.schedule = e.config.Schedule
	}

	cron := cron.New()
	cron.AddFunc(e.schedule, func() { e.runOneTime() })
	cron.Start()
}

func (e *Executor) runOneTime() error {
	var cmd *exec.Cmd
	var cmdArgs []string
	var err error
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var waitStatus syscall.WaitStatus
	var exitCode int = 0

	cmdName := strings.TrimSpace(e.config.Command)

	args := strings.TrimSpace(e.config.Args)
	if len(args) > 0 {
		cmdArgs = strings.Split(args, " ")
	}

	// execute command
	now := time.Now()

	if len(cmdArgs) > 0 {
		logp.Debug("Execbeat", "Executing command: [%v] with args [%w]", cmdName, cmdArgs)
		cmd = exec.Command(cmdName, cmdArgs...)
	} else {
		logp.Debug("Execbeat", "Executing command: [%v]", cmdName)
		cmd = exec.Command(cmdName)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Start()
	if err != nil {
		logp.Err("An error occured while executing command: %v", err)
		exitCode = 127
	}

	err = cmd.Wait()
	if err != nil {
		logp.Err("An error occured while executing command: %v", err)
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			exitCode = waitStatus.ExitStatus()
		}
	}

	logp.Info("Execbeat", "Executing command: [%v]", stdout.String())

	commandEvent := Exec{
		Command:  cmdName,
		StdOut:   stdout.String(),
		StdErr:   stderr.String(),
		ExitCode: exitCode,
	}

	event := ExecEvent{
		ReadTime:     now,
		DocumentType: e.documentType,
		Fields:       e.config.Fields,
		Exec:         commandEvent,
	}

	e.execbeat.client.PublishEvent(event.ToMapStr())

	return nil
}

func (e *Executor) Stop() {
}
