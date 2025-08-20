// main.go
package main

import (
	"github.com/diiineeei/mvc-sample/pkg/config/logger"
	"github.com/diiineeei/mvc-sample/pkg/router"
	"go.uber.org/zap"
)

func main() {
	// Inicializa o logger
	log, err := logger.NewLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer log.Sync()

	// Inicializa o roteador
	r := router.SetupRouter(log)

	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}

}
