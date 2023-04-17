package syclient

import (
	"fmt"
	"os/exec"
)

type Touchpad struct {
	active bool
}

func (t *Touchpad) Off() error {
	cmd := exec.Command("synclient", "TouchpadOff=1")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("syclient: %w", err)
	}

	t.active = false

	return nil
}

func (t *Touchpad) On() error {
	cmd := exec.Command("synclient", "TouchpadOff=0")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("syclient: %w", err)
	}

	t.active = true

	return nil
}

func (t *Touchpad) Status() bool {
	return t.active
}
