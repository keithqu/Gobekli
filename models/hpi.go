package models

// HPI ... The house price index type
type HPI struct {
	Year     uint    `json:"year"`
	Month    uint    `json:"month"`
	GeoType  string  `json:"geo_type"`
	GeoName  string  `json:"geo_name"`
	GeoCode  string  `json:"geo_code"`
	IndexNsa float64 `json:"index_nsa"`
	IndexSa  float64 `json:"index_sa"`
}
