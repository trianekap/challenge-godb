package main

import (
	"database/sql"
	"fmt"
	"time"

	"challenge-godb/entity"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "laundry_enigma"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	transactionDetail := entity.TransactionDetail{Id: 5, Customer_Id: 5, Laundry_Services_Id: "S02", Item: 2}

	detailService(transactionDetail)
}

func detailService(transactionDetail entity.TransactionDetail) {
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// customer := entity.Customer{Id: 6, Name: "Zamal", Address: "jl.tol", Phone_number: "089871293010", Taken_item: 1}
	viewCustomer(tx)
	// insertCustomer(customer, tx)
	// updateCustomer(customer, tx)
	// deleteCustomer(5, tx)

	// service := entity.LaundryServices{Id: "S05", Service_name: "Shoe Laundry", Price: "30000", Unit: "1 pair of shoes"}
	//viewLaundryService(tx)
	// insertService(service, tx)
	// updateService(service, tx)
	// deleteService("S05", tx)

	// trade := entity.Transaction{Id: "LT05", Transaction_date: time.Date(2023, 8, 20, 0, 0, 0, 0, time.Local), Customer_id: 5}
	// viewTransaction(tx, trade)
	// insertTransaction(trade, tx)
	// updateTransaction(trade, tx)
	// deleteTransaction("LT05", tx)

	// viewTransactionDetail(tx)
	// insertTransactionDetail(transactionDetail, tx)
	// updateTransactionDetail(transactionDetail.Item, transactionDetail.Customer_Id, tx)
	// deleteTransactionDetail(5, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func viewCustomer(tx *sql.Tx) {
	query := "SELECT * FROM customer;"
	rows, err := tx.Query(query)
	if err != nil {
		validate(err, "View", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var address string
		var phone_number string
		var taken_item int

		err := rows.Scan(&id, &name, &address, &phone_number, &taken_item)
		if err != nil {
			validate(err, "View", tx)
			return
		}

		fmt.Printf(" %d, %s, %s, %s, %d,\n", id, name, address, phone_number, taken_item)
	}
}

func insertCustomer(customer entity.Customer, tx *sql.Tx) {
	insertCustomerDetail := "INSERT INTO customer (id, name, address, phone_number, taken_item) VALUES ($1, $2, $3, $4, $5);"

	_, err := tx.Exec(insertCustomerDetail, customer.Id, customer.Name, customer.Address, customer.Phone_number, customer.Taken_item)
	validate(err, "Insert", tx)
}

func updateCustomer(customer entity.Customer, tx *sql.Tx) {
	updateCustomer := "UPDATE customer SET name = $2, address = $3, phone_number = $4 WHERE id = $1;"

	_, err := tx.Exec(updateCustomer, customer.Id, customer.Name, customer.Address, customer.Phone_number)
	validate(err, "Update", tx)
}

func deleteCustomer(customerId int, tx *sql.Tx) {
	deleteCustomer := "DELETE FROM customer WHERE id = $1;"

	_, err := tx.Exec(deleteCustomer, customerId)
	validate(err, "Delete", tx)
}

func viewLaundryService(tx *sql.Tx) {
	query := "SELECT * FROM laundry_services;"
	rows, err := tx.Query(query)
	if err != nil {
		validate(err, "View", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var service_name string
		var price string
		var unit string

		err := rows.Scan(&id, &service_name, &price, &unit)
		if err != nil {
			validate(err, "View", tx)
			return
		}

		fmt.Printf(" %s, %s, %s, %s,\n", id, service_name, price, unit)
	}
}

func insertService(service entity.LaundryServices, tx *sql.Tx) {
	insertService := "INSERT INTO laundry_services (id, service_name, price, unit) VALUES ($1, $2, $3, $4);"

	_, err := tx.Exec(insertService, service.Id, service.Service_name, service.Price, service.Unit)
	validate(err, "Insert", tx)
}

func updateService(service entity.LaundryServices, tx *sql.Tx) {
	updateService := "UPDATE laundry_services SET service_name = $2, price = $3, unit = $4 WHERE id = $1;"

	_, err := tx.Exec(updateService, service.Id, service.Service_name, service.Price, service.Unit)
	validate(err, "Update", tx)
}

func deleteService(laundry_servicesId string, tx *sql.Tx) {
	deleteService := "DELETE FROM laundry_services WHERE id = $1;"

	_, err := tx.Exec(deleteService, laundry_servicesId)
	validate(err, "Delete", tx)
}

func viewTransaction(tx *sql.Tx, transaction entity.Transaction) {
	query := "SELECT * FROM laundry_transactions;"
	rows, err := tx.Query(query)
	if err != nil {
		validate(err, "View", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var transaction_date time.Time
		var customer_id int

		err := rows.Scan(&id, &transaction_date, &customer_id)
		if err != nil {
			validate(err, "View", tx)
			return
		}

		fmt.Printf(" %s, %s, %d,\n", id, transaction_date.Format("2006-01-02 00:00:00"), customer_id)
	}
}

func insertTransaction(transactions entity.Transaction, tx *sql.Tx) {
	insertTransaction := "INSERT INTO laundry_transactions (id, transaction_date, customer_id) VALUES ($1, $2, $3);"

	_, err := tx.Exec(insertTransaction, transactions.Id, transactions.Transaction_date, transactions.Customer_id)
	validate(err, "Insert", tx)
}

func updateTransaction(transactions entity.Transaction, tx *sql.Tx) {
	updateTransaction := "UPDATE laundry_transactions SET transaction_date = $2, customer_id = $3 WHERE id = $1;"

	_, err := tx.Exec(updateTransaction, transactions.Id, transactions.Transaction_date, transactions.Customer_id)
	validate(err, "Update", tx)
}

func deleteTransaction(laundry_transactionId string, tx *sql.Tx) {
	deleteTransaction := "DELETE FROM laundry_transactions WHERE id = $1;"

	_, err := tx.Exec(deleteTransaction, laundry_transactionId)
	validate(err, "Delete", tx)
}

func viewTransactionDetail(tx *sql.Tx) {
	query := "SELECT * FROM transaction_details;"
	rows, err := tx.Query(query)
	if err != nil {
		validate(err, "View", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var customerId int
		var laundryServicesId string
		var item int

		err := rows.Scan(&id, &customerId, &laundryServicesId, &item)
		if err != nil {
			validate(err, "View", tx)
			return
		}

		fmt.Printf(" %d, %d, %s, %d\n", id, customerId, laundryServicesId, item)
	}
}

func insertTransactionDetail(transactionDetail entity.TransactionDetail, tx *sql.Tx) {
	insertTransactionDetail := "INSERT INTO transaction_details (id, customer_id, laundry_services_id, item) VALUES ($1, $2, $3, $4);"

	_, err := tx.Exec(insertTransactionDetail, transactionDetail.Id, transactionDetail.Customer_Id, transactionDetail.Laundry_Services_Id, transactionDetail.Item)
	validate(err, "Insert", tx)
}

func updateTransactionDetail(item int, customerId int, tx *sql.Tx) {
	updateTransactionDetail := "UPDATE transaction_details SET item = $1 WHERE customer_id = $2;"

	_, err := tx.Exec(updateTransactionDetail, item, customerId)
	validate(err, "Update", tx)
}

func deleteTransactionDetail(id int, tx *sql.Tx) {
	deleteDetail := "DELETE FROM transaction_details WHERE id = $1;"

	_, err := tx.Exec(deleteDetail, id)
	if err != nil {
		validate(err, "Delete", tx)
	} else {
		fmt.Println("Successfully Deleted Data!")
	}
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback!")
	} else {
		fmt.Println("Successfully " + message + " data!")
	}
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected!")
	}

	return db
}
