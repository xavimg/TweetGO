package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario es el modelo de usario de la BD */
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido        string             `bson:"apellido" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"id"`
	Password        string             `bson:"password" json:"password, omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar, omitempty"`
	Banner          string             `bson:"banner" json:"banner, omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia, omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion, omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb, omitempty"`
}
