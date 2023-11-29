package config

import "github.com/alecthomas/kingpin"

type Config struct {
	Port string
	File string
	Ext  string
}

func NewConfig() *Config {
	var cfg = Config{}

	kingpin.Flag("port", "Port for the application to listen on, e.g., \"8080\"").Default("8080").Short('p').StringVar(&cfg.Port)
	kingpin.Flag("file", "Path to the configuration file, e.g., \"/zip/test.zip\"").Short('f').Required().StringVar(&cfg.File)
	kingpin.Flag("ext", "file extension of the configuration file, e.g., \".c\"").Short('e').Default("").StringVar(&cfg.Ext)

	kingpin.Parse()

	return &cfg

}
