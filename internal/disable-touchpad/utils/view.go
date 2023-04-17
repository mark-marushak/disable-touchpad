package utils

import (
	"errors"
	"fmt"
	"strings"

	evdev "github.com/gvalkov/golang-evdev"
)

// Select a device from a list of accessible input devices.
//
// PathGlob string
// For linux it will: /dev/input/event*
func SelectDevice(pathGlob string) (*evdev.InputDevice, error) {
	devices, _ := evdev.ListInputDevices(pathGlob)

	lines := make([]string, 0)
	max := 0
	if len(devices) > 0 {
		for i := range devices {
			dev := devices[i]
			str := fmt.Sprintf("%-3d %-20s %-35s %s", i, dev.Fn, dev.Name, dev.Phys)
			if len(str) > max {
				max = len(str)
			}
			lines = append(lines, str)
		}
		fmt.Printf("%-3s %-20s %-35s %s\n", "ID", "Device", "Name", "Phys")
		fmt.Printf(strings.Repeat("-", max) + "\n")
		fmt.Printf(strings.Join(lines, "\n") + "\n")

		var choice int
		choice_max := len(lines) - 1

	ReadChoice:
		fmt.Printf("Select device [0-%d]: ", choice_max)
		_, err := fmt.Scan(&choice)
		if err != nil || choice > choice_max || choice < 0 {
			goto ReadChoice
		}

		return devices[choice], nil
	}

	errmsg := fmt.Sprintf("no accessible input devices found by %s", pathGlob)
	return nil, errors.New(errmsg)
}

func FormatEvent(ev *evdev.InputEvent) string {
	var res, f, code_name string

	code := int(ev.Code)
	etype := int(ev.Type)

	switch ev.Type {
	case evdev.EV_SYN:
		if ev.Code == evdev.SYN_MT_REPORT {
			f = "time %d.%-8d +++++++++ %s ++++++++"
		} else {
			f = "time %d.%-8d --------- %s --------"
		}
		return fmt.Sprintf(f, ev.Time.Sec, ev.Time.Usec, evdev.SYN[code])
	case evdev.EV_KEY:
		val, haskey := evdev.KEY[code]
		if haskey {
			code_name = val
		} else {
			val, haskey := evdev.BTN[code]
			if haskey {
				code_name = val
			} else {
				code_name = "?"
			}
		}
	default:
		m, haskey := evdev.ByEventType[etype]
		if haskey {
			code_name = m[code]
		} else {
			code_name = "?"
		}
	}

	evfmt := "time %d.%-8d type %d (%s), code %-3d (%s), value %d"
	res = fmt.Sprintf(evfmt, ev.Time.Sec, ev.Time.Usec, etype,
		evdev.EV[int(ev.Type)], ev.Code, code_name, ev.Value)

	return res
}
