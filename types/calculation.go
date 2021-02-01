package types

//GroupDimensions combines all cuts of the same dimension
func GroupDimensions(cuts []Cut) (dimensionGroups DimensionGroups) {
	dimensionGroups = make(DimensionGroups)
	for _, cut := range cuts {
		if dimension, ok := dimensionGroups[cut.Dimensions]; ok {
			dimension.Cuts = append(dimension.Cuts, cut)
			dimension.TotalLength += cut.Length * float64(cut.Quantity)
			dimensionGroups[cut.Dimensions] = dimension
		} else {
			dimensionGroups[cut.Dimensions] = DimensionGroup{
				[]Cut{cut},
				cut.Length * float64(cut.Quantity),
			}
		}
	}
	return
}

//NumBoard gets the number of boards needed for one dimension
func (dimensionGroup DimensionGroup) NumBoard(lenBoards float64) (shoppingList ShoppingList) {
	// boardQuantity = math.Ceil(dimensionGroup.TotalLength / lenBoards)
	// scrapBoard = (boardQuantity * lenBoards) - dimensionGroup.TotalLength
	shoppingList.Quantity = 1
	tempLength := lenBoards
	shoppingList.LenBoard = lenBoards
	shoppingList.Dimension = dimensionGroup.Cuts[0].Dimensions
	for _, cut := range dimensionGroup.Cuts {
		for i := 0; i < cut.Quantity; i++ {
			if (tempLength - cut.Length) > 0 {
				tempLength -= cut.Length
			} else {
				shoppingList.ScrapLengths = append(shoppingList.ScrapLengths, tempLength)
				shoppingList.Quantity++
				tempLength = lenBoards - cut.Length
			}
		}
	}
	shoppingList.ScrapLengths = append(shoppingList.ScrapLengths, tempLength)

	return
}
