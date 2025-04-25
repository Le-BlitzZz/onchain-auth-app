package config

import (
	"reflect"

	"github.com/urfave/cli/v2"
)

// Options hold the global configuration values without further validation or processing.
// Application code should retrieve option values via getter functions since they provide
// validation and return defaults if a value is empty.
type Options struct {
	DefaultUser      string `flag:"default-user"`
	DefaultPassword  string `flag:"default-password"`
	AssetsPath       string `flag:"assets-path"`
	HttpHost         string `flag:"http-host"`
	HttpPort         int    `flag:"http-port"`
	DatabaseDriver   string `flag:"database-driver"`
	DatabaseDsn      string `flag:"database-dsn"`
	DatabaseName     string `flag:"database-name"`
	DatabaseServer   string `flag:"database-server"`
	DatabaseUser     string `flag:"database-user"`
	DatabasePassword string `flag:"database-password"`
	DatabaseTimeout  int    `flag:"database-timeout"`
}

// NewOptions creates a new configuration entity by using ApplyCliContext
// to initialize options through the CLI.
func NewOptions(ctx *cli.Context) *Options {
	o := &Options{}
	o.ApplyCliContext(ctx)
	return o
}

// ApplyCliContext uses options from the CLI to initialize configuration for the entity.
func (o *Options) ApplyCliContext(ctx *cli.Context) {
	v := reflect.ValueOf(o).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)

		tagValue := v.Type().Field(i).Tag.Get("flag")

		switch t := fieldValue.Interface().(type) {
		case string:
			if ctx.IsSet(tagValue) || fieldValue.String() == "" {
				f := ctx.String(tagValue)
				fieldValue.SetString(f)
			}
		case int:
			if ctx.IsSet(tagValue) || fieldValue.Int() == 0 {
				f := ctx.Int64(tagValue)
				fieldValue.SetInt(f)
			}
		default:
			log.Warnf("cannot assign value of type %s from cli flag %s", t, tagValue)
		}

	}
}
