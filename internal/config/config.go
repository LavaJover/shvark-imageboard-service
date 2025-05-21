package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type ImageboardConfig struct {
	Env string 		`yaml:"env"`
	GRPCServer 		`yaml:"grpc_server"`
	ImageboardDB 	`yaml:"Imageboard_db"`
	LogConfig 		`yaml:"log_config"`
}

type GRPCServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ImageboardDB struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	LogLevel 	string 	`yaml:"log_level"`
	LogFormat 	string 	`yaml:"log_format"`
	LogOutput 	string 	`yaml:"log_output"`
}

func MustLoad() *ImageboardConfig {

	// Processing env config variable and file
	configPath := os.Getenv("Imageboard_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("Imageboard_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg ImageboardConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}