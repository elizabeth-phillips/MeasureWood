package types

//GroupDimensions combines all cuts of the same dimension
func GroupDimensions(cuts []Cut) (dimensionGroups DimensionGroups) {
	dimensionGroups = make(DimensionGroups)
	for _, cut := range cuts {
		if dimension, ok := dimensionGroups[cut.Dimensions]; ok {
			for i := 0; i < cut.Quantity; i++ {
				dimension.Cuts = append(dimension.Cuts, cut.Length)
			}
			dimension.TotalLength += cut.Length * float64(cut.Quantity)
			dimensionGroups[cut.Dimensions] = dimension
		} else {
			dimension := DimensionGroup{
				[]float64{cut.Length},
				cut.Dimensions,
				cut.Length * float64(cut.Quantity),
			}
			for i := 0; i < cut.Quantity-1; i++ {
				dimension.Cuts = append(dimension.Cuts, cut.Length)
			}
			dimensionGroups[cut.Dimensions] = dimension
		}
	}
	return
}

//PossibleShoppingList gets the number of boards needed for one dimension
func (dimensionGroup DimensionGroup) PossibleShoppingList(lenBoards float64) (shoppingList ShoppingList) {
	shoppingList = ShoppingList{
		dimensionGroup.Dimension,
		lenBoards,
		1,
		make([]float64, 0),
		make([][]float64, 1),
	}

	shoppingList.Boards[int(shoppingList.Quantity)-1] = make([]float64, 0)
	tempLength := lenBoards
	for _, cut := range dimensionGroup.Cuts {
		if (tempLength - cut) > 0 {
			tempLength -= cut
		} else {
			shoppingList.ScrapLengths = append(shoppingList.ScrapLengths, tempLength)
			shoppingList.Quantity++
			tempLength = lenBoards - cut
			shoppingList.Boards = append(shoppingList.Boards, make([]float64, 0))
		}
		shoppingList.Boards[int(shoppingList.Quantity)-1] = append(shoppingList.Boards[int(shoppingList.Quantity)-1], cut)

	}
	shoppingList.ScrapLengths = append(shoppingList.ScrapLengths, tempLength)

	return
}
