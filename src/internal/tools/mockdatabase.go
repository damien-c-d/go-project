package tools

import (
	"time"
)

type mockDatabase struct{}

var mockLoginDetails = map[string]LoginDetails{
	"damien": {
		AuthToken: "ABC123",
		Username:  "bob",
	},
	"bella": {
		AuthToken: "DEF456",
		Username:  "jane",
	},
	"addison": {
		AuthToken: "GHI789",
		Username:  "john",
	},
}

var mockPointDetails = map[string]PointDetails{
	"damien": {
		Username: "bob",
		Balance:  100,
	},
	"bella": {
		Username: "jane",
		Balance:  200,
	},
	"addison": {
		Username: "john",
		Balance:  300,
	},
}

func (d *mockDatabase) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDatabase) GetUserPointDetails(username string) *PointDetails {
	time.Sleep(time.Second * 1)

	var pointData = PointDetails{}
	pointData, ok := mockPointDetails[username]
	if !ok {
		return nil
	}

	return &pointData
}

func (d *mockDatabase) SetupDatabase() error {
	return nil
}
