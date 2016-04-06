package main

import (
	"encoding/json"
	"os/exec"
)

func run(cmd string, args []string) ([]byte, error) {
	return exec.Command(cmd, args...).Output()
}

func salt(out interface{}, args ...string) (err error) {
	args = append([]string{"--out", "json"}, args...)

	data, err := run("salt", args)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		return err
	}

	return nil
}
