package adapter

import (
	"errors"
	"testing"
	"time"

	"github.com/hmrkm/simple-auth/usecase"
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
		param       RequestVerify
		now         time.Time
		dbToken     string
		dbUserId    string
		dbExpiredAt time.Time
		dbErr       error
		dbEmail     string
		dbUserErr   error
		expected    ResponseVerify
		expectedErr error
	}{
		{
			"正常ケース",
			RequestVerify{
				Token: "token",
			},
			now,
			"token",
			"a",
			feature,
			nil,
			"aaa@example.com",
			nil,
			ResponseVerify{
				User: VerifyUser{
					Id:    "a",
					Email: "aaa@example.com",
				},
			},
			nil,
		},
		{
			"DBに見つからない異常ケース",
			RequestVerify{
				Token: "token",
			},
			now,
			"",
			"",
			now,
			usecase.ErrNotFound,
			"aaa@example.com",
			nil,
			ResponseVerify{},
			usecase.ErrNotFound,
		},
		{
			"DBエラー異常ケース",
			RequestVerify{
				Token: "token",
			},
			now,
			"",
			"",
			now,
			gorm.ErrInvalidDB,
			"aaa@example.com",
			nil,
			ResponseVerify{},
			gorm.ErrInvalidDB,
		},
		{
			"DBエラー異常ケース2",
			RequestVerify{
				Token: "token",
			},
			now,
			"",
			"",
			now,
			nil,
			"aaa@example.com",
			gorm.ErrInvalidDB,
			ResponseVerify{},
			gorm.ErrInvalidDB,
		},
		{
			"有効期限切れの異常ケース",
			RequestVerify{
				Token: "token",
			},
			now,
			"token",
			"a",
			past,
			nil,
			"aaa@example.com",
			nil,
			ResponseVerify{},
			usecase.ErrTokenWasExpired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := usecase.NewMockStore(ctrl)
			sm.EXPECT().First(gomock.Any(), "token=?", tc.param.Token).DoAndReturn(
				func(dest *usecase.Token, cond string, param string) error {
					if tc.dbErr == nil {
						*dest = usecase.Token{
							Token:     tc.dbToken,
							UserId:    tc.dbUserId,
							ExpiredAt: tc.dbExpiredAt,
						}
					}
					return tc.dbErr
				},
			)
			if tc.dbErr != nil {
				sm.EXPECT().IsNotFoundError(tc.dbErr).DoAndReturn(
					func(err error) bool {
						return errors.Is(usecase.ErrNotFound, err)
					},
				)
			} else {
				if !tc.dbExpiredAt.Before(tc.now) {
					sm.EXPECT().First(gomock.Any(), "id=?", tc.dbUserId).DoAndReturn(
						func(dest *usecase.User, cond string, param string) error {
							if tc.dbUserErr == nil {
								*dest = usecase.User{
									Id:    tc.dbUserId,
									Email: tc.dbEmail,
								}
							}
							return tc.dbUserErr
						},
					)
				}
			}

			ta := NewTokenAdapter(sm)

			actual, actualErr := ta.Verify(tc.param, tc.now)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Verify() ResponseVerify is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
