package main

import (
	"github.com/elizabeth-phillips/woodMeasure/db"
	"github.com/elizabeth-phillips/woodMeasure/types"
)

func main() {
	cuts := db.ReadPlan("firepitChair.json")
	groupings := types.GroupDimensions(cuts)
	groupings.PrintFullList(96)
}
