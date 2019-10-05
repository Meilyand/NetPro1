# Tugas 1 Networking Programming - Meilyand Evriyan Timor 1301161769

## Soal 1
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%201/Soal-1.png">
</p>

### Jawaban
Diagram di atas merupakan proses dari TCP three-way handshake, di mana dimulai dari belum terdapat koneksi hingga melewati  beberapa states sampai koneksi dibuat. Dengan menggunakan tiga jenis pesan yang sebagai penentu transisi antar state, yaitu:

1. SYN: merupakan synchronize message, digunakan untuk memulai dan membuat koneksi
2. ACK: merupakan acknowledgment message, sebagai penanda bahwa menerima pesan SYN atau FIN
3. FIN: merupakan finish message, sebagai penanda bahwa server ingin memutuskan koneksi

Penjelasan setiap state sebagai berikut:

<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%201/Tabel-1.png">
</p>

<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%201/Tabel-2.png">
</p>

<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%201/Tabel-3.png">
</p>


## Soal 2
### Program For
```javascript
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```
#### Penjelasan
* program for melakukan loop sesuai dengan keadaan yang ditentukan dengan nilai awalan      i = 1

* for pertama mencetak nilai i, selama    i <= 3. Didapat hasil 1 2 3

* for ke dua mencetak nilai j selama j <= 9, dengan awalan j = 7. Didapat hasil 7 8 9

* for ke tiga merupakan loop forever dengan mencetak string “loop” , namun terdapat break sehingga loop hanya dilakukan sekali saja. Didapat hasil berupa string “loop”

* for ke empat mencetak nilai n dengan awalan n = 0, selama n <= 5 dengan  kondisi jika n mod 2 = 0 maka nilai n tidak akan dicetak, melainkan mencetak angka ganjil. Didapat hasil 1 3 5
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%202/for.png">
</p>

### Program If/Else
```javascript
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```
#### Penjelasan
* program if/else melakukan pengecekan dari suatu kondisi jika memuhi maka melakukan aksi dari kondisi tersebut

* kondisi pertama, jika 7 mod 2 = 0, maka mencetak string “7 is even”, jika bukan maka mencetak string “7 is odd”. Didapat hasil 7 is odd

* kondisi ke dua, jika 8 mod 2 = 0, maka mencetak string “8 is divisible by 4”, jika bukan maka lanjut ke kondisi ke tiga. Didapat hasil 8 is divisible by 4

* kondisi ke tiga memiliki awalan nilai num = 9 dengan kondisi, jika num < 0 maka cetak “9 is negative”, jika num < 10 maka cetak “9 has 1 digit”, jika tidak keduanya maka cetak “9 has multiple digits”. Didapat hasil 9 has 1 digit
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%202/ifElse.png">
</p>


## Soal 3
### Program Array
```javascript
package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```
#### Penjelasan
* program array terdapat variabel integer a[5] sebagai array dengan nilai yg masih kosong lalu dicetak. Didapat hasil emp: [0 0 0 0 0]

* Lalu, array pada index ke empat diisi dengan nilai 100 dan dicetak untuk semua array. Didapat hasil set: [0 0 0 0 100]. Kemudian dicetak kembali untuk index ke empat, didapat hasil get: 100

* Kemudian mencetak panjang array dengan sintak len(a), didapat hasil len: 5

* Selanjutnya, inisialisasi array kedua dengan semua index array b diisi dengan nilai integer dan dicetak. Didapat hasil: [1,2,3,4,5]

* Terakhir dilakukan nested loop untuk membentuk array 2 dimensi dengan nilai array didapat dari penjumlahan i+j didalam nested loop. Didapat hasil [[0 1 2] [1 2 3]]
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%203/array.png">
</p>

### Program Function
```javascript
package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
```
#### Penjelasan
* program function memanggil function plus dan plusPlus pada fungsi main.

* Function plus mengembalikan nilai a+b, sedangkan function plusPlus mengembalikan nilai a+b+c. Lalu, pada function main, memanggil function plus(1,2) yg disimpan pada variabel res dan dicetak, didapat hasil 1+2 = 3.

* Kemudian nilai variabel res diupdate dengan memanggil function plusPlus(1,2,3) dan dicetak, didapat hasil 1+2+3 = 6.
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%203/function.png">
</p>


## Soal 4
### Program Struct
```javascript
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
```
#### Penjelasan
* program struct membuat tipe data buatan, yaitu person dengan atribut string name dan integer age.

