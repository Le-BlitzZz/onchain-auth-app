package config

import "github.com/urfave/cli/v2"

var Flags = CliFlags{
	{
		Flag: &cli.StringFlag{
			Name:    "default-user",
			Aliases: []string{"login"},
			Value:   "default",
			Usage:   "`USERNAME` of the default user that is created on first startup",
			EnvVars: []string{EnvVar("DEFAULT_USER")},
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "default-password",
			Aliases: []string{"pwd"},
			Usage:   "initial `PASSWORD` of the default user that is created on first startup",
			EnvVars: []string{EnvVar("DEFAULT_PASSWORD")},
		},
	},
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
	{
		Flag: &cli.StringFlag{
			Name:    "assets-path",
			Aliases: []string{"as"},
			Usage:   "assets `PATH` containing static resources",
			EnvVars: []string{EnvVar("ASSETS_PATH")},
		},
	},
}
