package constants

import "fmt"

const (
	major = 0
	minor = 1
	path  = 1
)

var Version = fmt.Sprint(major) + "." + fmt.Sprint(minor) + "." + fmt.Sprint(path) + "-ALPHA"
