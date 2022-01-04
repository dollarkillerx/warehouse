package main

import (
	"github.com/dollarkillerx/warehouse/internal/config"
	"github.com/dollarkillerx/warehouse/internal/server"
	"github.com/dollarkillerx/warehouse/pkg/utils"

	"fmt"
	"log"
)

func main() {
	utils.InitLogger(config.GetLoggerConfig())

	s := server.New()
	fmt.Println("run: ", config.GetListenAddr())
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}
