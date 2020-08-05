package main

import "time"

type Station struct {
	ID                int     `json:"id" db:"id"`
	Name              string  `json:"name" db:"name"`
	Distance          float64 `json:"-" db:"distance"`
	IsStopExpress     bool    `json:"is_stop_express" db:"is_stop_express"`
	IsStopSemiExpress bool    `json:"is_stop_semi_express" db:"is_stop_semi_express"`
	IsStopLocal       bool    `json:"is_stop_local" db:"is_stop_local"`
}

type DistanceFare struct {
	Distance float64 `json:"distance" db:"distance"`
	Fare     int     `json:"fare" db:"fare"`
}

type Fare struct {
	TrainClass     string    `json:"train_class" db:"train_class"`
	SeatClass      string    `json:"seat_class" db:"seat_class"`
	StartDate      time.Time `json:"start_date" db:"start_date"`
	FareMultiplier float64   `json:"fare_multiplier" db:"fare_multiplier"`
}

type Train struct {
	Date         time.Time `json:"date" db:"date"`
	DepartureAt  string    `json:"departure_at" db:"departure_at"`
	TrainClass   string    `json:"train_class" db:"train_class"`
	TrainName    string    `json:"train_name" db:"train_name"`
	StartStation string    `json:"start_station" db:"start_station"`
	LastStation  string    `json:"last_station" db:"last_station"`
	IsNobori     bool      `json:"is_nobori" db:"is_nobori"`
}

type Seat struct {
	TrainClass    string `json:"train_class" db:"train_class"`
	CarNumber     int    `json:"car_number" db:"car_number"`
	SeatColumn    string `json:"seat_column" db:"seat_column"`
	SeatRow       int    `json:"seat_row" db:"seat_row"`
	SeatClass     string `json:"seat_class" db:"seat_class"`
	IsSmokingSeat bool   `json:"is_smoking_seat" db:"is_smoking_seat"`
}

type Reservation struct {
	ReservationId int        `json:"reservation_id" db:"reservation_id"`
	UserId        *int       `json:"user_id" db:"user_id"`
	Date          *time.Time `json:"date" db:"date"`
	TrainClass    string     `json:"train_class" db:"train_class"`
	TrainName     string     `json:"train_name" db:"train_name"`
	Departure     string     `json:"departure" db:"departure"`
	Arrival       string     `json:"arrival" db:"arrival"`
	PaymentStatus string     `json:"payment_status" db:"payment_status"`
	Status        string     `json:"status" db:"status"`
	PaymentId     string     `json:"payment_id,omitempty" db:"payment_id"`
	Adult         int        `json:"adult" db:"adult"`
	Child         int        `json:"child" db:"child"`
	Amount        int        `json:"amount" db:"amount"`
}

type SeatReservation struct {
	ReservationId int    `json:"reservation_id,omitempty" db:"reservation_id"`
	CarNumber     int    `json:"car_number,omitempty" db:"car_number"`
	SeatRow       int    `json:"seat_row" db:"seat_row"`
	SeatColumn    string `json:"seat_column" db:"seat_column"`
}

// 未整理

type CarInformation struct {
	Date                string                 `json:"date"`
	TrainClass          string                 `json:"train_class"`
	TrainName           string                 `json:"train_name"`
	CarNumber           int                    `json:"car_number"`
	SeatInformationList []SeatInformation      `json:"seats"`
	Cars                []SimpleCarInformation `json:"cars"`
}

type SimpleCarInformation struct {
	CarNumber int    `json:"car_number"`
	SeatClass string `json:"seat_class"`
}

type SeatInformation struct {
	Row           int    `json:"row"`
	Column        string `json:"column"`
	Class         string `json:"class"`
	IsSmokingSeat bool   `json:"is_smoking_seat"`
	IsOccupied    bool   `json:"is_occupied"`
}

type SeatInformationByCarNumber struct {
	CarNumber           int               `json:"car_number"`
	SeatInformationList []SeatInformation `json:"seats"`
}

type TrainSearchResponse struct {
	Class            string            `json:"train_class"`
	Name             string            `json:"train_name"`
	Start            string            `json:"start"`
	Last             string            `json:"last"`
	Departure        string            `json:"departure"`
	Arrival          string            `json:"arrival"`
	DepartureTime    string            `json:"departure_time"`
	ArrivalTime      string            `json:"arrival_time"`
	SeatAvailability map[string]string `json:"seat_availability"`
	Fare             map[string]int    `json:"seat_fare"`
}

type User struct {
	ID             int64
	Email          string `json:"email"`
	Password       string `json:"password"`
	Salt           []byte `db:"salt"`
	HashedPassword []byte `db:"super_secure_password"`
}

type TrainReservationRequest struct {
	Date          string        `json:"date"`
	TrainName     string        `json:"train_name"`
	TrainClass    string        `json:"train_class"`
	CarNumber     int           `json:"car_number"`
	IsSmokingSeat bool          `json:"is_smoking_seat"`
	SeatClass     string        `json:"seat_class"`
	Departure     string        `json:"departure"`
	Arrival       string        `json:"arrival"`
	Child         int           `json:"child"`
	Adult         int           `json:"adult"`
	Column        string        `json:"Column"`
	Seats         []RequestSeat `json:"seats"`
}

type RequestSeat struct {
	Row    int    `json:"row"`
	Column string `json:"column"`
}

type TrainReservationResponse struct {
	ReservationId int64 `json:"reservation_id"`
	Amount        int   `json:"amount"`
	IsOk          bool  `json:"is_ok"`
}

type ReservationPaymentRequest struct {
	CardToken     string `json:"card_token"`
	ReservationId int    `json:"reservation_id"`
}

type ReservationPaymentResponse struct {
	IsOk bool `json:"is_ok"`
}

type PaymentInformationRequest struct {
	CardToken     string `json:"card_token"`
	ReservationId int    `json:"reservation_id"`
	Amount        int    `json:"amount"`
}

type PaymentInformation struct {
	PayInfo PaymentInformationRequest `json:"payment_information"`
}

type PaymentResponse struct {
	PaymentId string `json:"payment_id"`
	IsOk      bool   `json:"is_ok"`
}

type ReservationResponse struct {
	ReservationId int               `json:"reservation_id"`
	Date          string            `json:"date"`
	TrainClass    string            `json:"train_class"`
	TrainName     string            `json:"train_name"`
	CarNumber     int               `json:"car_number"`
	SeatClass     string            `json:"seat_class"`
	Amount        int               `json:"amount"`
	Adult         int               `json:"adult"`
	Child         int               `json:"child"`
	Departure     string            `json:"departure"`
	Arrival       string            `json:"arrival"`
	DepartureTime string            `json:"departure_time"`
	ArrivalTime   string            `json:"arrival_time"`
	Seats         []SeatReservation `json:"seats"`
}

type CancelPaymentInformationRequest struct {
	PaymentId string `json:"payment_id"`
}

type CancelPaymentInformationResponse struct {
	IsOk bool `json:"is_ok"`
}

type Settings struct {
	PaymentAPI string `json:"payment_api"`
}

type InitializeResponse struct {
	AvailableDays int    `json:"available_days"`
	Language      string `json:"language"`
}
