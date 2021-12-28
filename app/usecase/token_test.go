package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/hmrkm/simple-auth/domain"
	"gorm.io/gorm"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestVerifyToken(t *testing.T) {
	now := time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local)
	feature := time.Date(2020, 6, 1, 17, 44, 13, 1, time.Local)
	past := time.Date(2020, 6, 1, 17, 44, 12, 999, time.Local)

	testCases := []struct {
		name        string
		token       string
		now         time.Time
		dbToken     string
		dbUserId    string
		dbExpiredAt time.Time
		dbErr       error
		dbEmail     string
		dbUserErr   error
		expected    domain.User
		expectedErr error
	}{
		{
			"正常ケース",
			"token",
			now,
			"token",
			"a",
			feature,
			nil,
			"aaa@example.com",
			nil,
			domain.User{
				Id:    "a",
				Email: "aaa@example.com",
			},
			nil,
		},
		{
			"DBに見つからない異常ケース",
			"token",
			now,
			"",
			"",
			now,
			domain.ErrNotFound,
			"aaa@example.com",
			nil,
			domain.User{},
			domain.ErrNotFound,
		},
		{
			"DBエラー異常ケース",
			"token",
			now,
			"",
			"",
			now,
			gorm.ErrInvalidDB,
			"aaa@example.com",
			nil,
			domain.User{},
			gorm.ErrInvalidDB,
		},
		{
			"DBエラー異常ケース2",
			"token",
			now,
			"",
			"",
			now,
			nil,
			"aaa@example.com",
			gorm.ErrInvalidDB,
			domain.User{},
			gorm.ErrInvalidDB,
		},
		{
			"有効期限切れの異常ケース",
			"token",
			now,
			"token",
			"a",
			past,
			nil,
			"aaa@example.com",
			nil,
			domain.User{},
			domain.ErrTokenWasExpired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := domain.NewMockStore(ctrl)
			sm.EXPECT().First(gomock.Any(), "token=?", tc.token).DoAndReturn(
				func(dest *domain.Token, cond string, params ...interface{}) error {
					if tc.dbErr == nil {
						*dest = domain.Token{
							Token:     tc.dbToken,
							UserId:    tc.dbUserId,
							ExpiredAt: tc.dbExpiredAt,
						}
					}
					return tc.dbErr
				},
			)
			if tc.dbErr == nil {
				if !tc.dbExpiredAt.Before(tc.now) {
					sm.EXPECT().First(gomock.Any(), "id=?", tc.dbUserId).DoAndReturn(
						func(dest *domain.User, cond string, param string) error {
							if tc.dbUserErr == nil {
								*dest = domain.User{
									Id:    tc.dbUserId,
									Email: tc.dbEmail,
								}
							}
							return tc.dbUserErr
						},
					)
				}
			}

			ta := NewToken(sm)

			actual, actualErr := ta.Verify(tc.token, tc.now)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Verify() domain.User is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
