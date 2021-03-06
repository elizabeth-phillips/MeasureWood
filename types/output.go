package types

import (
	"fmt"
)

//PrintFullList prints the total number of boards needed
func (dimensionGroups DimensionGroups) PrintFullList(lenBoards float64) (shoppingList []ShoppingList) {
	for _, dimensionGroup := range dimensionGroups {
		shopList := dimensionGroup.PrintList(lenBoards)
		shoppingList = append(shoppingList, shopList)
	}
	return
}

//PrintList prints the number of boards of one dimension needed
func (dimensionGroup DimensionGroup) PrintList(lenBoards float64) (shoppingList ShoppingList) {
	fmt.Println("================")
	shoppingList = dimensionGroup.PossibleShoppingList(lenBoards)
	fmt.Printf("(%v) %v-inch %s board(s) needed\n", shoppingList.Quantity, shoppingList.LenBoard, shoppingList.Dimension)
	for i, cuts := range shoppingList.Boards {
		fmt.Printf("Board %v cuts: %s\n", i+1, printCuts(cuts))
	}
	fmt.Println("Scrap lengths:", shoppingList.ScrapLengths)
	fmt.Printf("%s total Length: %v\n", shoppingList.Dimension, dimensionGroup.TotalLength)
	return
}

func printCuts(cuts []float64) (output string) {
	for i, cut := range cuts {
		output += fmt.Sprint(cut)
		if i != len(cuts)-1 {
			output += ", "
		}
	}
	return
}
