package view

import (
	"github.com/AlbertClo/pylon/keybind"
)

type Model interface {
	GetCounter() int
	GetMessage() string
	GetKeys() keybind.Shortcuts
	GetMenuItems() []string
	GetSelectedItem() int
}
