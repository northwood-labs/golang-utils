package exiterrorf

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ExitErrorf enables standardized error message formatting.
func ExitErrorf(args ...interface{}) { // lint:allow_dead
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "Error: %+v\n\n", args...)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %+v\nContext: %+v\n\n", args...)
	}

	function, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "%s:%d %s\r\n", chopPath(file), line, runtime.FuncForPC(function).Name())

	os.Exit(1)
}

func chopPath(original string) string { // lint:allow_dead
	i := strings.LastIndex(original, "/")

	if i == -1 {
		return original
	}

	return original[i+1:]
}
