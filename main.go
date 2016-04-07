package main

import "github.com/tucnak/climax"

//go:generate sh version.sh

var ver = "pre"

func main() {
	truck := climax.New("truck")
	truck.Brief = "truck is a helper utility for salt"
	truck.Version = ver

	truck.AddCommand(sshCommand)
	truck.Run()
}
