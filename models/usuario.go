package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
//modelo de usuario mongodb
type Usuario struct {
	ID              primitive.ObjectID `bson: "_id,omiteempty" json:"id"`
	Nombre          string             `bson: "nombre" json:"nombre, omiteempty"`
	Apellidos       string             `bson: "apellidos" json:"apellidos, omiteempty"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento" json:"fechaNacimiento,omiteempty"`
	Email           string             `bson: "email" json:"email"`
	Password        string             `bson: "password" json:"password,omitempty"`
	Avatar          string             `bson: "avatar" json:"avatar,omitempty"`
	Banner          string             `bson: "banner" json:"banner,omitempty"`
	Ubicacion       string             `bson: "unicacion" json:"unicacion,omitempty"`
	SitioWeb        string             `bson: "sitioWeb" json:"sitioWeb,omitempty"`
}