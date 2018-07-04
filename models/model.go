package models

type AllNetWork struct {
	Name                   string  //名称
	Market_cap_share       string  //全球总市值占比
	Turnover_rate          string  //换手率
	Volume_24h             string  //24小时成交率
	Volume_24h_from        float64 //24小时成交量
	Market_cap_display_cny float64 //流通市值
	Timestamp              float64
}
type Bstk struct {
	Com_id                 string  //名称
	Market_name            string  //全球总市值占比
	Price_display_cny      float64 //法币价格
	Price_display          string  //价格
	Percent_change_display string  //今日涨幅
	Volume_24h             string  //流通市值
	Timestamp              float64
}
