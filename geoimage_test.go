package mapboxGeoImage

import "testing"

var box = NewBox("map", "123")
var vars = make(map[string]string)

func TestGeoCoderURLGeneration(t *testing.T) {
	vars["location_name"] = "Berlin"
	if box.generateGeocoderURL(vars) != "http://api.tiles.mapbox.com/v4/geocode/mapbox.places-v1/Berlin.json?access_token=123" {
		t.Error("fault URL")
	}
}
func TestImageURLGeneration(t *testing.T) {
	vars["location_name"] = "Berlin"
	vars["zoom"] = "9"
	vars["x"] = "200"
	vars["y"] = "300"
	if box.generateImageURL("0", "0", vars) != "http://api.tiles.mapbox.com/v4/map/pin-m-marker+4B98D3(0,0)/0,0,9/200x300.png?access_token=123" {
		t.Error("fault URL")
	}
}
