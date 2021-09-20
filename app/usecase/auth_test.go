package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/hmrkm/simple-auth/domain"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestVerifyAuth(t *testing.T) {
	hashedPasswd := domain.CreateHash("passwd")
	now := time.Now()
	user := domain.User{
		Id:       "a",
		Email:    "aaa@example.com",
		Password: hashedPasswd,
	}
	createErr := errors.New("create error")
	testCases := []struct {
		name            string
		email           string
		password        string
		now             time.Time
		tokenExpireHour int
		isValid         bool
		user            domain.User
		verifyErr       error
		token           domain.Token
		CreateErr       error
		expected        domain.Token
		expectedErr     error
	}{
		{
			"正常ケース",
			"aaa@example.com",
			"passwd",
			now,
			1,
			true,
			user,
			nil,
			domain.Token{
				Token:     "token",
				UserId:    "a",
				ExpiredAt: now.Add(1 * time.Hour),
			},
			nil,
			domain.Token{
				Token:     "token",
				UserId:    "a",
				ExpiredAt: now.Add(1 * time.Hour),
			},
			nil,
		},
		{
			"ユーザー認証異常ケース1",
			"aaa@example.com",
			"passwd",
			now,
			1,
			true,
			domain.User{},
			domain.ErrInvalidPassword,
			domain.Token{},
			nil,
			domain.Token{},
			domain.ErrInvalidPassword,
		},
		{
			"ユーザー認証異常ケース2",
			"aaa@example.com",
			"passwd",
			now,
			1,
			false,
			domain.User{},
			domain.ErrInvalidVerify,
			domain.Token{},
			nil,
			domain.Token{},
			domain.ErrInvalidVerify,
		},
		{
			"トークン作成失敗の異常ケース",
			"aaa@example.com",
			"passwd",
			now,
			1,
			true,
			domain.User{
				Email:    "aaa@example.com",
				Password: hashedPasswd,
			},
			nil,
			domain.Token{},
			createErr,
			domain.Token{},
			createErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usm := domain.NewMockUserService(ctrl)
			usm.EXPECT().Verify(gomock.Any(), gomock.Any()).Return(tc.user, tc.verifyErr)
			tsm := domain.NewMockTokenService(ctrl)
			if tc.verifyErr == nil {
				tsm.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.token, tc.CreateErr)
			}

			ta := NewAuth(usm, tsm)

			actual, actualErr := ta.Verify(tc.email, tc.password, tc.now, tc.tokenExpireHour)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Verify() domain.Token is missmatch :%s", diff)
			}
			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
