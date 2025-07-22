// Package controller provides HTTP handlers for managing book posts.
package controller

import (
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sweetnebulae/go_ent/data/request"
	"github.com/sweetnebulae/go_ent/data/response"
	"github.com/sweetnebulae/go_ent/helper"
	"github.com/sweetnebulae/go_ent/service"
	"net/http"
)

// PostController handles HTTP requests related to posts.
type PostController struct {
	PostService service.PostService
}

// NewPostController returns a new instance of PostController.
//
// It takes a reference to the PostService implementation and binds it to the controller.
func NewPostController(postService *service.PostServiceImpl) *PostController {
	return &PostController{postService}
}

// Create handles HTTP POST requests to create a new post.
//
// It reads and parses the request body, delegates to the PostService for creation,
// and writes the appropriate HTTP response.
func (c *PostController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	PostCreateRequest := request.CreateRequest{}
	helper.ReadRequestBody(r, &PostCreateRequest)

	err := c.PostService.Create(r.Context(), PostCreateRequest)
	if err != nil {
		return
	}
	PostResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteRequestBody(w, PostResponse)
}

// FindAll handles HTTP GET requests to fetch all posts.
//
// It returns a list of post responses from the PostService.
func (c *PostController) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result, _ := c.PostService.FindAll(r.Context())
	GetResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteRequestBody(w, GetResponse)
}

// FindById handles HTTP GET requests to fetch a specific post by its UUID.
//
// It parses the `id` from the URL and fetches the post details from the PostService.
func (c *PostController) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ParamId := p.ByName("id")
	id, err := uuid.Parse(ParamId)
	helper.ErrorPanic(err)
	result, _ := c.PostService.FindById(r.Context(), id)
	GetResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteRequestBody(w, GetResponse)
}

// Update handles HTTP PUT requests to update a specific post.
//
// It reads the updated data from the request body, parses the ID, and sends
// the update request to the PostService.
func (c *PostController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Invalid ID format",
		}
		helper.WriteRequestBody(w, errorResponse)
		return
	}
	updateRequest := request.UpdateRequest{}
	helper.ReadRequestBody(r, &updateRequest)

	updateRequest.Id = id

	err = c.PostService.Update(r.Context(), updateRequest)
	if err != nil {
		errorResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		}
		helper.WriteRequestBody(w, errorResponse)
		return
	}

	PostResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Book updated successfully",
	}
	helper.WriteRequestBody(w, PostResponse)
}

// Delete handles HTTP DELETE requests to remove a post by its UUID.
//
// It parses the ID from the URL and sends a delete request to the PostService.
func (c *PostController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ParamId := p.ByName("id")
	id, err := uuid.Parse(ParamId)
	helper.ErrorPanic(err)
	err = c.PostService.Delete(r.Context(), id)
	if err != nil {
		return
	}
	PostResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteRequestBody(w, PostResponse)
}
