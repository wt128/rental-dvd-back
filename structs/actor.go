package structs

import "time"

type Actor struct {
	ActorId    uint      `gorm:PrimaryKey json:"actor_id`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	LastUpdate time.Time `json:last_update`
}
