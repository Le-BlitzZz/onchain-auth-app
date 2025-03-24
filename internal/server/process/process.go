package process

import "os"

// ID is the process ID under which the server is running.
var ID = os.Getpid()

var Signal = make(chan os.Signal)
