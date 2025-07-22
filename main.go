package main

import (
	"github.com/sweetnebulae/go_ent/config"
	"github.com/sweetnebulae/go_ent/controller"
	"github.com/sweetnebulae/go_ent/repository"
	"github.com/sweetnebulae/go_ent/router"
	"github.com/sweetnebulae/go_ent/service"
)

// main is the entry point of the application.
//
// It performs the following steps:
//  1. Connects to the database and ensures cleanup with `defer`.
//  2. Initializes the repository, service, and controller layers.
//  3. Sets up HTTP routes and binds handlers.
//  4. Starts the web server to serve requests.
//  5. Keeps the program alive using a blocking `select {}` to avoid exit.
func main() {
	// Initialize database connection
	client := config.ConnectDB()
	defer config.DisconnectDB(client)

	// Dependency injection: Repository → Service → Controller → Router
	postRepository := repository.NewPostRepositoryImpl(client)
	postService := service.NewPostServiceImpl(postRepository)
	postController := controller.NewPostController(postService)
	routes := router.NewRouter(postController)

	// Start HTTP server
	config.StartServer(routes)

	// Block main goroutine to keep the service running
	select {}
}
