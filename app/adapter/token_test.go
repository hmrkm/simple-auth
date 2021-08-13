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
		param       GetV1VerifyParams
		now         time.Time
		dbToken     string
		dbUserId    string
		dbExpiredAt time.Time
		dbErr       error
		expected    bool
		expectedErr error
	}{
		{
			"正常ケース",
			GetV1VerifyParams{
				Token: "token",
			},
			now,
			"token",
			"a",
			now.Add(1 * time.Millisecond),
			nil,
			true,
			nil,
		},
		{
			"DBに見つからない異常ケース",
			GetV1VerifyParams{
				Token: "token",
			},
			now,
			"",
			"",
			now,
			usecase.ErrNotFound,
			false,
			usecase.ErrNotFound,
		},
		{
			"有効期限切れの異常ケース",
			GetV1VerifyParams{
				Token: "token",
			},
			now,
			"token",
			"a",
			now.Add(-1 * time.Millisecond),
			nil,
			false,
			usecase.ErrTokenWasExpired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := usecase.NewMockStore(ctrl)
			sm.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
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
			ta := NewTokenAdapter(sm)

			actual, actualErr := ta.Verify(tc.param, tc.now)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Verify() isValid is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
