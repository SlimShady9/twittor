package models

/* GraboTweet es el modelo para insertar un tweet */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
