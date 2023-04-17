//go:build linux
// +build linux

// Input device event monitor.
package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	dispatcher2 "disable-touchpad/internal/disable-touchpad/dispatcher"
	"disable-touchpad/internal/disable-touchpad/respository/syclient"

	evdev "github.com/gvalkov/golang-evdev"
	cli "github.com/urfave/cli/v2"
)

const touchpadEvent = "/dev/input/event4"

func main() {
	// init new CLI app
	app := cli.NewApp()
	app.Name = "Touchpad dispatcher"
	app.Usage = "When the touchpad dispatcher run it controls you keyboard event to disable touchpad when you typing and enable after 1 second by default"
	app.Action = action
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "path",
			Value: "/dev/input/event*",
			Usage: "Path to folder with events",
		},
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	if err := app.RunContext(ctx, nil); err != nil {
		log.Fatalf("Touchpad dispatcher run: %v", err)
		return
	}
}

func action(c *cli.Context) error {
	var dev *evdev.InputDevice
	var events []evdev.InputEvent
	var err error

	dev, err = evdev.Open(touchpadEvent)
	if err != nil {
		log.Fatal(err)
	}

	dispatcher := dispatcher2.NewDispatcher(&syclient.Touchpad{})
	if err = dispatcher.On(); err != nil {
		return fmt.Errorf("make sure touchpad is enabled: %w", err)
	}

	go func() {
		err = dispatcher.Watch(c.Context)
		if err != nil {
			log.Fatalf("dispatcher watch: %v", err)
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 50)
	for range ticker.C {
		select {
		case <-c.Context.Done():
			return nil
		default:
			events, err = dev.Read()
			if err != nil {
				return fmt.Errorf("read event: %w", err)
			}

			if len(events) > 0 {
				if err = dispatcher.Disable(); err != nil {
					return fmt.Errorf("dipatcher disable: %v", err)
				}
			}
		}

	}

	return err
}
