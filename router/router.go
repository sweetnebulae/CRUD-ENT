package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sweetnebulae/go_ent/controller"
)

// NewRouter sets up the HTTP routes and maps them to the corresponding controller methods.
//
// Routes:
//   - GET    /api/posts        : Retrieve all posts (books).
//   - GET    /api/posts/:id    : Retrieve a single post by UUID.
//   - POST   /api/post         : Create a new post.
//   - PUT    /api/post/:id     : Update an existing post.
//   - DELETE /api/post/:id     : Delete a post by UUID.
//
// Parameters:
//   - postController: pointer to PostController that implements the handler logic.
//
// Returns:
//   - A configured *httprouter.Router instance ready to serve.
func NewRouter(postController *controller.PostController) *httprouter.Router {
	router := httprouter.New()

	// Define RESTful routes
	router.GET("/api/posts", postController.FindAll)
	router.GET("/api/posts/:id", postController.FindById)
	router.POST("/api/post", postController.Create)
	router.PUT("/api/post/:id", postController.Update)
	router.DELETE("/api/post/:id", postController.Delete)

	return router
}
