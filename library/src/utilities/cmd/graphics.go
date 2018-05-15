package cmd

import (
	"time"
)

func Loading() {

	for i := 0; i < 5; i++ {
		time.Sleep(220 * time.Millisecond)
		print("\r|")
		time.Sleep(220 * time.Millisecond)
		print("\r/")
		time.Sleep(220 * time.Millisecond)
		print("\r-")
		time.Sleep(220 * time.Millisecond)
		print("\r\\")
		time.Sleep(220 * time.Millisecond)
		print("\r|")
		time.Sleep(220 * time.Millisecond)
		print("\r/")
		time.Sleep(220 * time.Millisecond)
		print("\r-")
		time.Sleep(220 * time.Millisecond)
		print("\r\\")
	}
	time.Sleep(220 * time.Millisecond)
	print("\r")
}

func ChangeColor(text, color string) string {
	var text_with_color = ""
	switch color {
	case "black":
		text_with_color = "\x1b[30;1m" + text + "\x1b[0m"
	case "red":
		text_with_color = "\x1b[31;1m" + text + "\x1b[0m"
	case "green":
		text_with_color = "\x1b[32;1m" + text + "\x1b[0m" // green
	case "yellow":
		text_with_color = "\x1b[33;1m" + text + "\x1b[0m" // yellow
	case "blue":
		text_with_color = "\x1b[34;1m" + text + "\x1b[0m" // blue
	case "magenta":
		text_with_color = "\x1b[35;1m" + text + "\x1b[0m" // magenta
	case "cyan":
		text_with_color = "\x1b[36;1m" + text + "\x1b[0m" // cyan
	case "white":
		text_with_color = "\x1b[37;1m" + text + "\x1b[0m" // white
	case "brightBlack":
		text_with_color = "\x1b[90;1m" + text + "\x1b[0m" // brightBlack
	case "brightRed":
		text_with_color = "\x1b[91;1m" + text + "\x1b[0m" // brightRed
	case "brightGreen":
		text_with_color = "\x1b[92;1m" + text + "\x1b[0m" // brightGreen
	case "brightYellow":
		text_with_color = "\x1b[93;1m" + text + "\x1b[0m" // brightYellow
	case "brightBlue":
		text_with_color = "\x1b[94;1m" + text + "\x1b[0m" // brightBlue
	case "brightMagenta":
		text_with_color = "\x1b[95;1m" + text + "\x1b[0m" // brightMagenta
	case "brightCyan":
		text_with_color = "\x1b[96;1m" + text + "\x1b[0m" // brightCyan
	case "brightWhite":
		text_with_color = "\x1b[97;1m" + text + "\x1b[0m" // brightWhite
	default:
		text_with_color = "\x1b[34;1m" + text + "\x1b[0m"
	}
	return text_with_color
}
