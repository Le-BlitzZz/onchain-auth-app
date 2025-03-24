package config

import "github.com/urfave/cli/v2"

var Flags = CliFlags{
	{
		Flag: &cli.StringFlag{
			Name:    "http-host",
			Aliases: []string{"ip"},
			Value:   "0.0.0.0",
			Usage:   "Web server `IP` address",
			EnvVars: []string{EnvVar("HTTP_HOST")},
		},
	},
	{
		Flag: &cli.IntFlag{
			Name:    "http-port",
			Aliases: []string{"port"},
			Value:   8080,
			Usage:   "Web server port `NUMBER`",
			EnvVars: []string{EnvVar("HTTP_PORT")},
		},
	},
}
