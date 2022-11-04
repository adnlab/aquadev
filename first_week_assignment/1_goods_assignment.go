package main

import "fmt"

// membuat type untuk penyimpanan data barang-barang perikanan
type fisheries struct {
	name  string
	price int
}

// membuat fungsi untuk mengurutkan data barang
func goods_sorting(g []fisheries) []fisheries {
	for i := len(g); i > 0; i-- {
		for j := 1; j < i; j++ {
			if g[j].price < g[j-1].price {
				p_temp := g[j].price
				g[j].price = g[j-1].price
				g[j-1].price = p_temp

				n_temp := g[j].name
				g[j].name = g[j-1].name
				g[j-1].name = n_temp
			}
		}
	}
	return g
}

// membuat fungsi untuk menampilkan rekomendasi barang-barang
// yang paling banyak dapat dibeli dengan batas harga tertentu
func limit_buying(g []fisheries, limit int) {
	sorted_goods := goods_sorting(g)

	sum := 0
	// best_buy := make([]fisheries, 0)

	fmt.Printf("Total produk dengan harga dibawah Rp %d\n", limit)
	fmt.Printf("Harga total Rp %d\n", limit)
	for _, goods := range sorted_goods {
		// best_buy = append(best_buy, g[i])
		fmt.Printf("%s - %d\n", goods.name, goods.price)

		if sum += goods.price; sum >= limit {
			break
		}
	}
}

// membuat fungsi untuk mencari barang-barang dengan harga tertentu
func search_price(g []fisheries, num int) {
	// found := make([]fisheries, 0)

	fmt.Printf("Daftar produk dengan harga Rp %d\n", num)

	for _, goods := range g {
		if goods.price == num {
			// found = append(found, g[i])
			fmt.Printf("%s - %d\n", goods.name, goods.price)
		}
	}
}

// membuat fungsi untuk mencari barang-barang dengan harga paling mahal dan paling murah
func top_down_goods(g []fisheries) {
	sorted_goods := goods_sorting(g)

	fmt.Printf("Daftar produk termurah: %s Rp %d\n", sorted_goods[0].name, sorted_goods[0].price)
	fmt.Printf("Daftar produk termahal: %s Rp %d\n", sorted_goods[len(sorted_goods)-1].name, sorted_goods[len(sorted_goods)-1].price)
}

func main() {
	// mendefinisikan persediaan barang-barang yang dapat dibeli
	var allGoods = []fisheries{
		{name: "Benih Lele", price: 50000},
		{name: "Pakan Lele Cap Menara", price: 25000},
		{name: "Probiotik A", price: 75000},
		{name: "Probiotik Nila B", price: 10000},
		{name: "Pakan Nila", price: 20000},
		{name: "Benih Nila", price: 20000},
		{name: "Cupang", price: 5000},
		{name: "Benih Nila", price: 30000},
		{name: "Benih Cupang", price: 10000},
		{name: "Probiotik B", price: 10000},
	}

	fmt.Println("=================================================")

	// memanggil fungsi untuk mendapatkan rekomendasi barang-barang
	// yang paling banyak dapat diperoleh dengan batas Rp. 100.000
	limit_buying(allGoods, 100000)

	fmt.Println("=================================================")

	// memanggil fungsi pencarian barang dengan harga paling mahal
	// dan paling murah
	top_down_goods(allGoods)

	fmt.Println("=================================================")

	// memanggil fungsi pencari barang dengan harga Rp. 10.000
	search_price(allGoods, 10000)

}
