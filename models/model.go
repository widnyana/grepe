package models

type Customer struct {
}

type Film struct {
	ID              int     `json:"id,omitempty" bson:"_id"`
	Actors          []Actor `json:"actors,omitempty" bson:"Actors"`
	Category        string  `json:"category,omitempty" bson:"Category"`
	Description     string  `json:"description,omitempty" bson:"Description"`
	Length          string  `json:"length,omitempty" bson:"Length"`
	Rating          string  `json:"rating,omitempty" bson:"Rating"`
	RentalDuration  string  `json:"rental_duration,omitempty" bson:"Rental Duration"`
	ReplacementCost string  `json:"replacement_cost,omitempty" bson:"Replacement Cost"`
	SpecialFeatures string  `json:"special_features,omitempty" bson:"Special Features"`
	Title           string  `json:"title,omitempty" bson:"Title"`
}

type Store struct {
	ID        int         `json:"id,omitempty" bson:"_id"`
	Address   string      `json:"address,omitempty" bson:"Address"`
	Inventory []Inventory `json:"inventory,omitempty" bson:"Inventory"`
	//tbd
}

type Inventory struct {
	FilmID      int    `json:"film_id,omitempty" bson:"filmId"`
	Title       string `json:"title,omitempty" bson:"Film Title"`
	InventoryID int    `json:"inventory_id,omitempty" bson:"inventoryId"`
}

// Actor schema
type Actor struct {
	FirstName string `json:"first_name,omitempty" bson:"First name"`
	LastName  string `json:"last_day,omitempty" bson:"Last name"`
	ActorID   int    `json:"actor_id,omitempty" bson:"actorId"`
}
