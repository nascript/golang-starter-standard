package token_test

import (
	"github.com/stretchr/testify/assert"
	"skilledin-green-skills-api/pkg/token"
	"testing"
	"time"
)

func TestJSONWebToken_ClaimJWTToken(t *testing.T) {
	type fields struct {
		Issuer    string
		SecretKey []byte
		Payload   interface{}
		IssuedAt  time.Time
		ExpiredAt time.Time
	}

	issuedAt, _ := time.Parse("",
		"2022-12-05 17:57:44.321843 +0800 WITA m=+25.737606459")
	expiredAt, _ := time.Parse("",
		"2022-12-06 17:57:44.321851 +0800 WITA m=+86425.737614876")

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "NEW JWT TEST SHOULD SUCCESS",
			fields: fields{
				Issuer: "BECOOP_TEST",
				Payload: map[string]string{
					"data": "hello world",
				},
				IssuedAt:  issuedAt,
				ExpiredAt: expiredAt,
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJCRUNPT1BfVEVTVCIsImV4cCI6LTYyMTM1NTk2ODAwLCJpYXQiOi02MjEzNTU5NjgwMCwicGF5bG9hZCI6eyJkYXRhIjoiaGVsbG8gd29ybGQifX0.XEO-61nJMVhUJJXAMpoCa1jzxiIJC12i6CnFMDzIwtg",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jwt := &token.JSONWebToken{
				Issuer:    tt.fields.Issuer,
				SecretKey: tt.fields.SecretKey,
				IssuedAt:  tt.fields.IssuedAt,
				ExpiredAt: tt.fields.ExpiredAt,
			}
			got, err := jwt.Claim(tt.fields.Payload)
			if !tt.wantErr(t, err, "ClaimJWTToken()") {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
