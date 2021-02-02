package types

//Cut holds the info for one individual cut
type Cut struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Dimensions string  `json:"dimensions"`
	Length     float64 `json:"length"`
}

//DimensionGroup holds like dimensions
type DimensionGroup struct {
	Cuts        []float64
	Dimension   string
	TotalLength float64
}

//DimensionGroups groups all like cuts together
type DimensionGroups map[string]DimensionGroup

//ShoppingList gets the shopping list for one dimension
type ShoppingList struct {
	Dimension    string      `json:"dimension"`
	LenBoard     float64     `json:"len_board"`
	Quantity     float64     `json:"quantity"`
	ScrapLengths []float64   `json:"scrap_lengths"`
	Boards       [][]float64 `json:"boards"`
}
