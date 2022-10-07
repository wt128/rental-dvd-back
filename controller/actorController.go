package controller

import (
	"fmt"
	"goapi/infrastructure"
	"goapi/structs"
	"goapi/utils"
	"net/http"
)

func actorAll(w http.ResponseWriter, r *http.Request) {
	db := infrastructure.Db
	actors := []structs.Actor{}
	db.Limit(10).Find(&actors)
	raw := utils.StructToJson[[]structs.Actor](actors)
	fmt.Fprint(w, string(raw))
}
