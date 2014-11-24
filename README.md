#Mapbox GeoCode + Image Server

This is a simple mapbox geocoder and image server, it just combines two of mapbox services to provide an easier API.

Sample Server: https://mapboxgeoimage.herokuapp.com/Berlin/10/1280x120
![berlinimg](https://mapboxgeoimage.herokuapp.com/Berlin/10/1280x120)

##What it does
- It uses Mapbox' Geocoder API
- It uses Mapbox' Static Maps API
- It caches an image for later use

##How to build your own API Server

```golang
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

  n.Run(":" + os.Getenv("PORT"))
}
```
