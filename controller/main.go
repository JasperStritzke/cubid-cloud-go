package main

import (
	"flag"
	"fmt"
	"github.com/AlexanderGrom/go-event"
	"github.com/JasperStritzke/cubid-cloud/controller/event_names"
	"github.com/JasperStritzke/cubid-cloud/controller/executor"
	"github.com/JasperStritzke/cubid-cloud/controller/rest"
	"github.com/JasperStritzke/cubid-cloud/controller/server_group"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/mathutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Controller struct {
	stop *chan int
}

func main() {
	startup := time.Now()

	nologs := flag.Bool("nologs", false, "use to disable logging")
	_ = flag.Bool("freshstart", false, "use to start the system with a blank configuration")

	flag.Parse()

	if !*nologs {
		logger.ActivateLogs()
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	rest.InitRouter()
	_ = executor.NewManager()
	_ = server_group.NewGroupManager()

	logger.Info("Started. Press CTRL+C (^C) to stop.")
	// hide input because this application doesn't accept console input
	fmt.Print("\033[8m")

	<-quit

	fmt.Println(" - SIGTERM called")
	_ = event.Go(event_names.Shutdown)

	logger.Warn("Shutting down...")

	//TODO Shutdown logic

	logger.Info("Successfully stopped system.")

	logger.Info("System was up " + mathutil.RoundAndString(time.Now().Sub(startup).Hours(), 2) + " hours")
}
