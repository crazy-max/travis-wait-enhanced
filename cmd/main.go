package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version = "dev"
	cli     Cli
)

// Cli holds command line args, flags and cmds
type Cli struct {
	Version        kong.VersionFlag
	Timeout        time.Duration `kong:"name='timeout',default='20m',help='Timeout for this command.'"`
	Interval       time.Duration `kong:"name='interval',default='1m',help='Interval at which to print keep-alive messages.'"`
	PrintName      bool          `kong:"name='print-name',default='true',help='Print the name of this tool to identify keep-alive messages.'"`
	PrintString    string        `kong:"name='print-string',default='Still running...',help='Keep-alive message printed in each interval.'"`
	PrintTimestamp bool          `kong:"name='print-timestamp',default='true',help='Print the current timestamp after each keep-alive message.'"`
	PrintNewline   bool          `kong:"name='print-newline',default='true',help='Print a newline character after each keep-alive message.'"`
	Command        []string      `kong:"required,arg,name='command',help='Command to execute.'"`
}

func main() {
	// Parse command line
	kctx := kong.Parse(&cli,
		kong.Name("travis-wait-enhanced"),
		kong.Description(`Prevent Travis CI from thinking a long-running process has stalled. More info: https://github.com/crazy-max/travis-wait-enhanced`),
		kong.UsageOnError(),
		kong.Vars{
			"version": fmt.Sprintf("%s", version),
		},
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	// Logger
	output := zerolog.ConsoleWriter{
		Out: os.Stdout,
	}
	output.FormatTimestamp = func(i interface{}) string {
		return kctx.Model.Name
	}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Handle os signals
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-channel
		log.Warn().Msgf("Caught signal %v", sig)
		os.Exit(0)
	}()

	// Start
	ticker := time.NewTicker(cli.Interval)
	baseString := cli.PrintString
	if cli.PrintName {
		baseString = kctx.Model.Name + " " + baseString
	}

	go func() {
		for t := range ticker.C {
			if cli.PrintTimestamp {
				_, _ = io.WriteString(os.Stdout, fmt.Sprintf("%s %s", baseString, t.Format(time.RFC1123)))
			} else {
				_, _ = io.WriteString(os.Stdout, baseString)
			}
			if cli.PrintNewline {
				_, _ = io.WriteString(os.Stdout, "\n")
			}
		}
	}()
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), cli.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, cli.Command[0], cli.Command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		log.Fatal().Msgf("Timeout reached. Terminating %s", cli.Command[0])
	} else if err != nil {
		log.Fatal().Err(err).Msgf("Non-zero exit code for %s", cli.Command[0])
	}
}
