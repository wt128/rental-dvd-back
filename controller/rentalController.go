package controller

import (
	"fmt"
	"goapi/infrastructure"
	"goapi/structs"
	"goapi/utils"
	"net/http"
)

/*
	俳優の一覧を10件取得する
	   pages uint
*/
func RentalAll(w http.ResponseWriter, r *http.Request) {
	db := infrastructure.Db
	rentals := []structs.Rental{}
	db.Limit(10).Find(&rentals)
	raw := utils.StructToJson[[]structs.Rental](rentals)
	fmt.Fprint(w, string(raw))
}
