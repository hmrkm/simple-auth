package adapter

import (
	"errors"
	"testing"
	"time"

	"github.com/hmrkm/simple-auth/usecase"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestVerifyToken(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name        string
		param       RequestVerify
		now         time.Time
		dbToken     string
		dbUserId    string
		dbExpiredAt time.Time
		dbErr       error
		dbEmail     string
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
			now.Add(1 * time.Millisecond),
			nil,
			"aaa@example.com",
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
			ResponseVerify{},
			usecase.ErrNotFound,
		},
		{
			"有効期限切れの異常ケース",
			RequestVerify{
				Token: "token",
			},
			now,
			"token",
			"a",
			now.Add(-1 * time.Millisecond),
			nil,
			"aaa@example.com",
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
				if tc.dbExpiredAt.After(tc.now) {
					sm.EXPECT().First(gomock.Any(), "id=?", tc.dbUserId).DoAndReturn(
						func(dest *usecase.User, cond string, param string) error {
							if tc.dbErr == nil {
								*dest = usecase.User{
									Id:    tc.dbUserId,
									Email: tc.dbEmail,
								}
							}
							return tc.dbErr
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
