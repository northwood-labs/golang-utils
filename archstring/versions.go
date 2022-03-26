package archstring

import (
	"fmt"
)

// GetFriendlyName takes the values of GOOS and GOARCH, and returns a "friendly"
// name for the pairing (U.S. English).
func GetFriendlyName(osStr, archStr string) string {
	var osFriendly string
	var archFriendly string

	osFriendly = osStr
	archFriendly = archStr

	if osStr == "js" && archStr == "wasm" {
		return "WebAssembly"
	}

	if osStr == "darwin" && archStr == "arm64" {
		return "macOS on Apple Silicon"
	}

	if val, ok := OSMap[osStr]; ok {
		osFriendly = val
	}

	if val, ok := ArchMap[archStr]; ok {
		archFriendly = val
	}

	return fmt.Sprintf("%s on %s", osFriendly, archFriendly)
}
