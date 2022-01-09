package models

/*Relacion modelo para grabar la relacion de un usuario con otro */
type Relation struct {
	UserID         string `bson:"userID" json:"userID"`
	UserRelationID string `bson:"userRelationID" json:"userRelationID"`
}

type RelationRetrieved struct {
	Status bool `json:"status"`
}
