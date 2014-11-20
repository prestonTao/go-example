package models

import (
	"time"
)

type RoleLogin struct {
	TriggerTime  time.Time
	ManServerId  string
	GameServerId string
	AccountsId
	AccountsName
	TerraceId
	SceneId
	RoleId
	RoleName
	RoleLevel
	RoleGender
	RoleOccupation
	RoleGold
	AccountsCowry
	OnlineTimeCount
	TriggerType
	LoginIP
	LoginCountry
	Block
	RoleCharacteristic
	StorageTime
}
