package main

import (
	"fmt"
	"os"
	"os/signal"
	"payment-simulator/core"
	"payment-simulator/internal"
	"payment-simulator/utils"
	"runtime"
	"strconv"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	zap.S().Info("initiate configuration service")

	path, err := utils.GetPathNow()
	if err != nil {
		fmt.Println("error when get working directory")
	}
	logPath := utils.GetEnv("LOG_PATH", path+"/log/testlog.log")

	internal.Initiate(logPath, 1, 3, true)

	maxProc, _ := strconv.Atoi(utils.GetEnv("MAX_PROCESSOR", "1"))
	runtime.GOMAXPROCS(maxProc)
	// fmt.Println("2")

	var errChan = make(chan error, 1)
	Route := core.InitRouter()
	// fmt.Println("3")

	go func() {
		listenAddress := utils.GetEnv("LISTEN_ADDRESS", "0.0.0.0:8110")
		zap.S().Infow("initiate configuration service")

		// otelzap.Ctx(context.TODO()).Info("Starting @:" + listenAddress)

		errChan <- core.GinServerUp(listenAddress, Route.SetRouter())
		// fmt.Println("error while running api, exiting...", errChan)
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		// fmt.Println("got an interrupt, exiting...")
		zap.S().Infow("got an interrupt, exiting...")

		// sugarLogger.Error("Got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			// fmt.Println("error while running api, exiting...", err)
			zap.S().Infow("got an interrupt, exiting...", err)

			// sugarLogger.Error("Error while running api, exiting... ", zap.Error(err))
		}
	}
}
