package model

type Cancelation struct {
	IsSuccess  bool `json:"IsSuccess"`
	NumOfFreed int  `json:"numOfFreedTable"`
	Remaining  int  `json:"remainingTable"`
}
