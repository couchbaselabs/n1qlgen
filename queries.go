package main

import (
	"fmt"
	"math/rand"
)

// RouteQueryGenerator generates queries on route docs with pseudo random source airport
type RouteQueryGenerator struct {
	airports []string
	seed     int
}

func NewRouteQueryGenerator() *RouteQueryGenerator {
	airports := []string{"SFO", "ATL", "MNL", "TUN", "CDG", "TPE", "TLV", "PHX", "MRS"}
	return &RouteQueryGenerator{airports: airports, seed: 0}
}

func (g *RouteQueryGenerator) query() string {
	qstr := "SELECT  id, sourceairport, destinationairport, " +
		"(SELECT s.day, s.flight, str_to_tz(s.utc, 'Europe/London') as time " +
		"FROM `travel-sample`.schedule s " +
		"WHERE  str_to_tz(s.utc, 'Europe/London') > '18:00:00' " +
		"ORDER BY s.utc)  after_10pm " +
		"FROM `travel-sample` " +
		"WHERE type = 'route' and sourceairport = '%s'"
	g.seed += 1
	return fmt.Sprintf(qstr, g.airports[rand.Intn(g.seed)%len(g.airports)])
}

// HotelQueryGenerator generates queries on hotel docs
// from random cities with limit between 0 and 20
type HotelQueryGenerator struct {
	cities []string
	seed   int
}

func NewHotelQueryGenerator() *HotelQueryGenerator {
	cites := []string{"Medway", "Gillingham", "Giverny", "Giverny", "Glasgow", "Highland",
		"Glossop", "Padfield", "Glossop", "Glossop", "Santa Barbara",
		"Swansea", "Llanrhidian", "Swansea", "Preaux", "Greenhead",
		"Northumberland", "Northumberland", "Half Moon Bay"}
	return &HotelQueryGenerator{cities: cites, seed: 0}
}

func (g *HotelQueryGenerator) query() string {
	qstr := "SELECT name, (SELECT raw avg(s.ratings.Overall) " +
		"FROM   t.reviews  as s)[0] AS overall_avg_rating " +
		"FROM   `travel-sample` AS t " +
		"WHERE type = 'hotel' and city ='%s' " +
		"ORDER BY overall_avg_rating DESC " +
		"LIMIT %d; "
	g.seed += 1
	return fmt.Sprintf(qstr, g.cities[rand.Intn(g.seed)%len(g.cities)], rand.Intn(g.seed)%20)
}
