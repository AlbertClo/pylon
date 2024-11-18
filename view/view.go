package view

import (
	"github.com/AlbertClo/pylon/keybind"
	"github.com/AlbertClo/pylon/types"
)

type Model interface {
	GetCounter() int
	GetMessage() string
	GetKeys() keybind.Shortcuts
	GetMenuItems() []types.MenuItem
	GetSelectedItem() int
}
