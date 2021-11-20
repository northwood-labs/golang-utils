package debug

import spew "github.com/davecgh/go-spew/spew"

// GetSpew is analogous to PrettyPrint in Python or print_r() in PHP.
func GetSpew() spew.ConfigState { // lint:allow_dead
	return spew.ConfigState{
		Indent:   "    ",
		SortKeys: true,
		SpewKeys: true,
	}
}
