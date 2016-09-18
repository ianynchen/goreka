package eureka

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/*
Signal handler for ctrl-C and kill operations. Invoke this method in your
service implementation before starting server. This handler unregisters with
Eureka server first before exiting.
*/
func SigTermHandler() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		Unregister()
		os.Exit(1)
	}()
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	health := make(map[string]interface{})
	health["health"] = "OK"
	if err := json.NewEncoder(w).Encode(health); err != nil {
		panic(err)
	}
}

type EurekaHandler interface {
	Register(appId string) bool
	Unregister(appId, instanceId string) bool
	SendHeartbeat(appId, instanceId string) bool
	QueryForAllInstances() bool
	QueryForAllAppIdInstances(appId string) bool
	QueryForSpecific(appId, instanceId string) bool
	TakeOutOfService(appId, instanceId string) bool
	PutBackIntoService(appId, instanceId string) bool
	UpdateMetaData(appId, instanceId string) bool
	QueryForAllInstancesUnderVip(vip string) bool
	QueryForAllInstancesUnderSecureVip(svip string) bool
}
