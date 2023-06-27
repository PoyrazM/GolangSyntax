package slices

import "log"

func Cities() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	log.Println(cities)

	citiesCopy := make([]string, len(cities))
	log.Println(citiesCopy)
	copy(citiesCopy, cities)
	log.Println(citiesCopy)

	cities = append(cities, "Adana")
	log.Println(cities)
	log.Println(citiesCopy)

	log.Println(cities[1:4])
	log.Println(cities[1:])
	log.Println(cities[:2])
}
