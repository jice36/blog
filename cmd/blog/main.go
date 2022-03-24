package main

import (
	"flag"
	"log"

	"github.com/jice36/blog/internal/server"
	"github.com/jice36/blog/internal/service"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yml", "path to config file") //todo пароль поправить в конфиге
}

func main() {
	flag.Parse()

	config, cErr := server.NewConfigBlog(configPath)
	if cErr != nil {
		log.Fatal(cErr)
	}

	service := service.NewService()
	defer service.Conn.Close()

	s := server.NerServer(config, service)

	sErr := s.StartServer()
	if sErr != nil {
		s.Log.Fatal(sErr)
	}
}

