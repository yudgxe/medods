package model

type Auth struct {
	tableName struct{} `pg:"auth,discard_unknown_columns"`

	Uuid         string `pg:"uuid,pk" json:"uuid"`
	RefreshToken string `pg:"refresh_token" json:"refresh_token"`
}
