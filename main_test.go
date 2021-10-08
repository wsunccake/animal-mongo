package main

import (
	"animal-mongo/util"
	"testing"
)

func TestMain(m *testing.M) {
	util.ConnectDatabase(MONGODB_URI)
	defer util.DisconnectDatabase()
}
