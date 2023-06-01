package model

type HistoryData struct {
	Num    int    `json:"num"`
	Length int    `json:"length"`
	Mode   string `json:"mode"`
	//AtTime time.Time `json:"at_time"`
	//CostTime time.Duration `json:"cost_time"`
	//Food     int           `json:"food"`
}