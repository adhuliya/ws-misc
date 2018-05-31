// Convenience function for the project.
package util

import (
	"app/logs" // logger must be initialized at program startup
	"os"
)

func CheckError(err error) {
	if err != nil {
		logs.Logger.Error(err)
	}
}

// will force the program to exit
func FatalError(err error) {
	if err != nil {
		logs.Logger.Error(err)
		os.Exit(1)
	}
}
