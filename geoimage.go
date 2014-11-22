package mapboxGeoImage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MapboxGeoFeatures struct {
	Id         string
	Type       string
	Text       string
	Place_name string
	Center     []float64
}

type MapboxGeoResponse struct {
	Features []MapboxGeoFeatures
}

type MapboxConfig struct {
	MapId       string
	AccessToken string
}

type BoxInterface interface {
	HandleMapboxGeoToImageRequest()
}

type Box struct {
	inMemoryKV map[string][]byte
	config     MapboxConfig
}

func NewBox(id string, accessToken string) Box {
	return Box{
		inMemoryKV: make(map[string][]byte),
		config: MapboxConfig{
			MapId:       id,
			AccessToken: accessToken,
		}}
}

func (b *Box) generateGeocoderURL(vars map[string]string) string {
	return "http://api.tiles.mapbox.com/v4/geocode/mapbox.places-v1/" + vars["location_name"] + ".json?access_token=" + b.config.AccessToken
}

func (b *Box) generateImageURL(lat string, lng string, vars map[string]string) string {
	return "http://api.tiles.mapbox.com/v4/" + b.config.MapId + "/pin-m-marker+4B98D3(" + lat + "," + lng + ")/" + lat + "," + lng + "," + vars["zoom"] + "/" + vars["x"] + "x" + vars["y"] + ".png?access_token=" + b.config.AccessToken
}

func (b *Box) HandleMapboxGeoToImageRequest(w http.ResponseWriter, req *http.Request) {
	if b.inMemoryKV[req.RequestURI] != nil {
		w.Write(b.inMemoryKV[req.RequestURI])
		return
	}
	vars := mux.Vars(req)
	res, err := http.Get(b.generateGeocoderURL(vars))
	if err != nil {
		panic(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var featureResponse MapboxGeoResponse
	err_json := json.Unmarshal(robots, &featureResponse)
	if err_json != nil {
		fmt.Println("error:", err_json)
	}

	lat := strconv.FormatFloat(featureResponse.Features[0].Center[0], 'f', 6, 64)
	lng := strconv.FormatFloat(featureResponse.Features[0].Center[1], 'f', 6, 64)
	url := b.generateImageURL(lat, lng, vars)
	log.Println(url)
	pic, perr := http.Get(url)
	if perr != nil {
		panic(perr)
	}
	picture, _ := ioutil.ReadAll(pic.Body)
	pic.Body.Close()
	b.inMemoryKV[req.RequestURI] = picture
	w.Write(picture)
}
