package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// membuat fungsi untuk mendapatkan deskripsi webpage dari API dan mengembalikan data yang diperoleh
func getDescription() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/description")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

// membuat fungsi untuk mendapatkan daftar pekerjaan dari API dan mengembalikan data yang diperoleh
func getJobList() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/jobs")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

// membuat type untuk menyimpan cache dari data yang telah diambil dari API
type Cache struct {
	DescString string
	JobList    string
	IsFilled   bool
}

// membuat fungsi untuk mengambil data deskripsi dan daftar pekerjaan secara concurrent
// yang kemudian disimpan dalam type cache
func (c *Cache) aggregate() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	if !c.IsFilled {

		go func() {
			c.DescString = getDescription()
			wg.Done()
		}()

		go func() {
			c.JobList = getJobList()
			wg.Done()
		}()

		wg.Wait()

		c.IsFilled = true
	}

	fmt.Println(c.DescString)
	fmt.Println(c.JobList)
}

func main() {

	start1 := time.Now()
	fmt.Println("dari calculate ", start1)

	chc := Cache{}
	chc.aggregate()

	// menghitung waktu yang dibutuhkan untuk mengambil data dari API
	fmt.Printf("took %v\n", time.Since(start1))

	fmt.Printf("\n\n")

	start2 := time.Now()
	fmt.Println("dari calculate ", start2)

	chc.aggregate()

	// menghitung waktu yang dibutuhkan untuk mengambil data dari cache
	fmt.Printf("took %v\n", time.Since(start2))

	fmt.Printf("\n\n")
}
