package models

/*Tweet es un tweet */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
