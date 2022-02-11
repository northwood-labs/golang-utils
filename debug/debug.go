package debug

import (
	"os"
	"strings"

	spew "github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GetSpew is analogous to PrettyPrint in Python or print_r() in PHP.
func GetSpew() spew.ConfigState { // lint:allow_dead
	return spew.ConfigState{
		Indent:   "    ",
		SortKeys: true,
		SpewKeys: true,
	}
}

// Vardump prints the name of the variable and dumps its contents.
func Vardump(varname string, args ...interface{}) {
	pp := GetSpew()

	for i := range args {
		arg := args[i]

		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})

		log.Debug().
			Str("variable", varname).
			Msg(
				strings.TrimSpace(
					pp.Sdump(arg),
				),
			)
	}
}
