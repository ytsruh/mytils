package tmux

import "github.com/charmbracelet/bubbles/table"

var columns = []table.Column{
	{Title: "Description", Width: 30},
	{Title: "Command", Width: 30},
}

var rows = []table.Row{
	{"Start a new session", "tmux"},
	{"Start a new named session", "tmux new -s [name]"},
	{"List tmux sessions", "tmux ls"},
	{"Detach from current session", "tmux d"},
	{"Rename current session", "tmux $"},
	{"Show all tmux", "tmux info"},
	{"Command mode", "ctrl+b, :"},
	{"Exit command mode", "ctrl+b, ctrl+c"},
	{"Help", "ctrl+b, ?"},
	{"Switch window", "ctrl+b, [number]"},
	{"Close current window", "ctrl+b, &"},
	{"Split window horizontally", "ctrl+b, %"},
	{"Split window vertically", "ctrl+b, \""},
	{"Close current pane", "ctrl+b, x"},
	{"Switch pane", "ctrl+b, [up/down/left/right]"},
	{"Show pane numbers", "ctrl+b, q"},
}