package config

import "github.com/urfave/cli/v2"

var Flags = CliFlags{
	{
		Flag: &cli.StringFlag{
			Name:    "default-user",
			Aliases: []string{"login"},
			Value:   "default",
			Usage:   "`USERNAME` of the default user that is created on first startup",
			EnvVars: EnvVars("DEFAULT_USER"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "default-password",
			Aliases: []string{"pwd"},
			Usage:   "initial `PASSWORD` of the default user that is created on first startup",
			EnvVars: EnvVars("DEFAULT_PASSWORD"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "http-host",
			Aliases: []string{"ip"},
			Value:   "0.0.0.0",
			Usage:   "Web server `IP` address",
			EnvVars: EnvVars("HTTP_HOST"),
		},
	},
	{
		Flag: &cli.IntFlag{
			Name:    "http-port",
			Aliases: []string{"port"},
			Value:   8080,
			Usage:   "Web server port `NUMBER`",
			EnvVars: EnvVars("HTTP_PORT"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "database-driver",
			Aliases: []string{"db"},
			Value:   "mysql",
			Usage:   "database `DRIVER`",
			EnvVars: EnvVars("DATABASE_DRIVER"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "database-dsn",
			Aliases: []string{"dsn"},
			Usage:   "database connection `DSN`",
			EnvVars: EnvVars("DATABASE_DSN"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "database-name",
			Aliases: []string{"db-name"},
			Value:   "authonchain",
			Usage:   "database schema `NAME`",
			EnvVars: EnvVars("DATABASE_NAME"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "database-server",
			Aliases: []string{"db-server"},
			Usage:   "database `HOST` incl. port e.g. \"mariadb:3306\"",
			EnvVars: EnvVars("DATABASE_SERVER"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "database-user",
			Aliases: []string{"db-user"},
			Value:   "photoprism",
			Usage:   "database user `NAME`",
			EnvVars: EnvVars("DATABASE_USER"),
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "database-password",
			Aliases: []string{"db-pass"},
			Usage:   "database user `PASSWORD`",
			EnvVars: EnvVars("DATABASE_PASSWORD"),
		},
	},
	{
		Flag: &cli.IntFlag{
			Name:    "database-timeout",
			Usage:   "timeout in `SECONDS` for establishing a database connection (1-60)",
			EnvVars: EnvVars("DATABASE_TIMEOUT"),
			Value:   15,
		},
	},
	{
		Flag: &cli.StringFlag{
			Name:    "assets-path",
			Aliases: []string{"as"},
			Usage:   "assets `PATH` containing static resources",
			EnvVars: EnvVars("ASSETS_PATH"),
		},
	},
}
