#Mapbox GeoCode + Image Server

This is a simple mapbox geocoder and image server, it just combines two of mapbox services to provide an easier API.
If you're using the `api_server`, the URL-Schema looks like:

```
/{place}/{zoom: 0-22}/{X-width}x{Y-width} //X,Y width (0x0)->(1280x1280)
e.g.
http://.../Berlin/12/850x115
http://.../New-York-City/6/1280x115
```

The API-Server is just handy, use the go package for more sophisticated stuff!
