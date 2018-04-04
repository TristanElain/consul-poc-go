package controller

import (
	"consul-poc/model"
	"consul-poc/provider"
	"consul-poc/utils"
	"html/template"
	"net/http"
)

// BaseController -
type BaseController struct{}

const tplDir string = "templates/"

// HandleRoot - handle root request
func (controller *BaseController) HandleRoot(w http.ResponseWriter, r *http.Request) {
	applicationProps := model.NewEmptyApplicationProperties()
	propertiesProvider, err := provider.NewPropertiesProvider()
	utils.FatalError(err)

	err = propertiesProvider.GetProperties("poc/consul/global", nil, &applicationProps)
	utils.FatalError(err)

	controller.executeTemplate(w, "index_consul_poc.gohtml", applicationProps)
}

// HandleMapProps - handle root request
func (controller *BaseController) HandleMapProps(w http.ResponseWriter, r *http.Request) {
	propertiesProvider, err := provider.NewPropertiesProvider()
	utils.FatalError(err)

	var props map[string]interface{}
	props, err = propertiesProvider.GetPropertiesMap("poc/consul/global", nil)
	utils.FatalError(err)

	controller.executeTemplate(w, "map_consul_poc.gohtml", props)
}

func (controller *BaseController) executeTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	templ, err := template.ParseFiles(tplDir + templateName)
	utils.FatalError(err)

	err = templ.Execute(w, data)
	utils.FatalError(err)
}

// NewBaseController - create a new BaseController structure
func NewBaseController() BaseController {
	baseController := BaseController{}
	return baseController
}
