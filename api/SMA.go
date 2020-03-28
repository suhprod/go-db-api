package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type data struct {
	Datetime        string `json:"datetime"`
	Subject         string `json:"subject"`
	Sscore          string `json:"sscore"`
	Sdispersion     string `json:"sdispersion"`
	Rawsscore200    string `json:"rawsscore200"`
	MktCapDominance string `json:"mkt_cap_dominance"`
	TradeDominance  string `json:"trade_dominance"`
	TweetDominance  string `json:"tweet_dominance"`
}
type tokendetails struct {
	APIToken    string `json:"api_token"`
	Function    string `json:"function"`
	ReqQuotaRem int32  `json:"request_quota_remaining"`
	Expires     string `json:"expires"`
	IP          string `json:"ip_address"`
	RecQuotaRem int32  `json:"records_quota_remaining"`
}
type searchParams struct {
	Format    string `json:"format"`
	Subject   string `json:"subject"`
	Ontology  string `json:"ontology"`
	Items     string `json:"items"`
	Dates     string `json:"dates"`
	Frequency string `json:"frequency"`
	Cycle     string `json:"cycle"`
	Filter    string `json:"filter"`
	Sort      string `json:"sort"`
	Domain    string `json:"domain"`
	Limit     string `json:"limit"`
	Tag       string `json:"tag"`
	Timezone  string `json:"timezone"`
	DataPref  string `json:"data_pref"`
	AppServer string `json:"appServer"`
	Received  string `json:"received"`
	Completed string `json:"completed"`
}

type response struct {
	Tokendetails tokendetails `json:"tokendetails"`
	SearchParams searchParams `json:"search_params"`
	Data         []data       `json:"data"`
}

//SMAResponse struct
type SMAResponse struct {
	Response response `json:"response"`
}

func prepareSMAUrl(path string, params map[string]string) string {
	baseURL := "" + path + "?"
	for key, value := range params {
		baseURL += key + "=" + value + "&"
	}
	return baseURL
}

//Sentiment Function
func Sentiment() SMAResponse {

	var params = make(map[string]string)
	params["items"] = os.Getenv("STATIC_DATAPOINTS")
	params["ontology"] = "crypto"
	params["frequency"] = "1d"
	params["filter"] = "tie_flag+eq+1"
	params["format"] = "json"
	params["timezone"] = "UTC"
	params["domain"] = "finance"
	params["limit"] = "10000"
	params["tag"] = "open_api_request"
	params["subject"] = "all"
	params["datetime"] = "datetime+eq+recent"

	requestURL := prepareSMAUrl("", params)
	headerData := url.Values{
		"function": {""},
		"api_key":  {""},
	}

	resp, err := http.PostForm(requestURL, headerData)

	if err != nil {
		panic(nil)
	}

	var obj SMAResponse
	bytes, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &obj)
	fmt.Println("jsonObj", obj)

	return obj
}
