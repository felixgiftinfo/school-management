package auths

import (
	jwtto "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gopkg.in/square/go-jose.v2/jwt"
)

// Our handler struct
type AuthHandler struct {
	Logger zerolog.Logger
}

// decoded AWS IdToken JWT - just the bits we want
type AwsJwtClaims struct {
	Email string
	Uid   string
}

// Jwt Token structure
type JwtClaims struct {
	Email string `json:"email"`
	jwtto.StandardClaims
}

func NewAuthToken(logger zerolog.Logger) *AuthHandler {
	return &AuthHandler{Logger: logger}
}

func (h *AuthHandler) ParseIdToken(token string) (*AwsJwtClaims, error) {

	tok, err := jwt.ParseSigned(token)
	if err != nil {
		h.Logger.Error().
			Err(err).
			Msg("Error parsing JWT")
		return nil, err
	}

	var claims map[string]interface{}
	err = tok.UnsafeClaimsWithoutVerification(&claims)
	h.Logger.Debug().Interface("claims", claims).Send()
	if err != nil {
		h.Logger.Error().
			Err(err).
			Msg("Error parsing JWT")

		return nil, err
	}
	awsClaims := AwsJwtClaims{
		Email: claims["email"].(string),
		Uid:   claims["sub"].(string),
	}
	return &awsClaims, nil
}

// Create a JWT, usually for testing purposes
func (h *AuthHandler) CreateJwt(email string, sub string, issuer string) string {
	claims := JwtClaims{
		Email: email,
		StandardClaims: jwtto.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    issuer,
			Subject:   sub,
		},
	}

	signedToken, _ := jwtto.
		NewWithClaims(jwtto.SigningMethodHS256, claims).
		SignedString([]byte(uuid.NewString()))

	return signedToken
}
