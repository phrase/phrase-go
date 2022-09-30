//go:build linux || darwin
// +build linux darwin

package phrase

import (
	"os"
)

func defaultConfigDir() string {
	return os.Getenv("HomePath")
}
