package cmd

import (
	"flag"
)

func Execute() *Commands {
	// command results struct
	var commands Commands

	// the commands
	commands.Nuke = flag.Bool("nuke", false, "drop tables and recreate db but dont serve")
	commands.Migrate = flag.Bool("migrate", false, "migrate db but dont serve")
	commands.Migrate_serve = flag.Bool("migrate-serve", false, "migrate db and serve")

	flag.Parse()

	return &commands
}

// a struct of the results
type Commands struct {
	Migrate       *bool
	Nuke          *bool
	Migrate_serve *bool
}
