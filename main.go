package main

import "github.com/tucnak/climax"

func main() {
	truck := climax.New("truck")
	truck.Brief = "truck is a helper utility for salt"
	truck.Version = "pre"

	truck.AddCommand(sshCommand)
	truck.Run()
}
