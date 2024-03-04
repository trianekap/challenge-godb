package entity

import "time"

type Customer struct {
	Id           int
	Name         string
	Address      string
	Phone_number string
	Taken_item   int
}

type LaundryServices struct {
	Id           string
	Service_name string
	Price        string
	Unit         string
}

type Transaction struct {
	Id               string
	Transaction_date time.Time
	Customer_id      int
}

type TransactionDetail struct {
	Id                  int
	Customer_Id         int
	Laundry_Services_Id string
	Item                int
}
