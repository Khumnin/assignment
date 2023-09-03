package model

type Reservation struct {
	IsSuccess   bool   `json:"IsSuccess"`
	BookingId   string `json:"bookingId"`
	NumOfBooked int    `json:"numOfBookedTable"`
	Remaining   int    `json:"remainingTable"`
}
