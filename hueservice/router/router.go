package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gitlab.com/scotthome/hueservice/models"
	routing "gitlab.com/scotthome/hueservice/utils"
	"io"
	"net/http"
)

type HTTPRouter struct {
	router           *mux.Router
	lightsController LightsController
}

func NewHTTPRouter(lightsController LightsController) (*HTTPRouter, error) {
	httpRouter := &HTTPRouter{lightsController: lightsController}
	router := mux.NewRouter()
	httpRouter.router = router
	httpRouter.registerHandlers()
	return httpRouter, nil
}

func (httpRouter *HTTPRouter) Serve() error {
	return http.ListenAndServe(":80", cors.Default().Handler(routing.GenerateRouter(httpRouter.router)))
}

func (httpRouter *HTTPRouter) registerHandlers() {
	httpRouter.router.HandleFunc("/", httpRouter.health).Methods("GET")
	httpRouter.router.HandleFunc("/v1/rooms", httpRouter.getRooms).Methods("GET")

	httpRouter.router.HandleFunc("/v1/rooms/state", httpRouter.setGroupState).Methods("POST")
	httpRouter.router.HandleFunc("/v1/rooms/scene", httpRouter.setGroupScene).Methods("POST")
	httpRouter.router.HandleFunc("/v1/rooms/bri", httpRouter.setGroupBri).Methods("POST")
}

func (httpRouter *HTTPRouter) health(writer http.ResponseWriter, request *http.Request) {
	routing.SendPayload(map[string]interface{}{"alive": "true"}, writer)
}

func (httpRouter *HTTPRouter) getRooms(w http.ResponseWriter, r *http.Request) {
	lights, err := httpRouter.lightsController.Rooms()
	if err != nil {
		routing.RespondWithHTTPError("failed to get lights", err, w)
	}
	routing.SendPayload(lights, w)
}

func (httpRouter *HTTPRouter) setGroupScene(w http.ResponseWriter, r *http.Request) {
	var groupMod models.GroupModification
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&groupMod)
	if err != nil && err != io.EOF {
		routing.RespondWithHTTPError("Error while inserting player achievement", err, w)
		return
	}
	err = httpRouter.lightsController.SetRoomScene(groupMod.GroupName, groupMod.SceneName)
	if err != nil {
		routing.RespondWithHTTPError("failed to get lights", err, w)
		return
	}
	routing.SendPayload(map[string]interface{}{"success": "true"}, w)
}

func (httpRouter *HTTPRouter) setGroupState(w http.ResponseWriter, r *http.Request) {
	var groupMod models.GroupModification
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&groupMod)
	if err != nil && err != io.EOF {
		routing.RespondWithHTTPError("Error while inserting player achievement", err, w)
		return
	}
	err = httpRouter.lightsController.SetRoomState(groupMod.GroupName, groupMod.On)
	if err != nil {
		routing.RespondWithHTTPError("failed to get lights", err, w)
		return
	}
	routing.SendPayload(map[string]interface{}{"success": "true"}, w)
}

func (httpRouter *HTTPRouter) setGroupBri(w http.ResponseWriter, r *http.Request) {
	var groupMod models.GroupModification
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&groupMod)
	if err != nil && err != io.EOF {
		routing.RespondWithHTTPError("Error while setting lights brightness", err, w)
		return
	}
	err = httpRouter.lightsController.SetRoomBri(groupMod.GroupName, groupMod.Brightness)
	if err != nil {
		routing.RespondWithHTTPError("failed to get lights", err, w)
		return
	}
	routing.SendPayload(map[string]interface{}{"success": "true"}, w)
}
