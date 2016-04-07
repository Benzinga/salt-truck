package main

import (
	"os/exec"

	"gopkg.in/yaml.v2"
)

func run(cmd string, args []string) ([]byte, error) {
	return exec.Command(cmd, args...).Output()
}

func salt(out interface{}, args ...string) (err error) {
	args = append([]string{"--out", "yaml"}, args...)

	data, err := run("salt", args)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, out)
	if err != nil {
		return err
	}

	return nil
}
