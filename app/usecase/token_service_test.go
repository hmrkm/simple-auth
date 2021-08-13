package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestCreate(t *testing.T) {
	hashedPasswd := CreateHash("passwd")
	user := User{
		"a",
		"aaa",
		hashedPasswd,
	}
	now := time.Now()
	dbErr := errors.New("db error")
	testCases := []struct {
		name        string
		user        User
		now         time.Time
		expireHour  int
		dbToken     string
		dbUserId    string
		dbExpiredAt time.Time
		dbErr       error
		expected    Token
		expectedErr error
	}{
		{
			"正常ケース",
			user,
			now,
			1,
			"token",
			"a",
			now.Add(1 * time.Hour),
			nil,
			Token{
				"token",
				"a",
				now.Add(1 * time.Hour),
			},
			nil,
		},
		{
			"作成失敗異常ケース",
			user,
			now,
			1,
			"token",
			"a",
			now.Add(1 * time.Hour),
			dbErr,
			Token{},
			dbErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := NewMockStore(ctrl)
			sm.EXPECT().Create(gomock.Any()).DoAndReturn(
				func(target *Token) error {
					if tc.dbErr == nil {
						*target = Token{
							Token:     tc.dbToken,
							UserId:    tc.dbUserId,
							ExpiredAt: tc.dbExpiredAt,
						}
					}
					return tc.dbErr
				},
			)
			us := NewTokenService(sm)

			actual, actualErr := us.Create(tc.user, tc.now, tc.expireHour)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Create() isValid is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Create() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
