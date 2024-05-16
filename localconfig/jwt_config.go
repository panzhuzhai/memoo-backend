package localconfig

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"memoo-backend/constants"
	"memoo-backend/middleware/jwt"
	"memoo-backend/middleware/sign/btc"
	"memoo-backend/middleware/sign/eth"
	"memoo-backend/serializer"
	"net/http"
	"strconv"
	"time"
)

var identityKey = "address"

type Login struct {
	Address   string `json:"address" binding:"required" example:"tb1pk7e88y59xje503pzsknn0z3mgpyvefagf2dlrkvlfxpnugy6shmq7qg7q4"`
	Message   string `json:"message" binding:"required" example:"1710501746195"`
	Signature string `json:"signature" binding:"required" example:"IGSo5ZI0H7xpDizzG/1zAYvljp98GoTGIoVZwMLdYAmkMPBwZyvtbgbE1Vs0SogS3juBko3W9BZ5XpGkJY8JEYc="`
	Chain     string `json:"chain" binding:"required" example:"Bitcoin"`
	PublicKey string `json:"publicKey" `
}

type User struct {
	Address string
	Chain   string
}

type Token struct {
	Token  string `json:"token"`
	Expire uint64 `json:"expire" example:"1710505373965"`
}

func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, serializer.Success(Token{
		Token:  token,
		Expire: uint64(expire.UnixNano() / int64(time.Millisecond)),
	}))
}

func LogoutResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, serializer.Success(nil))
}

func RefreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, serializer.Success(Token{
		Token:  token,
		Expire: uint64(expire.UnixNano() / int64(time.Millisecond)),
	}))
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, serializer.Fail(code, message))
}

func InitJwt() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "B2",
		SigningAlgorithm: "RS256",
		//Key:              []byte("secret key"),
		Timeout:     24 * time.Hour,
		MaxRefresh:  24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Address,
					"chain":     v.Chain,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Address: claims[identityKey].(string),
				Chain:   claims["chain"].(string),
			}
		},
		Authenticator:   authenticator,
		Unauthorized:    unauthorized,
		LoginResponse:   loginResponse,
		LogoutResponse:  logoutResponse,
		RefreshResponse: refreshResponse,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		PrivKeyFile:   "resources/jwt/private_key.pem",
		PubKeyFile:    "resources/jwt/public_key.pem",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginParams Login
	if err := c.ShouldBind(&loginParams); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	messageTime, err := strconv.Atoi(loginParams.Message)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, jwt.ErrMissingExpConfig
	}
	if time.Now().UnixNano()/int64(time.Millisecond)-int64(messageTime) >= AppConfig.JwtAttribute.SignValidTimeMs {
		return nil, jwt.ErrExpiredSign
	}
	address := loginParams.Address
	if constants.Ethereum == loginParams.Chain {
		verify, err := eth.SignVerify(loginParams.Message, loginParams.Signature, loginParams.Address)
		if err != nil || !verify {
			return nil, jwt.ErrInvalidSignature
		}
	} else if constants.Bitcoin == loginParams.Chain {
		verified, err := btc.VerifyWithChain(btc.SignedMessage{
			Address:   loginParams.Address,
			Message:   loginParams.Message,
			Signature: loginParams.Signature,
		}, BtcNetwork)
		if err != nil || !verified {
			return nil, jwt.ErrInvalidSignature
		}
		if err != nil {
			return nil, err
		}
	}

	return &User{
		Address: address,
		Chain:   loginParams.Chain,
	}, nil

}

func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, serializer.Success(Token{
		Token:  token,
		Expire: uint64(expire.UnixNano() / int64(time.Millisecond)),
	}))
}

func logoutResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, serializer.Success(nil))
}

func refreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, serializer.Success(Token{
		Token:  token,
		Expire: uint64(expire.UnixNano() / int64(time.Millisecond)),
	}))
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, serializer.Fail(code, message))
}
