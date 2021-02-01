package types

type Cut struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Dimensions string  `json:"dimensions"`
	Length     float64 `json:"length"`
}

type DimensionGroup struct {
	Cuts        []Cut
	TotalLength float64
}

type DimensionGroups map[string]DimensionGroup

type ShoppingList struct {
	Dimension    string    `json:"dimension"`
	LenBoard     float64   `json:"len_board"`
	Quantity     float64   `json:"quantity"`
	ScrapLengths []float64 `json:"scrap_lengths"`
}
