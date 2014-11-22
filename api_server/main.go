package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/johannesboyne/mapboxGeoImage"
)

var mapId string
var accessToken string

func main() {
	mapId = os.Getenv("MAP_ID")
	accessToken = os.Getenv("ACCESS_TOKEN")
	MapGeoImg := mapboxGeoImage.NewBox(mapId, accessToken)
	log.Println(mapId + " - " + accessToken)
	router := mux.NewRouter()
	router.HandleFunc("/{location_name}/{zoom:[0-9]+}/{x:[0-9]+}x{y:[0-9]+}", MapGeoImg.HandleMapboxGeoToImageRequest)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
