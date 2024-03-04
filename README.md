# challenge-godb

1. Pastikan anda memiliki database yang telah diatur dan siap digunakan oleh aplikasi. Atur informasi koneksi database pada bagian berikut dalam kode aplikasi (enigma_laundry.go) di Visual Studio Code:

example :

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "laundry_enigma"
)

Ganti nilai-nilai diatas sesuai dengan konfigurasi database Anda.

2. Penggunaan Dan menjalankan Aplikasi
Di menu utama terdapat beberapa opsi yang tersedia:
    1. Customer
    2. Laundry Services
    3. Laundry Transactions
    4. Transaction Details
Untuk mengelola data customer Anda dapat melakuka view, insert, update dan delete data. Cara menjalankannya aplikasinya:
    1. Untuk melakukan menu view customer Anda bisa melakukan Uncomment di menu viewCustomer(tx),
       // customer := entity.Customer{Id: 6, Name: "Zamal", Address: "jl.tol", Phone_number: "089871293010", Taken_item: 1}
	     viewCustomer(tx)
	     // insertCustomer(customer, tx)
	     // updateCustomer(customer, tx)
	     // deleteCustomer(5, tx)

       output :
        ~/Development/GOLANG/src/challenge-godb [12:30:02]
          $ go run main.go
          Successfully Connected!
           1, Rizal, Jln.Kh.hasan arif, 08123456789, 4,
           2, Fauzan, jl.Maleer, 082123128932, 1,
           3, Wildan, jl.Banyuresmi, 089876251627, 3,
           4, Sena, jl.Cirendeu, 087654216281, 5,
           5, Desy, Perum Arumaya, 085162718291, 2,

    2. 
  





