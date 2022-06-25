package exiterrorf

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// ExitErrorf enables standardized error message formatting.
func ExitErrorf(err error, args ...string) { // lint:allow_dead
	if len(args) == 0 {
		args = []string{""}
	}

	err = Errorf(err, args[0])
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

// Errorf enables standardized error message formatting.
func Errorf(err error, args ...string) error { // lint:allow_dead
	function, file, line, _ := runtime.Caller(2)
	meta := fmt.Sprintf(
		"[%s:%d %s]\n",
		chopPath(file),
		line,
		runtime.FuncForPC(function).Name(),
	)

	if len(args) == 0 {
		args[0] = ""
	}

	return fmt.Errorf("[ERROR]: %w %s", errors.Wrap(err, args[0]), meta)
}

func chopPath(original string) string { // lint:allow_dead
	i := strings.LastIndex(original, "/")

	if i == -1 {
		return original
	}

	return original[i+1:]
}