* Lalu dicetak dengan langsung memasukkan nilai ke parameternya yaitu dengan format: person({“Bob”,20}). Didapat hasil {Bob 20}

* Atau bisa juga dengan format: person({name: “Alice”, age: 30}). Didapat hasil {Alice 30}.

* Jika yang diberi nilai hanya salah satu atribut, maka yang lain didefault nol untuk integer dan string kosong untuk string. Didapat hasil {Fred 0}

* Jika type person menjadi nilai dari suatu variabel, contoh s;= person{name: “Sean”, age: 50}, maka untuk memanggil salah satu atribut harus dengan format s.name atau s.age. Didapat hasil Sean

* Jika ingin mengisi nilai dengan variabel s harus dengan format sp:= &s (memakai & sebagai pointer). Lalu dapat memanggil atribut person. Jika dilakukan sp.age = 51, maka nilai age akan diupdate menjadi 51 yang semulanya 50.
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%204/struct.png">
</p>

### Program Method
```javascript
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
```
#### Penjelasan
* program method dapat mengakases ke atribut struct dengan cara deklarasi variabel objek func dan nama fungsi, seperti: func (r *rect) area() int {
    return r.width * r.height
}
      dan
* func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

* Untuk memanggil fungsi dengan format, misal: r.area() dengan r:=rect{width: 10, height: 5}.

* Jika ingin mengisi r pada variabel lain misal sp, maka formatnya adalah rp := &r (menggunakan & sbg pointer), lalu memanggil fungsi dengan rp.area() 
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%204/method.png">
</p>


## Soal 5
### Program Multiple Return
```javascript
package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
```
#### Penjelasan
* program ini dapat mengembalikan nilai return sebanyak maksimal 2 kali sesuai parameter fungsi vals.


* Terlihat jika a,b := vals(), maka a akan bernilai 3 dan b bernilai 7. Sedangkan yang terakhir _, c := vals, maka c akan bernilai 7 karena 3 telah diisi oleh _.
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%205/multipleReturn.png">
</p>

### Program Command Line
```javascript
package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
```
#### Penjelasan
* program ini seperti melakukan pengecekan kondisi. Jika kita ekseskusi command “go run commandLine.go” maka word, numb, fork, svar, tail akan bernilai default seperti gambar di bawah yang pertama. Namun, hasilnya akan berbeda jika dengan command line yg berberda, yaitu seperti gambar di bawah yang kedua
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%205/commandLine.png">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%205/commandLine2.png">
</p>


## Soal 6
### Program Simple Web Application
```javascript
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}
```
#### Penjelasan
* program ini bertindak sebagai server untuk melayani request dari client dengan mengkases localhost:8000 pada browser. Port yang digunakan 8000 dan jika berhasil maka server akan menampilkan pesan “Hello, you’ve requested” 
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%206/simpleWebApp.png">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%206/simpleWebBrowser.png">
</p>


## Soal 7
### File config.json
```javascript
{
    "server" : {
        "port" : ":8000"
    }
}
```
#### Penjelasan
* Pertama jalankan perintah go get -u github.com/spf13/viper pada command line. Lalu bikin file config.json seperti di bawah dengan port 8000. Kemudian bikin code simpleWebApp.go dan jalankan, lanjut ke browser dengan localhost:8000. 
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%207/config1.png">
</p>

### File simpleWebApp.go
```javascript
package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(viper.GetString("server.port"), nil)
}
```
#### Penjelasan
* Fungsi viper.SetConfigType() digunakan untuk set jenis file konfigurasi.

* Fungsi .AddConfigPath() digunakan untuk mendaftarkan path folder dimana file-file konfigurasi berada. Fungsi ini bisa dipanggil beberapa kali, jika memang ada banyak file konfigurasi tersimpan dalam path berbeda.

* Statement .SetConfigName() dieksekusi dengan parameter berisi nama file konfigurasi secara eksplisit tanpa ekstensi. Misalkan nama file adalah app.config.json, maka parameter cukup ditulis app.config.
  
* Fungsi .ReadInConfig() digunakan untuk memproses file-file konfigurasi sesuai dengan path   dan nama yang sudah ditentukan.
#### Hasil
<p align="center">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%207/simpleWebApp1.png">
  <img src="https://github.com/Meilyand/NetPro1/blob/master/Nomor%207/simpleWebBrowser2.png">
</p>
