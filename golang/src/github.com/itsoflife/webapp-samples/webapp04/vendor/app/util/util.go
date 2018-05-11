// Convenience function for the project.
package util

import (
  "app/logs"  // logger must be initialized at program startup
)

func CheckError(err error) {
    if err != nil {
        logs.Logger.Error(err)
    }
}


