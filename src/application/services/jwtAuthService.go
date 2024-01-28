package services

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	u "ddgodeliv/domains/user"
	r "ddgodeliv/infrastructure/repository"
)

type ClaimsUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type customClaims struct {
	ClaimsUser
	jwt.RegisteredClaims
}

type JwtAuthService struct {
	userRepository    r.IUserRepository
	sessionRepository r.ISessionRepository
	secret            *[]byte
}

func GetNewJwtAuthService(userRepository r.IUserRepository, secret *[]byte, sessionRepository r.ISessionRepository) *JwtAuthService {
	return &JwtAuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		secret:            secret,
	}
}

func (jas JwtAuthService) claims() *customClaims {
	return &customClaims{ClaimsUser: ClaimsUser{}, RegisteredClaims: jwt.RegisteredClaims{}}
}

func (jas JwtAuthService) keyFunc(t *jwt.Token) (interface{}, error) {
	return *jas.secret, nil
}

func (jas JwtAuthService) generateToken(user *ClaimsUser, expiresAt time.Time) (string, error) {

	claims := &customClaims{
		ClaimsUser:       *user,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expiresAt)},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(*jas.secret)
	if err != nil {
		return "", fmt.Errorf("Bad secret key")
	}

	return tokenStr, nil
}

func (jas JwtAuthService) getUnparsedToken(authorization string, cookie *http.Cookie) string {
	if authorization != "" {
		token, found := strings.CutPrefix(authorization, "Bearer ")
		if found {
			return token
		}
	}
	if cookie != nil {
		return cookie.Value
	}
	return ""
}

func (jas JwtAuthService) ValidateAuth(authorization string, cookie *http.Cookie) (*ClaimsUser, error) {

	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return nil, fmt.Errorf("Missing Token")
	}

	claims := jas.claims()

	parsedToken, err := jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc)
	if err != nil || !parsedToken.Valid {
		jas.sessionRepository.Delete(unparsedToken)

		return nil, fmt.Errorf("Expired session")
	}

	return &claims.ClaimsUser, nil
}

func (jas JwtAuthService) Login(user u.IUser) (token string, expiresAt time.Time, err error) {

	if user.GetEmail() == "" || user.GetPassword() == "" {
		return token, expiresAt, fmt.Errorf("Invalid Email or Password")
	}

	dbUser, err := jas.userRepository.FindByEmail(user.GetEmail())
	if err != nil {
		return token, expiresAt, fmt.Errorf("Invalid Email")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.GetPassword()), []byte(user.GetPassword()),
	); err != nil {
		return token, expiresAt, fmt.Errorf("Invalid Password")
	}

	createdAt := time.Now()
	expiresAt = createdAt.Add(time.Hour * 12)

	token, err = jas.generateToken(
		&ClaimsUser{
			Id:    dbUser.GetId(),
			Email: dbUser.GetEmail(),
			Name:  dbUser.GetName(),
		},
		expiresAt,
	)
	if err != nil {
		return "", expiresAt, fmt.Errorf("Could not generate token")
	}

	if err := jas.sessionRepository.Set(token, createdAt, dbUser.GetId()); err != nil {
		return "", expiresAt, err
	}

	return token, expiresAt, nil
}

func (jas JwtAuthService) Logout(authorization string, cookie *http.Cookie) error {
	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return fmt.Errorf("Missing Token")
	}

	if err := jas.sessionRepository.Delete(unparsedToken); err != nil {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}
