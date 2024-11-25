package helper

import (
	"fmt"
	"runtime"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("SEBELUM UNIT TEST")

	//* untuk menjalankan semua testing dalam 1 package
	m.Run()

	fmt.Println("SETELAH UNIT TEST")

	//* TestMain ini biasa digunakan ketika , unit test kita membutuhkan
	//* connect ke databse. dan setelah selesai menjalankan unit test , akan
	//* menutup database.
}

func TestTableHelloWorld(t *testing.T) {
	//* membuat slice struct
	tests := []struct {
		name     string
		request  string
		expectec string
	}{
		//* isi data slice struct
		{
			name:     "TestAzka",
			request:  "Azka",
			expectec: "Hello Azka",
		},
		{
			name:     "TestFitra",
			request:  "Fitra",
			expectec: "Hello Fitra",
		},
		{
			name:     "TestRamadhan",
			request:  "Ramadhan",
			expectec: "Hello Ramadhan",
		},
	}

	//* menjalankan table test
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := HelloWorld(test.request)
			require.Equal(t, test.expectec, res, "Result must be 'Hello Azka'")
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Azka", func(t *testing.T) {
		res := HelloWorld("Azka")
		require.Equal(t, "Hello Azka", res, "Result must be 'Hello Azka'")
	})
	t.Run("Fitra", func(t *testing.T) {
		res := HelloWorld("Fitra")
		require.Equal(t, "Hello Fitra", res, "Result must be 'Hello Fitra'")
	})
}

func TestSkip(t *testing.T) {
	res := HelloWorld("Azka")

	if runtime.GOOS == "darwin" {
		t.Skip("Can not run in MAC OS")
	}

	//* require.Equal akan memanggil juga t.FailNow() ,
	//* sehingga ketentuan pada t.FailNow akan dijalankan juga.
	require.Equal(t, "Hello Azka", res, "Result must be 'Hello Azka'")

	//* test success, print this.
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldRequire(t *testing.T) {
	res := HelloWorld("Azka")

	//* require.Equal akan memanggil juga t.FailNow() ,
	//* sehingga ketentuan pada t.FailNow akan dijalankan juga.
	require.Equal(t, "Hello Azka", res, "Result must be 'Hello Azka'")

	//* test success, print this.
	fmt.Println("TestHelloWorldRequire Done")
}
func TestHelloWorldAssert(t *testing.T) {
	res := HelloWorld("Azka")

	//* assert.Equal akan memanggil juga t.Fail() ,
	//* sehingga ketentuan pada t.Fail akan dijalankan juga.
	assert.Equal(t, "Hello Azka", res, "Result must be 'Hello Azka'")

	//* test success, print this.
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldAzka(t *testing.T) {
	res := HelloWorld("Azka")
	fmt.Printf("res: %v\n", res)
	if res != "Hello Azka" {
		//* error
		//* jangan menggunakan panic untuk menggagalkan unit test.
		// panic("Result is not 'Hello Azka'")

		//* solusinya gunakan seperti di bawah ini
		// t.Fail() // di commnent karena akan di optimasi dengan menggunakan error()
		//* Fail() akan menggagalkan unit test,
		//* namun tetap melanjutkan eksekusi unit test. Namun diakhir
		//* ketika selesai, maka unit test tersebut dianggap gagal.

		//* optimation t.Fail(), gunakan Error()
		//* Error() -> Lebih seperti melakukan log (print)
		//* error, namun setelah melakukan log error, dia
		//* akan secara otomatis memanggil function Fail(),
		//* sehingga mengakibatkan unit test dianggap gagal.
		t.Error()
	}

	//* test success, print this.
	fmt.Println("TestHelloWorldAzka Done")
}

func TestHelloWorlFitra(t *testing.T) {
	res := HelloWorld("Fitra")
	fmt.Printf("res: %v\n", res)
	if res != "Hello Fitra" {

		//* error
		//* jangan menggunakan panic untuk menggagalkan unit test.
		// panic("Result is not 'Hello Fitra'")

		//* solusinya gunakan seperti di bawah ini
		// t.FailNow() // di commnent karena akan di optimasi dengan menggunakan Fatal()
		//* FailNow() akan menggagalkan unit test saat ini
		//* juga, tanpa melanjutkan eksekusi unit test.

		//* optimation t.Fail(), gunakan Fatal()
		//* Fatal() -> mirip dengan Error(), hanya saja, setelah
		//* melakukan log error, dia akan memanggil
		//* FailNow(), sehingga mengakibatkan eksekusi unit test berhenti.

		t.Fatal()

	}

	//* test success, print this.
	fmt.Println("TestHelloWorlFitra Done")
}
