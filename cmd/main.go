package main

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version  = "dev"
	timeout  time.Duration
	interval time.Duration
	fullcmd  []string
)

func main() {
	// Parse command line
	kingpin.Flag("timeout", "Timeout for this command.").Default("20m").DurationVar(&timeout)
	kingpin.Flag("interval", "Interval at which to print keep-alive messages.").Default("1m").DurationVar(&interval)
	kingpin.Arg("command", "Command to execute.").Required().StringsVar(&fullcmd)
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(version).Author("CrazyMax")
	kingpin.CommandLine.Name = "travis-wait-enhanced"
	kingpin.CommandLine.Help = `Prevent Travis CI from thinking a long-running process has stalled. More info: https://github.com/crazy-max/travis-wait-enhanced`
	kingpin.Parse()

	// Logger
	output := zerolog.ConsoleWriter{
		Out: os.Stdout,
	}
	output.FormatTimestamp = func(i interface{}) string {
		return kingpin.CommandLine.Name
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
	ticker := time.NewTicker(interval)
	go func() {
		for t := range ticker.C {
			log.Info().Msgf("Still running at %s...", t.Format(time.RFC1123))
		}
	}()
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, fullcmd[0], fullcmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		log.Fatal().Msgf("Timeout reached. Terminating %s", fullcmd[0])
	} else if err != nil {
		log.Fatal().Err(err).Msgf("Non-zero exit code for %s", fullcmd[0])
	}
}
