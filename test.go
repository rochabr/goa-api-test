package main

import (
	"github.com/goadesign/goa"
	"goa-api-test/app"
)

// TestController implements the test resource.
type TestController struct {
	*goa.Controller
}

// NewTestController creates a test controller.
func NewTestController(service *goa.Service) *TestController {
	return &TestController{Controller: service.NewController("TestController")}
}

// Create runs the create action.
func (c *TestController) Create(ctx *app.CreateTestContext) error {
	// TestController_Create: start_implement

	// Put your logic here

	return nil
	// TestController_Create: end_implement
}

// List runs the list action.
func (c *TestController) List(ctx *app.ListTestContext) error {
	// TestController_List: start_implement

	// Put your logic here

	res := app.TestCollection{}
	return ctx.OK(res)
	// TestController_List: end_implement
}

// Show runs the show action.
func (c *TestController) Show(ctx *app.ShowTestContext) error {
	// TestController_Show: start_implement

	// Put your logic here

	res := &app.Test{}
	return ctx.OK(res)
	// TestController_Show: end_implement
}

// Update runs the update action.
func (c *TestController) Update(ctx *app.UpdateTestContext) error {
	// TestController_Update: start_implement

	// Put your logic here

	return nil
	// TestController_Update: end_implement
}
