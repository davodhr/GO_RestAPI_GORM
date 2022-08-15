package models

import "gorm.io/gorm"

// Rectangle represent coordinates of a rectangle
type Rectangle struct {
	gorm.Model
	X_rect int64  `json:"x"`
	Y_rect int64  `json:"y"`
	W_rect int64  `json:"width"`
	H_rect int64  `json:"height"`
	Time   string `json:"time,omitempty"`
}

// Data structure in post method for recieving new data
type Data_in struct {
	Main_Rectangle Rectangle   `json:"main"`
	Inputs         []Rectangle `json:"inputs"`
}

// Data structure to be stored in DB
type Data_db struct {
	Rectangle Rectangle
	Time      string `json:"time"`
}

type PureRectangle struct {
	X_rect int64  `json:"x"`
	Y_rect int64  `json:"y"`
	W_rect int64  `json:"width"`
	H_rect int64  `json:"height"`
	Time   string `json:"time,omitempty"`
}
