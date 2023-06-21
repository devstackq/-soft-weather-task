package main

import (
	"fmt"
	"log"
	"net/http"
	"soft-weater/internal/clients/http/service-1"
	"soft-weater/internal/config"
	"soft-weater/internal/database"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/managers"
	v1 "soft-weater/internal/resources/http/v1"
)

func Run() error {

	opts := config.LoadConfig()

	ds, err := database.New(drivers.DataStoreConfig{
		URL:           opts.DSURL,
		DataStoreName: opts.DSDB,
		DataBaseName:  opts.DSName,
	})
	if err != nil {
		return err
	}

	if err = ds.Connect(opts.DataStoreConfiguration); err != nil {
		return fmt.Errorf("cannot connect to datastore: %s", err)
	}

	taskMan := managers.NewTask(ds)
	solveMan := managers.NewSolver(
		ds.HistoryRepository(),
		ds.TaskRepository(),
		service.NewService(),
	)

	server := v1.NewServer(opts, taskMan, solveMan)

	return http.ListenAndServe(opts.ServerConfiguration.Port, server.Router)

}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
}
