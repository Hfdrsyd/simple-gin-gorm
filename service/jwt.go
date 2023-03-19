package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(UserID uint64, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (uint64, error)
}

type jwtCustomClaim struct {
	UserID uint64 `json:"id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer: "Admin",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "jwt_secret_key"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID uint64, role string) (string, error) {
	claims := jwtCustomClaim{
		UserID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
			Issuer: j.issuer,
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil{
		return "", err
	}
	return tx, nil
}

func (j *jwtService) parseToken(t_ *jwt.Token)(any, error){
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok{
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string)(*jwt.Token, error){
	parsedToken, err := jwt.Parse(token, j.parseToken)
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func (j *jwtService) GetUserIDByToken(token string)(uint64, error){
	t_Token, err := j.ValidateToken(token)
	if err != nil{
		return 0, err
	}
	claims := t_Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["id"])
	userID, _ := strconv.ParseUint(id,10,64)
	return userID, nil
}