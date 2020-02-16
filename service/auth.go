package service

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	api "github.com/beforesecond/gqlgen-todos/api"
	"github.com/beforesecond/gqlgen-todos/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Auth service
func Auth(g *echo.Group) {
	g.POST("", authTokenHandler)
	g.POST("/register", authRegisterHandler)
	g.POST("/revoke", authRevokeHandler, verifyAccessTokenMiddleware)
}

type authRequest struct {
	GrantType    grantType `json:"grant_type"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"` // unit: seconds
	RefreshToken string `json:"refresh_token,omitempty"`
	UID          string `json:"uid"`
}

type grantType string

// Grant Type
const (
	grantTypePassword     = "password"
	grantTypeRefreshToken = "refresh_token"
)

type errorString struct {
	s string
}

func AuthTokenHandlerByGraphQL(input models.InputLogin) (*models.AuthResponse, error) {
	var body authRequest
	// if err := c.Bind(&body); err != nil {
	// 	return c.String(http.StatusBadRequest, "Bad Request")
	// }

	if input.GrantType == grantTypePassword {
		// handle password grant type => return refresh token
		user, err := api.FindUser(input.Username, input.Password)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if user == nil {
			// user or password wrong = unauthorized
			err = errors.New("Unauthorized")
			return nil, err
		}
		refreshToken, err := generateRefreshToken(user.ID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		accessToken, err := generateAccessToken(user.ID, accessTokenDuration)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &models.AuthResponse{
			accessToken,
			"bearer",
			int(accessTokenDuration.Seconds()),
			refreshToken,
			user.ID,
		}, nil
	}
	if input.GrantType == grantTypeRefreshToken {
		// handle refresh token grant type => return access token

		// get user id from context
		claims, err := validateToken(body.RefreshToken)
		if err != nil {
			return nil, err
		}
		// verify refresh token in database
		if ok, err := api.ValidateToken(body.RefreshToken, claims.ID, refreshTokenDuration); !ok {
			if err != nil {
				log.Println(err)
				return nil, err
			}
			return nil, err
		}

		accessToken, err := generateAccessToken(claims.ID, accessTokenDuration)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &models.AuthResponse{
			accessToken,
			"bearer",
			int(accessTokenDuration.Seconds()),
			"",
			claims.ID,
		}, nil
	}
	err := errors.New("Unauthorized")
	return nil, err
}

func authTokenHandler(c echo.Context) error {
	var body authRequest
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if body.GrantType == grantTypePassword {
		// handle password grant type => return refresh token
		user, err := api.FindUser(body.Username, body.Password)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		if user == nil {
			// user or password wrong = unauthorized
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		refreshToken, err := generateRefreshToken(user.ID)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		accessToken, err := generateAccessToken(user.ID, accessTokenDuration)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, AuthResponse{
			accessToken,
			"bearer",
			int64(accessTokenDuration.Seconds()),
			refreshToken,
			user.ID,
		})
	}
	if body.GrantType == grantTypeRefreshToken {
		// handle refresh token grant type => return access token

		// get user id from context
		claims, err := validateToken(body.RefreshToken)
		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		// verify refresh token in database
		if ok, err := api.ValidateToken(body.RefreshToken, claims.ID, refreshTokenDuration); !ok {
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, "Internal Server Error")
			}
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		accessToken, err := generateAccessToken(claims.ID, accessTokenDuration)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, AuthResponse{
			accessToken,
			"bearer",
			int64(accessTokenDuration.Seconds()),
			"",
			claims.ID,
		})
	}

	return c.String(http.StatusUnauthorized, "Unauthorized")
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthRegisterGraphQLHandler(input models.InputUser) (*models.UserModel, error) {
	var err error

	var user models.UserModel
	user.Username = input.Username
	user.SetPassword(input.Password)

	err = api.SaveUser(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func authRegisterHandler(c echo.Context) error {
	var body registerRequest
	var err error
	if err = c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	var user models.UserModel
	user.Username = body.Username
	user.SetPassword(body.Password)

	err = api.SaveUser(&user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.String(http.StatusCreated, "Created")
}

type tokenClaim struct {
	ID   string    `json:"id"`
	Type tokenType `json:"type"`
	jwt.StandardClaims
}

type tokenType int

// Token Type
const (
	_                               = iota
	TokenTypeRefreshToken tokenType = iota
	TokenTypeAccessToken
)

const accessTokenDuration = time.Duration(time.Minute * 5)
const refreshTokenDuration = time.Duration(time.Hour * 24 * 7)

// Token Errors
var (
	ErrInvalidToken = errors.New("token: invalid token")
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	key, err := ioutil.ReadFile("key/key.rsa")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Fatal(err)
	}

	key, err = ioutil.ReadFile("key/key.pub")
	if err != nil {
		log.Fatal(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		log.Fatal(err)
	}
}

type authRevokeRequest struct {
	Token string `json:"token"`
}

func authRevokeHandler(c echo.Context) error {
	var body authRevokeRequest
	var err error
	if err = c.Bind(&body); err != nil || body.Token == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err = api.DeleteToken(body.Token); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getTokenFromHeader(c echo.Context) string {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	token = strings.TrimSpace(token)
	if token == "" || len(token) < 8 || strings.ToLower(token[:7]) != "bearer " {
		return ""
	}
	token = strings.TrimSpace(token[7:])
	return token
}

func validateToken(token string) (*tokenClaim, error) {
	tok, err := jwt.ParseWithClaims(token, &tokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Check is token use correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return secret for this signing method
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tok.Claims.(*tokenClaim); ok && tok.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}

func verifyAccessTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := getTokenFromHeader(c)
		claims, err := validateToken(token)
		if err != nil || claims.Type != TokenTypeAccessToken {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		// set user id to context
		c.Set("userID", claims.ID)
		return next(c)
	}
}

func generateToken(id string, expiresIn time.Duration, tokenType tokenType) (string, error) {
	expiresAt := int64(0) // not expires
	now := time.Now()
	if expiresIn > 0 {
		expiresAt = now.Add(expiresIn).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenClaim{
		id,
		tokenType,
		jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	})
	return token.SignedString(privateKey)
}

func generateRefreshToken(id string) (string, error) {
	token, err := generateToken(id, 0, TokenTypeRefreshToken)
	if err != nil {
		return "", err
	}
	if err = api.CreateToken(token, id); err != nil {
		return "", err
	}
	return token, nil
}

func generateAccessToken(id string, expiresIn time.Duration) (string, error) {
	token, err := generateToken(id, expiresIn, TokenTypeAccessToken)
	if err != nil {
		return "", err
	}
	return token, nil
}
