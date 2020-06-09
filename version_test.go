package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		name    string
		ver     string
		wantErr bool
	}{
		{
			name: "happy path",
			ver:  "1.2.3",
		},
		{
			name: "happy path with pre suffix",
			ver:  "1.2.3_alpha1",
		},
		{
			name: "happy path with post suffix",
			ver:  "1.2.3-r1",
		},
		{
			name: "happy path with 2 pre suffixes",
			ver:  "0.1.0_alpha_pre2",
		},
		{
			name: "happy path with alphabets",
			ver:  "1.0b",
		},
		{
			name:    "sad path",
			ver:     "1.0bc",
			wantErr: true,
		},
		{
			name:    "invalid symbol",
			ver:     "1.0!",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewVersion(tt.ver)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" vs "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.Compare(b)

			var expected int
			switch tt.expected {
			case ">":
				expected = apkVersionGreater
			case "<":
				expected = apkVersionLess
			case "=":
				expected = apkVersionEqual
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestVersion_CompareMultipleTimes(t *testing.T) {
	a, _ := NewVersion("1.2.3")
	b, _ := NewVersion("1.2.3")

	got := a.Compare(b)
	assert.Equal(t, 0, got)

	b, _ = NewVersion("1.2.3_pre1")
	got = a.Compare(b)
	assert.True(t, got > 0)

	b, _ = NewVersion("1.2.3-r1")
	got = a.Compare(b)
	assert.True(t, got < 0)
}

func TestVersion_Equal(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" is equal to "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.Equal(b)

			var expected bool
			switch tt.expected {
			case ">", "<":
				expected = false
			case "=":
				expected = true
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestVersion_GreaterThan(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" is greater than "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.GreaterThan(b)

			var expected bool
			switch tt.expected {
			case "<", "=":
				expected = false
			case ">":
				expected = true
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestVersion_LessThan(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" is less than "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.LessThan(b)

			var expected bool
			switch tt.expected {
			case ">", "=":
				expected = false
			case "<":
				expected = true
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}
