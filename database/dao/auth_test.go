package dao

import (
	"medods/database/model"
	"testing"
)

func Test_CreateOrUpdate(t *testing.T) {
	const uuid string = "UUID"

	fatalOnError(t, Auth().CreateOrUpdate(&model.Auth{Uuid: uuid, RefreshToken: "TOKEN"}))
	fatalOnError(t, Auth().CreateOrUpdate(&model.Auth{Uuid: uuid, RefreshToken: "NEWTOKEN"}))
}
