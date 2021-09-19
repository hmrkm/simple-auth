package adapter

import (
	"errors"
	"testing"
	"time"

	"github.com/hmrkm/simple-auth/usecase"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestVerifyAuth(t *testing.T) {
	hashedPasswd := usecase.CreateHash("passwd")
	now := time.Now()
	user := usecase.User{
		Id:       "a",
		Email:    "aaa@example.com",
		Password: hashedPasswd,
	}
	createErr := errors.New("create error")
	testCases := []struct {
		name            string
		req             RequestAuth
		now             time.Time
		tokenExpireHour int
		isValid         bool
		user            usecase.User
		verifyErr       error
		token           usecase.Token
		CreateErr       error
		expected        ResponseAuth
		expectedErr     error
	}{
		{
			"正常ケース",
			RequestAuth{
				Email:    "aaa@example.com",
				Password: "passwd",
			},
			now,
			1,
			true,
			user,
			nil,
			usecase.Token{
				Token:     "token",
				UserId:    "a",
				ExpiredAt: now.Add(1 * time.Hour),
			},
			nil,
			ResponseAuth{
				Token:     "token",
				ExpiredAt: int(now.Add(1*time.Hour).UnixNano() / 1000),
			},
			nil,
		},
		{
			"ユーザー認証異常ケース1",
			RequestAuth{
				Email:    "aaa@example.com",
				Password: "passwd",
			},
			now,
			1,
			true,
			usecase.User{},
			usecase.ErrInvalidPassword,
			usecase.Token{},
			nil,
			ResponseAuth{},
			usecase.ErrInvalidPassword,
		},
		{
			"ユーザー認証異常ケース2",
			RequestAuth{
				Email:    "aaa@example.com",
				Password: "passwd",
			},
			now,
			1,
			false,
			usecase.User{},
			usecase.ErrInvalidVerify,
			usecase.Token{},
			nil,
			ResponseAuth{},
			usecase.ErrInvalidVerify,
		},
		{
			"トークン作成失敗の異常ケース",
			RequestAuth{
				Email:    "aaa@example.com",
				Password: "passwd",
			},
			now,
			1,
			true,
			usecase.User{
				Email:    "aaa@example.com",
				Password: hashedPasswd,
			},
			nil,
			usecase.Token{},
			createErr,
			ResponseAuth{},
			createErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usm := usecase.NewMockUserService(ctrl)
			usm.EXPECT().Verify(gomock.Any(), gomock.Any()).Return(tc.user, tc.verifyErr)
			tsm := usecase.NewMockTokenService(ctrl)
			if tc.verifyErr == nil {
				tsm.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.token, tc.CreateErr)
			}

			ta := NewAuthAdapter(usm, tsm)

			actual, actualErr := ta.Verify(tc.req, tc.now, tc.tokenExpireHour)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Verify() isValid is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
