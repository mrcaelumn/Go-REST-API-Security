package controller

import (
	"github.com/goadesign/goa"
	"github.com/mrcaelumn/Go-REST-API-Security/api/app"
)

// ActionController implements the Action resource.
type ActionController struct {
	*goa.Controller
}

// NewActionController creates a Action controller.
func NewActionController(service *goa.Service) *ActionController {
	return &ActionController{Controller: service.NewController("ActionController")}
}

// GetToken runs the getToken action.
func (c *ActionController) GetToken(ctx *app.GetTokenActionContext) error {
	// ActionController_GetToken: start_implement

	// Put your logic here

	return nil
	// ActionController_GetToken: end_implement
}

// Request runs the request action.
func (c *ActionController) Request(ctx *app.RequestActionContext) error {
	// ActionController_Request: start_implement

	// Put your logic here

	return nil
	// ActionController_Request: end_implement
}
