package controllers

import (
	"net/http"

	m "../models"
	"github.com/gernest/utron/controller"
)

//Todo is a controller for Todo list
type CoreUsers struct {
	controller.BaseController
	Routes []string
}

//Home renders a todo list
func (t *CoreUsers) GetUsers() {
	t.RenderJSON(m.GetUsers(true), http.StatusOK)
}

//Home renders a todo list
func (t *CoreUsers) AddUser() {
	data := m.User{}
	req := t.Ctx.Request()
	req.ParseForm()
	if err := decoder.Decode(&data, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.RenderJSON(m.ErrorResult(err.Error(), "400"), http.StatusBadRequest)
		return
	} else {
		t.RenderJSON(m.AddUser(data, true), http.StatusOK)
	}
}

func (t *CoreUsers) ToggleActive() {
	t.RenderJSON(m.ToggleUserActive(t.Ctx.Params["id"]), http.StatusOK)
}

//Home renders a todo list
func (t *CoreUsers) EditPassword() {
	data := m.User{}
	req := t.Ctx.Request()
	req.ParseForm()
	if err := decoder.Decode(&data, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.RenderJSON(m.ErrorResult(err.Error(), "400"), http.StatusBadRequest)
		return
	} else {
		t.RenderJSON(m.ChangeUserPassword(t.Ctx.Params["id"], data.Password), http.StatusOK)
	}
}

//NewTodo returns a new  todo list controller
func NewCoreUsers() controller.Controller {
	return &CoreUsers{
		Routes: []string{
			"get;/api/coreusers;GetUsers",
			"post;/api/coreusers;AddUser",
			"post;/api/coreusers/active/{id};ToggleActive",
			"post;/api/coreusers/password/{id};EditPassword",
		},
	}
}
