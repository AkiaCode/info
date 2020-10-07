package info

import (
	"github.com/fatih/color"
)

func log(msg string) {
	color.Green(msg)
}

func debug(msg string) {
	color.Black(msg)
}

func error(msg string) {
	color.RedString(msg)
}

func warn(msg string) {
	color.Red(msg)
}

func success(msg string) {
	color.GreenString(msg)
}
