package models

type Meal struct {
	Carbs        string  `json:"carbs"`
	Proteins     string  `json:"proteins"`
	Oils         string  `json:"oils"`
	Vegetables   string  `json:"vegetables"`
	Beverages    string  `json:"beverages"`
	Fruits       string  `json:"fruits"`
	Calories     float64 `json:"calories"`
	ProteinValue float64 `json:"proteinValue"`
	Cost         float64 `json:"cost"`
}
