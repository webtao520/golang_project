package golog

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Color int

const (
	NoColor Color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Purple
	DarkGreen
	White
)

var logColorPrefix = []string{
	"",
	"\x1b[030m",
	"\x1b[031m",
	"\x1b[032m",
	"\x1b[033m",
	"\x1b[034m",
	"\x1b[035m",
	"\x1b[036m",
	"\x1b[037m",
}

var colorByName = map[string]Color{
	"none":      NoColor,
	"black":     Black,
	"red":       Red,
	"green":     Green,
	"yellow":    Yellow,
	"blue":      Blue,
	"purple":    Purple,
	"darkgreen": DarkGreen,
	"white":     White,
}

func matchColor(name string) Color {

	lower := strings.ToLower(name)

	for cname, c := range colorByName {

		if cname == lower {
			return c
		}
	}

	return NoColor
}

func ColorFromLevel(l Level) Color {
	switch l {
	case Level_Warn:
		return Yellow
	case Level_Error, Level_Fatal:
		return Red
	}

	return NoColor
}

var logColorSuffix = "\x1b[0m"

type ColorMatch struct {
	Text  string
	Color string

	c Color
}

type ColorFile struct {
	Rule []*ColorMatch
}

func (self *ColorFile) ColorFromText(text string) Color {

	for _, rule := range self.Rule {
		if strings.Contains(text, rule.Text) {
			return rule.c
		}
	}

	return NoColor
}

func (self *ColorFile) Load(data string) error {

	err := json.Unmarshal([]byte(data), self)
	if err != nil {
		return err
	}

	for _, rule := range self.Rule {

		rule.c = matchColor(rule.Color)

		if rule.c == NoColor {
			return fmt.Errorf("color name not exists: %s", rule.Text)
		}

	}

	return nil
}

func NewColorFile() *ColorFile {
	return &ColorFile{}
}
