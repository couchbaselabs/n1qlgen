package main

import (
	"fmt"
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
	return fmt.Sprintf(qstr, pickOne(g.airports, g.seed))
}

// HotelRatingQueryGenerator generates queries on ratings of hotels
// from random cities with limit between 0 and 20
type HotelRatingGenerator struct {
	cities []string
	seed   int
}

func NewHotelRatingGenerator() *HotelRatingGenerator {
	cites := []string{"Medway", "Gillingham", "Giverny", "Giverny", "Glasgow", "Highland",
		"Glossop", "Padfield", "Glossop", "Glossop", "Santa Barbara",
		"Swansea", "Llanrhidian", "Swansea", "Preaux", "Greenhead",
		"Northumberland", "Northumberland", "Half Moon Bay"}
	return &HotelRatingGenerator{cities: cites, seed: 0}
}

func (g *HotelRatingGenerator) query() string {
	qstr := "SELECT name, (SELECT raw avg(s.ratings.Overall) " +
		"FROM   t.reviews  as s)[0] AS overall_avg_rating " +
		"FROM   `travel-sample` AS t " +
		"WHERE type = 'hotel' and city ='%s' " +
		"ORDER BY overall_avg_rating DESC " +
		"LIMIT %d; "
	g.seed += 1
	return fmt.Sprintf(qstr, pickOne(g.cities, g.seed), limitGen(g.seed))
}

// HotelReviewQueryGenerator generates queries on reviews
// of random hotels with limit between 0 and 20
type HotelReviewGenerator struct {
	hotels []string
	seed   int
}

func NewHotelReviewGenerator() *HotelReviewGenerator {

	hotels := []string{"Medway Youth Hostel", "The Balmoral Guesthouse", "The Robins",
		"Le Clos Fleuri", "Glasgow Grand Central", "Glencoe Youth Hostel",
		"The George Hotel", "Windy Harbour Farm Hotel", "Avondale Guest House",
		"The Bulls Head", "Bacara Resort & Spa", "Rhossili Bunkhouse",
		"Hill House Holiday Cottage", "Number 38 The Gower", "La Pradella",
		"The Greenhead Hotel and Hostel", "Once Brewed YHA Hostel"}
	return &HotelReviewGenerator{hotels: hotels, seed: 0}
}

func (g *HotelReviewGenerator) query() string {
	qstr := "SELECT name, cnt_reviewers " +
		"FROM   `travel-sample` AS t " +
		"LET cnt_reviewers = (SELECT raw count(*) " +
		"FROM t.reviews AS s " +
		"WHERE s.ratings.Overall >= 4)[0] " +
		"WHERE type = 'hotel' and name = '%s' " +
		"ORDER BY cnt_reviewers DESC " +
		"LIMIT %d;"
	g.seed += 1
	return fmt.Sprintf(qstr, pickOne(g.hotels, g.seed), limitGen(g.seed))
}
