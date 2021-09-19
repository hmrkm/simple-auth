package io

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/hmrkm/simple-auth/usecase"
	"gorm.io/gorm"
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

func TestFirst(t *testing.T) {
	testCases := []struct {
		name        string
		conds       string
		params      interface{}
		dbId        string
		dbName      string
		dbErr       error
		expected    MockTable
		expectedErr error
	}{
		{

			"正常ケース",
			"id=?",
			"1",
			"1",
			"aaa",
			nil,
			MockTable{
				Id:   "1",
				Name: "aaa",
			},
			nil,
		},
		{

			"レコードが見つからない異常ケース",
			"id=?",
			"1",
			"1",
			"aaa",
			gorm.ErrRecordNotFound,
			MockTable{},
			usecase.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mysql, sqlMock := NewMysqlMock()

			actual := MockTable{}
			if tc.dbErr == nil {
				sqlMock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `mock_tables` WHERE id=(?)",
				)).
					WithArgs(tc.params).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
						AddRow(tc.dbId, tc.dbName))
			} else {
				sqlMock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `mock_tables` WHERE id=(?)",
				)).
					WithArgs(tc.params).
					WillReturnError(tc.dbErr)
			}

			actualErr := mysql.First(&actual, tc.conds, tc.params)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("First() value is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("First() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name        string
		mockTable   MockTable
		expectedErr error
	}{
		{

			"正常ケース",
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

			sqlMock.ExpectBegin()
			sqlMock.ExpectExec(regexp.QuoteMeta(
				"INSERT INTO `mock_tables` (`id`,`name`) VALUES (?,?)",
			)).
				WithArgs(tc.mockTable.Id, tc.mockTable.Name).
				WillReturnResult(sqlmock.NewResult(1, 1))
			sqlMock.ExpectCommit()
			sqlMock.ExpectClose()

			actualErr := mysql.Create(tc.mockTable)

			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Create() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}

func TestIsNotFoundError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{

			"正常ケース",
			gorm.ErrRecordNotFound,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mysql, _ := NewMysqlMock()

			actual := mysql.IsNotFoundError(tc.err)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("IsNotFoundError() value is missmatch :%s", diff)
			}
		})
	}
}
