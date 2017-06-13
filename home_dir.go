package cgiutil

import (
	"fmt"
	"os"
	"strings"
)

// GetHomeDir returns the home directory calculated from
// the SCRIPT_FILENAME environment variable.
// It assumes the script filename is like /home/xxx/yyy/.../zzz.cgi
// and the home directory is like /home/xxxx.
// It panics if this assumption is not correct.
func GetHomeDir() string {
	scriptFilename := os.Getenv("SCRIPT_FILENAME")
	if !strings.HasPrefix(scriptFilename, "/home/") {
		panic(fmt.Sprintf("scriptFilename %q does not start with \"/home/\"", scriptFilename))
	}
	i := strings.IndexRune(scriptFilename[len("/home/"):], '/')
	if i == -1 {
		panic(fmt.Sprintf("scriptFilename %q does not contain '/' after \"/home/\"", scriptFilename))
	}
	return scriptFilename[:len("/home/")+i]
}
