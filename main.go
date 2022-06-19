package main

import (
	"context"
	"exl-server/api/handler"
	"exl-server/repository"
	"exl-server/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Initializing")

	router := mux.NewRouter()
	ctx := context.Background()

	fileRepository := &repository.FileRepositoryImpl{}
	tenantRepository := &repository.TenantRepositoryImpl{}
	tenantService := service.NewTenantServiceImpl(tenantRepository)
	fileService := service.NewFileServiceImpl(fileRepository, tenantService)

	handler.MakeFileHandler(ctx, router, fileService)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8090",
		Handler:      router,
	}

	log.Println("Starting server")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	log.Println("Server Exited")
	
	
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutDownCtx); err != nil {
		log.Printf("Server Shutdown Failed:%+v\n", err)
	}
	log.Printf("Server Shutdown Properly\n")
	
}