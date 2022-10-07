package structs

type Rental struct {
	RentalId    uint   `gorm:"primaryKey" json:"rental_id"`
	RentalDate  string `json:"rental_date"`
	InventoryId uint   `json:"inventory_id"`
	CustomerId  uint   `json:"customer_id"`
	ReturnDate  string `json:"return_date"`
	StaffId     uint   `json:"staff_id"`
	LastUpdate  string `json:"last_update"`
}
