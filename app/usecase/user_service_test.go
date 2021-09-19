package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
)

func TestVerify(t *testing.T) {
	hashedPasswd := CreateHash("passwd")
	testCases := []struct {
		name         string
		email        string
		password     string
		dbId         string
		dbEmail      string
		dbPassword   string
		dbErr        error
		expectedUser User
		expectedErr  error
	}{
		{
			"正常ケース",
			"aaa",
			"passwd",
			"a",
			"aaa@example.com",
			hashedPasswd,
			nil,
			User{
				Id:       "a",
				Email:    "aaa@example.com",
				Password: hashedPasswd,
			},
			nil,
		},
		{
			"DBに見つからない異常ケース",
			"aaa",
			"passwd",
			"",
			"",
			"",
			ErrNotFound,
			User{},
			ErrNotFound,
		},
		{
			"DBエラーの異常ケース",
			"aaa",
			"passwd",
			"",
			"",
			"",
			gorm.ErrInvalidDB,
			User{},
			gorm.ErrInvalidDB,
		},
		{
			"パスワードが一致しない異常ケース",
			"aaa@example.com",
			"password",
			"a",
			"aaa@example.com",
			hashedPasswd,
			nil,
			User{},
			ErrInvalidPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := NewMockStore(ctrl)
			sm.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
				func(dest *User, cond string, param string) error {
					if tc.dbErr == nil {
						*dest = User{
							Id:       tc.dbId,
							Email:    tc.dbEmail,
							Password: tc.dbPassword,
						}
					}
					return tc.dbErr
				},
			)
			if tc.dbErr != nil {
				sm.EXPECT().IsNotFoundError(tc.dbErr).DoAndReturn(
					func(err error) bool {
						return errors.Is(ErrNotFound, err)
					},
				)
			}
			us := NewUserService(sm)

			actualUser, actualErr := us.Verify(tc.email, tc.password)

			if diff := cmp.Diff(tc.expectedUser, actualUser); diff != "" {
				t.Errorf("Verify() user is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
