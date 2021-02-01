package types

import (
	"fmt"
)

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

func (dimensionGroups DimensionGroups) PrintFullList(lenBoards float64) (shoppingList []ShoppingList) {
	for _, dimensionGroup := range dimensionGroups {
		shopList := dimensionGroup.PrintList(lenBoards)
		shoppingList = append(shoppingList, shopList)
	}
	return
}

func (dimensionGroup DimensionGroup) PrintList(lenBoards float64) (shoppingList ShoppingList) {
	fmt.Println("================")
	shoppingList = dimensionGroup.NumBoard(lenBoards)
	fmt.Printf("(%v) %v-inch %s board(s) needed\n", shoppingList.Quantity, shoppingList.LenBoard, shoppingList.Dimension)
	fmt.Println("Scrap lengths:", shoppingList.ScrapLengths)
	fmt.Printf("%s total Length: %v\n", shoppingList.Dimension, dimensionGroup.TotalLength)
	return
}

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
