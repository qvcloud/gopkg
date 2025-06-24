package version

import "fmt"

var (
	Version = "v0.0.0"
	Go      = ""
	Commit  = ""
	Build   = "0000-00-00"
)

func ShowVersion() {
	fmt.Println("version:", Version)
	fmt.Println("go:", Go)
	fmt.Println("commit:", Commit)
	fmt.Println("build:", Build)
}
