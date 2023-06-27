package arrays

import "log"

func Cities() {
	var cities [5]string
	cities[0] = "İzmir"
	cities[1] = "Antalya"
	cities[2] = "İstanbul"
	cities[3] = "Bursa"
	cities[4] = "Muğla"
	log.Println(cities)

	for i := 0; i < len(cities); i++ {
		log.Println(cities[i])
	}
}
