package domain

import "testing"

func TestUserTableName(t *testing.T) {
	tableName := User{}.TableName()

	if tableName != "users" {
		t.Errorf("TableName is not users")

	}
}
