package password_test

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/assert"
	"skilledin-green-skills-api/pkg/password"
	"testing"
)

// Mock rand.Reader that always returns an error
type mockRandReader struct {
	err error
}

func (r *mockRandReader) Read(_ []byte) (n int, err error) {
	return 0, r.err
}

func TestHash_Must_Success(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "test hash make and verify",
			args:    args{s: "secret"},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:    "test hash make and verify",
			args:    args{s: "12345"},
			want:    true,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := password.Hash{Stored: "", Supplied: tt.args.s}
			pwd, err := hash.Make(password.Parallelization)
			hash.Stored = pwd
			assert.Nil(t, err)
			valid, err := hash.Compare(password.Parallelization)
			assert.Nil(t, err)
			assert.Equalf(t, tt.want, valid, "Hash(%v)", tt.args.s)
		})
	}
}

func TestHash_Make_Error(t *testing.T) {
	t.Run("ERROR FROM READER", func(t *testing.T) {
		expectedErr := errors.New("random read error")
		mockRand := &mockRandReader{err: expectedErr}

		origRandReader := rand.Reader
		defer func() { rand.Reader = origRandReader }()
		rand.Reader = mockRand

		hash := password.Hash{Supplied: "password123"}
		secret, err := hash.Make(password.Parallelization)
		assert.NotNil(t, err)
		assert.Empty(t, secret)
		assert.Equal(t, err.Error(), expectedErr.Error())
	})
	t.Run("ERROR FROM HASH", func(t *testing.T) {
		hash := password.Hash{Supplied: "password123"}
		secret, err := hash.Make(-1)
		assert.NotNil(t, err)
		assert.Empty(t, secret)
	})
}

func TestHash_Compare_Error(t *testing.T) {
	t.Run("ERROR WHEN VALIDATE LEN", func(t *testing.T) {
		hash := password.Hash{Stored: "", Supplied: ""}
		valid, err := hash.Compare(password.Parallelization)
		assert.NotNil(t, err)
		assert.False(t, valid)
	})

	t.Run("ERROR WHEN HEX DECODE", func(t *testing.T) {
		hash := password.Hash{Stored: "1234.gibberish", Supplied: ""}
		valid, err := hash.Compare(password.Parallelization)
		assert.NotNil(t, err)
		assert.False(t, valid)
		assert.Equal(t, err.Error(), password.ErrorPasswordUnableToVerify.Error())
	})

	t.Run("ERROR WHEN HEX DECODE", func(t *testing.T) {
		hash := password.Hash{Stored: "1234.1234", Supplied: ""}
		valid, err := hash.Compare(-1)
		assert.NotNil(t, err)
		assert.False(t, valid)
		assert.Equal(t, err.Error(), password.ErrorPasswordUnableToVerify.Error())
	})
}
