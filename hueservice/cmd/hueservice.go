package main

import (
	"github.com/spf13/viper"
	"gitlab.com/scotthome/hueservice/controller"
	router2 "gitlab.com/scotthome/hueservice/router"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	viper.SetEnvPrefix("scott")
	viper.AutomaticEnv()
	zap.S().Infow("Setting up lights controller...")
	c, err := controller.NewDefaultLightsController()
	if err != nil {
		zap.S().Fatalw("failed to initialize lights controller", "error", err)
	}
	zap.S().Infow("Set up lights controller!")
	r, err := router2.NewHTTPRouter(c)
	if err != nil {
		zap.S().Fatalw("failed to initialize lights controller", "error", err)
	}
	zap.S().Infow("Serving...")
	err = r.Serve()
	if err != nil {
		zap.S().Fatalw("failed to initialize lights controller", "error", err)
	}
}
