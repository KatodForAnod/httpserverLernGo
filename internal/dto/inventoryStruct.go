package dto

import "fmt"

// InventoryStruct - dto inventory
type InventoryStruct struct {
	NameOfInventory        string
	DescriptionOfInventory string
	ID                     int
	Assigned               UserStruct
}

// InventoryStructToString - get all information from dto inventory in string
func (s InventoryStruct) InventoryStructToString() string {
	return fmt.Sprintf("%s: %s - ID: %d @ %s\n",
		s.NameOfInventory,
		s.DescriptionOfInventory,
		s.ID, s.Assigned.UserName)
}
