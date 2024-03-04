<h2>Pengaturan Koneksi Database</h2>

<p>Pastikan Anda memiliki database yang telah diatur dan siap digunakan oleh aplikasi. Silakan atur informasi koneksi database pada bagian berikut dalam kode aplikasi (<code>enigma_laundry.go</code>):</p>

<pre><code>const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "laundry_enigma"
)
</code></pre>

<p>Ganti nilai-nilai di atas sesuai dengan konfigurasi database Anda.</p>

<h2>Penggunaan dan Menjalankan Aplikasi</h2>

<p>Aplikasi ini memiliki beberapa opsi yang tersedia di menu utama:</p>

<ol>
  <li><strong>Customer</strong></li>
  <li><strong>Laundry Services</strong></li>
  <li><strong>Laundry Transactions</strong></li>
  <li><strong>Transaction Details</strong></li>
</ol>

<p>Untuk mengelola data pelanggan, Anda dapat melihat, menambahkan, memperbarui, dan menghapus data. Berikut adalah cara menjalankan aplikasi:</p>

<ol>
  <li>Untuk melakukan menu "View Customer", Anda dapat melakukan Uncomment di bagian kode yang bersangkutan:</li>
</ol>

<pre><code>// customer := entity.Customer{Id: 6, Name: "Zamal", Address: "jl.tol", Phone_number: "089871293010", Taken_item: 1}
viewCustomer(tx)
// insertCustomer(customer, tx)
// updateCustomer(customer, tx)
// deleteCustomer(5, tx)
</code></pre>

<p>Output:</p>

<pre><code>~/Development/GOLANG/src/challenge-godb [12:30:02]
$ go run main.go
Successfully Connected!
 1, Rizal, Jln.Kh.hasan arif, 08123456789, 4,
 2, Fauzan, jl.Maleer, 082123128932, 1,
 3, Wildan, jl.Banyuresmi, 089876251627, 3,
 4, Sena, jl.Cirendeu, 087654216281, 5,
 5, Desy, Perum Arumaya, 085162718291, 2,
</code></pre>

<h3>2.2. Insert Customer</h3>

<p>Jika Anda ingin melakukan insert data customer, lakukan seperti kode berikut:</p>

<pre><code>customer := entity.Customer{Id: 6, Name: "Zamal", Address: "jl.tol", Phone_number: "089871293010", Taken_item: 1}
// viewCustomer(tx)
insertCustomer(customer, tx)
// updateCustomer(customer, tx)
// deleteCustomer(5, tx)
</code></pre>

<p>Lalu outputnya akan seperti:</p>

<pre><code>go run main.go
Successfully Connected!
Successfully Insert data!
</code></pre>

<p>Dan nanti akan ada di database.</p>

<h3>2.3. Update Customer</h3>

<p>Untuk melakukan menu update customer sama seperti sebelumnya Anda hanya perlu melakukan sedikit perubahan:</p>

<pre><code>customer := entity.Customer{Id: 6, Name: "jajang", Address: "jl.tolaten", Phone_number: "0123919192892", Taken_item: 2}
// viewCustomer(tx)
// insertCustomer(customer, tx)
updateCustomer(customer, tx)
// deleteCustomer(5, tx)
</code></pre>

<p>Outputnya akan:</p>

<pre><code>$ go run main.go
Successfully Connected!
Successfully Update data!
</code></pre>

<p>Kemudian cek di database maka data dengan id tersebut akan berubah.</p>

<h3>2.4. Delete Customer</h3>

<p>Di menu delete customer, Anda bisa lakukan comment di customer lalu tuliskan id yang ingin Anda hapus. Berikut contohnya dengan id 6:</p>

<pre><code>// customer := entity.Customer{Id: 6, Name: "jajang", Address: "jl.tolaten", Phone_number: "0123919192892", Taken_item: 2}
// viewCustomer(tx)
// insertCustomer(customer, tx)
// updateCustomer(customer, tx)
deleteCustomer(6, tx)
</code></pre>

<p>Maka outputnya:</p>

<pre><code>go run main.go
Successfully Connected!
Successfully Delete data!
</code></pre>

<p>Kemudian lihat kembali di database maka data dengan id tersebut sudah terhapus.</p>

<p>Setelah selesai menggunakan aplikasi, pastikan Anda telah menyimpan semua perubahan yang diperlukan sebelum keluar dari aplikasi.</p>
