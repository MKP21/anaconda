package anaconda

import (
	"net/url"
	"strconv"
)

type Location struct {
	Name  string `json:"name"`
	Woeid int    `json:"woeid"`
}

type Trend struct {
	Name            string `json:"name"`
	Query           string `json:"query"`
	Url             string `json:"url"`
	PromotedContent string `json:"promoted_content"`
	Volume          string `json:"tweet_volume"`
}

type TrendResponse struct {
	Trends    []Trend    `json:"trends"`
	AsOf      string     `json:"as_of"`
	CreatedAt string     `json:"created_at"`
	Locations []Location `json:"locations"`
}

type TrendLocation struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	ParentId    int    `json:"parentid"`
	PlaceType   struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"placeType"`
	Url   string `json:"url"`
	Woeid int32  `json:"woeid"`
}

// https://developer.twitter.com/en/docs/trends/trends-for-location/api-reference/get-trends-place
func (a TwitterApi) GetTrendsByPlace(id int64, v url.Values) (trendResp TrendResponse, err error) {
	response_ch := make(chan response)
	v = cleanValues(v)
	v.Set("id", strconv.FormatInt(id, 10))
	a.queryQueue <- query{a.baseUrl + "/trends/place.json", v, &[]interface{}{&trendResp}, _GET, response_ch}
	return trendResp, (<-response_ch).err
}

// https://developer.twitter.com/en/docs/trends/locations-with-trending-topics/api-reference/get-trends-available
func (a TwitterApi) GetTrendsAvailableLocations(v url.Values) (locations []TrendLocation, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/trends/available.json", v, &locations, _GET, response_ch}
	return locations, (<-response_ch).err
}

// https://developer.twitter.com/en/docs/trends/locations-with-trending-topics/api-reference/get-trends-closest
func (a TwitterApi) GetTrendsClosestLocations(lat float64, long float64, v url.Values) (locations []TrendLocation, err error) {
	response_ch := make(chan response)
	v = cleanValues(v)
	v.Set("lat", strconv.FormatFloat(lat, 'f', 6, 64))
	v.Set("long", strconv.FormatFloat(long, 'f', 6, 64))
	a.queryQueue <- query{a.baseUrl + "/trends/closest.json", v, &locations, _GET, response_ch}
	return locations, (<-response_ch).err
}
