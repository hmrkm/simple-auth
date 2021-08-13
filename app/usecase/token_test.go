package usecase

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestIsValid(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name      string
		ExpiredAt time.Time
		now       time.Time
		expected  bool
	}{
		{
			"正常ケース",
			now.Add(1 * time.Hour),
			now,
			true,
		},
		{
			"異常ケース",
			now.Add(-1 * time.Hour),
			now,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token := Token{
				ExpiredAt: tc.ExpiredAt,
			}

			actual := token.IsValid(tc.now)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("IsValid() is missmatch :%s", diff)
			}
		})
	}
}
