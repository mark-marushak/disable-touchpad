package dispatcher

import (
	"context"
	"fmt"
	"time"
)

type ITouchpad interface {
	Off() error
	On() error
	Status() bool
}

type Dispatcher struct {
	touchpad ITouchpad
	deadline time.Time
}

func NewDispatcher(touchpad ITouchpad) *Dispatcher {
	return &Dispatcher{
		touchpad: touchpad,
		deadline: time.Now(),
	}
}

func (d *Dispatcher) Disable() error {
	if d.touchpad.Status() {
		if err := d.touchpad.Off(); err != nil {
			return fmt.Errorf("touchpad off: %w", err)
		}
	}

	// update deadline when got a new event
	d.deadline = time.Now().Add(time.Second)

	return nil
}

func (d *Dispatcher) Watch(ctx context.Context) error {
	ticker := time.NewTicker(time.Millisecond * 100)
	for range ticker.C {
		select {
		case <-ctx.Done():
			return nil
		default:
			if time.Now().After(d.deadline) {
				if !d.touchpad.Status() {
					if err := d.touchpad.On(); err != nil {
						return fmt.Errorf("touchpad on: %w", err)
					}
				}
			}
		}
	}

	return nil
}

func (d *Dispatcher) On() error {
	return d.touchpad.On()
}

func (d *Dispatcher) Off() error {
	return d.touchpad.Off()
}
