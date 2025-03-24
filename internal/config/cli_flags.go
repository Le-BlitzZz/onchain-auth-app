package config

import "github.com/urfave/cli/v2"

type CliFlags []CliFlag

func (f CliFlags) Cli() (result []cli.Flag) {
	result = make([]cli.Flag, 0, len(f))

	for _, flag := range f {
		result = append(result, flag.Flag)
	}

	return result
}
