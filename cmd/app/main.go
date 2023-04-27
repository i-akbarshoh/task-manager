package main

import (
	"fmt"
	"github.com/i-akbarshoh/task-manager/pkg/config"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/datastore"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/router"
	"github.com/i-akbarshoh/task-manager/pkg/usecase/controller"
	"github.com/i-akbarshoh/task-manager/pkg/usecase/repository"
	"net/http"
	"os"
	"sync"
)

func main() {
	//load config
	config.ReadConfig()

	// postgres database initialized and migrate
	db := datastore.NewDB()
	defer db.Close()

	//router and repository initialized
	repo := repository.NewPostgres(db)
	c := controller.New(repo)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := router.NewRouter(c)
		if err := http.ListenAndServe(":"+config.C.Server.Port, r); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	fmt.Println("Program is running on", config.C.Server.Host+":"+config.C.Server.Port)
	wg.Wait()
}
