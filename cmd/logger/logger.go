package cmd

import (
	color "github.com/fatih/color"
)

// Initialize a nw logger struct to interact w
func NewLogger() *Logger {
	return &Logger{
		Text:   "",
		Status: 0,
		Color:  "",
	}
}

// Set new values to the logger struct.
func (log *Logger) Set(text string, status int, color string) *Logger {
	log.Text = text
	log.Status = status
	log.Color = color
	return log
}

// In case there is an error use this method to log the text as error message in red color.
func (log *Logger) Error(Message string, Error error) *Logger {
	if Error != nil {
		log.Text = "|-| " + Message + Error.Error()
	} else {
		// The error included in the message
		log.Text = "|-| " + Message

	}
	log.Status = 0
	log.Color = "red"
	return log.log()
}

// Success stage, log a message as success process.
func (log *Logger) Success(text string) *Logger {
	log.Text = "|+| " + text
	log.Status = 1
	log.Color = "green"
	return log.log()
}

// State stage, log a message as state inside a process.
func (log *Logger) State(text string) *Logger {
	log.Text = "|+| " + text
	log.Status = 1
	log.Color = "white"
	return log.log()
}

/*
- Function log to log out a text with an exact color.
- Available colors [red, green, blue]
*/
func (lg *Logger) log() *Logger {
	c := color.New(color.FgWhite) // Default color is white

	switch lg.Color {
	case "red":
		c = color.New(color.FgRed)
	case "green":
		c = color.New(color.FgGreen)
	case "white":
		c = color.New(color.FgWhite)
	case "blue":
		c = color.New(color.FgBlue)
	}
	c.Println(lg.Text)
	return lg
}
