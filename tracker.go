package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/easypostTester/app"
)

// TrackerController implements the tracker resource.
type TrackerController struct {
	*goa.Controller
}

// NewTrackerController creates a tracker controller.
func NewTrackerController(service *goa.Service) *TrackerController {
	return &TrackerController{Controller: service.NewController("TrackerController")}
}

// Create runs the create action.
func (c *TrackerController) Create(ctx *app.CreateTrackerContext) error {
	// TrackerController_Create: start_implement

	// Put your logic here

	// TrackerController_Create: end_implement
	res := &app.EasypostTracker{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *TrackerController) List(ctx *app.ListTrackerContext) error {
	// TrackerController_List: start_implement

	// Put your logic here

	// TrackerController_List: end_implement
	res := &app.EasypostTrackers{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TrackerController) Show(ctx *app.ShowTrackerContext) error {
	// TrackerController_Show: start_implement

	// Put your logic here

	// TrackerController_Show: end_implement
	res := &app.EasypostTracker{}
	return ctx.OK(res)
}
