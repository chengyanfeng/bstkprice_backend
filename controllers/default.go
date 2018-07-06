package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"bstkprice_backend/models"
)

var newWork = models.AllNetWork{}
var listRstk = make(map[int]models.Bstk)
var returnMap=make(map[string]interface{})
type P map[string]interface{}
type MainController struct {
	beego.Controller
}

func (c *MainController) GetData() {
	returnMap["listData"]=listRstk
	returnMap["summaryData"] = newWork
	c.Data["json"] = returnMap
	c.ServeJSON()
}

func Get(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	fmt.Printf("data", data)
	statusCode = resp.StatusCode
	content = string(data)
	return
}

func (c *MainController) GetToken() {

	s, statusCode := Get("https://mytoken.io/api/ticker/currencydetail?currency_on_market_id=821689817&timestamp=1530685209731&code=7b2c3a40edcc4847759b9cbed7bbc19a&platform=m&v=1.0.0&language=zh_CN&legal_currency=CNY")
	if statusCode != 200 {
		return
	}
	p := *JsonDecode([]byte(s))
	data := p["data"].(map[string]interface{})
	currency := data["currency"].(string)
	Market_cap_share := data["market_cap_share"].(string)
	Market_cap_display_cny := data["market_cap_display_cny"].(float64)
	Turnover_rate := data["turnover_rate"].(string)
	Volume_24h := data["volume_24h"].(string)
	Volume_24h_from := data["volume_24h_from"].(float64)
	Percent_change_display := data["percent_change_display"].(string)
	Price_display := data["price_display"].(string)
	Price_btc := data["price_btc"].(float64)
	Price_usd := data["price_usd"].(float64)
	Timestamps := p["timestamp"].(interface{})
	Timestamp := Timestamps.(float64)
	newWork.Name = currency
	newWork.Market_cap_share = Market_cap_share
	newWork.Market_cap_display_cny = Market_cap_display_cny
	newWork.Turnover_rate = Turnover_rate
	newWork.Volume_24h = Volume_24h
	newWork.Volume_24h_from = Volume_24h_from
	newWork.Percent_change_display = Percent_change_display
	newWork.Price_display = Price_display
	newWork.Price_btc = Price_btc
	newWork.Price_usd = Price_usd
	newWork.Timestamp = Timestamp
}
func (c *MainController) GetBstk() {

	s, statusCode := Get("https://mytoken.io/api/ticker/currencyexchangelist?currency_id=345463&page=1&size=10000&timestamp=1530692456025&code=359ac61abffb68f9532faa99fa649131&platform=m&v=1.0.0&language=zh_CN&legal_currency=CNY")
	if statusCode != 200 {
		return
	}
	p := *JsonDecode([]byte(s))
	data := p["data"].(map[string]interface{})
	Timestamp := p["timestamp"].(float64)
	list := data["list"].([]interface{})
	for k, v := range list {
		listv := v.(map[string]interface{})
		rstk := models.Bstk{}
		Com_id := listv["com_id"].(string)
		Volume_24h := listv["volume_24h"].(string)
		Market_name := listv["market_name"].(string)
		Percent_change_display := listv["percent_change_display"].(string)
		Price_display := listv["price_display"].(string)
		Price_display_cny := listv["price_display_cny"].(float64)
		rstk.Com_id = Com_id
		rstk.Volume_24h = Volume_24h
		rstk.Market_name = Market_name
		rstk.Percent_change_display = Percent_change_display
		rstk.Price_display = Price_display
		rstk.Price_display_cny = Price_display_cny
		rstk.Timestamp = Timestamp
		listRstk[k] = rstk
	}

}
func JsonDecode(b []byte) (p *P) {
	p = &P{}
	err := json.Unmarshal(b, p)
	fmt.Print(err)
	return
}
