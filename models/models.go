// modles define data structure
package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type BookingCollection map[string]Booking

type Booking map[string][]string

type jsonMovie Movie

type Movie struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Rating   float32 `json:"rating"`
	Director string  `json:"director"`
}

type Catalog map[string]Movie

type Showtimes map[string][]string

type Timestamp time.Time

type UserCollection map[string]User

type User struct {
	Username   string     `json:"id"`
	Name       string     `json:"name"`
	LastActive *Timestamp `json:"last_active"`
}

type UserBookings map[string][]Movie
