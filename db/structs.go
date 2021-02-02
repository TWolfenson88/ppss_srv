package db

type Station struct {
	Cnsi int    `json:"stan_id"`
	Road int    `json:"dor_kod"`
	Esr  int    `json:"st_kod"`
	Name string `json:"name"`
	Flag string `json:"flag"`
	Len  int    `json:"len"`
}
