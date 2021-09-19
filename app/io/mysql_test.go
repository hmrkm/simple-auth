package io

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		name        string
		conds       string
		params      interface{}
		dbId        string
		dbName      string
		expected    MockTable
		expectedErr error
	}{
		{

			"正常ケース",
			"id=?",
			"1",
			"1",
			"aaa",
			MockTable{
				Id:   "1",
				Name: "aaa",
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mysql, sqlMock := NewMysqlMock()

			actual := MockTable{}
			sqlMock.ExpectQuery(regexp.QuoteMeta(
				"SELECT * FROM `mock_tables` WHERE id=(?)",
			)).
				WithArgs(tc.params).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
					AddRow(tc.dbId, tc.dbName))

			actualErr := mysql.Find(&actual, tc.conds, tc.params)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Find() value is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Find() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
