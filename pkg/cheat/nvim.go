package cheat

import "github.com/charmbracelet/bubbles/table"

var nvimColumns = []table.Column{
	{Title: "Mode", Width: 10},
	{Title: "Description", Width: 30},
	{Title: "Command", Width: 30},
}

var nvimRows = []table.Row{
	{"None", "Start a new session", "nvim"},
	{"None", "Command mode", "shift :"},
	{"Command", "Quit", ":q"},
	{"Command", "Quit without saving", ":q!"},
	{"Command", "Save and quit", ":wq"},
	{"Command", "Save", ":w"},
	{"Command", "Save as", ":w [filename]"},
}
