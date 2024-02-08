package handlers

import "oapi-codegen-cultibio/controller"

func (h *MyHandler) CheckHandler() controller.ICheckController {
	checkController := controller.NewCheckController()
	return checkController
}
