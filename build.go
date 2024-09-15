package build

import "log"

var (
	// The build timestamp, e.g: Wed Jul  6 14:29:49 +07 2022
	Time string

	// The build user, e.g: Foo
	User string

	// The build number, e.g: 20220802
	Number string

	// The build git source branch, e.g: master
	Source string

	// The build git commit sha
	Commit string

	// The build version, e.g: v1.0
	Version string
)

func Print() {
	log.Println("- build.Time:", Time)
	log.Println("- build.User:", User)
	log.Println("- build.Number:", Number)
	log.Println("- build.Source:", Source)
	log.Println("- build.Commit:", Commit)
	log.Println("- build.Version:", Version)
}
