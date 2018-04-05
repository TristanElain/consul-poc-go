package controller

import (
	"consul-poc/manager"
	"consul-poc/model"
	"consul-poc/provider"
	"consul-poc/utils"
	"html/template"
	"net/http"
	"strconv"
)

// BaseController -
type BaseController struct{}

const tplDir string = "templates/"

// HandleForm - handle form request
func (controller *BaseController) HandleForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.FatalError(err)
	health := r.Form.Get("health")

	healthManager := manager.GetHealthManager()
	if health != "" {
		var healthBool bool
		healthBool, err = strconv.ParseBool(health)
		utils.FatalError(err)

		healthManager.SetHealth(healthBool)
	}
	controller.executeTemplate(w, "healthForm.gohtml", healthManager.IsHealthy())
}

// HandleProps - handle props request
func (controller *BaseController) HandleProps(w http.ResponseWriter, r *http.Request) {
	applicationProps := model.NewEmptyApplicationProperties()
	propertiesProvider := provider.GetPropertiesProvider()

	err := propertiesProvider.GetProperties("poc/consul/global", nil, &applicationProps)
	utils.FatalError(err)

	controller.executeTemplate(w, "index_consul_poc.gohtml", applicationProps)
}

// HandleMapProps - handle map props request
func (controller *BaseController) HandleMapProps(w http.ResponseWriter, r *http.Request) {
	propertiesProvider := provider.GetPropertiesProvider()

	props, err := propertiesProvider.GetPropertiesMap("poc/consul/global", nil)
	utils.FatalError(err)

	controller.executeTemplate(w, "map_consul_poc.gohtml", props)
}

func (controller *BaseController) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	healthManager := manager.GetHealthManager()

	if healthManager.IsHealthy() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

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
