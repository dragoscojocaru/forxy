package request

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetHost(t *testing.T) {

	testCases := []struct {
		url      string
		expected string
		err      error
	}{
		{
			url:      "https://forxy.io",
			expected: "forxy.io",
			err:      nil,
		},
		{
			url:      "https://forxy.io/",
			expected: "forxy.io",
			err:      nil,
		},
		{
			url:      "https://forxy.io/forxy/path",
			expected: "forxy.io",
			err:      nil,
		},
		{
			url:      "//forxy.io/",
			expected: "forxy.io",
			err:      nil,
		},
		{
			url:      "123!@#",
			expected: "",
			err:      *new(error),
		},
		{
			url:      "forxy.io/path",
			expected: "",
			err:      *new(error),
		},
	}

	for _, test := range testCases {
		t.Run("test", func(t *testing.T) {
			t.Parallel()
			host, err := GetHost(test.url)
			assert.Equal(t, test.expected, host)
			if test.err == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
