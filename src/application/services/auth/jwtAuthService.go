package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	a "ddgodeliv/domains/auth"
	u "ddgodeliv/domains/user"
)

type JwtAuthService struct {
	userRepository    u.IUserRepository
	sessionRepository a.ISessionRepository
	secret            *[]byte
}

func GetNewJwtAuthService(
	userRepository u.IUserRepository,
	sessionRepository a.ISessionRepository,
	secret *[]byte,
) *JwtAuthService {
	return &JwtAuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		secret:            secret,
	}
}

func (jas JwtAuthService) keyFunc(t *jwt.Token) (interface{}, error) {
	return *jas.secret, nil
}

func (jas JwtAuthService) generateToken(user *a.SessionUser, expiresAt time.Time) (string, error) {

	tokenStr, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256, a.GetNewPopulatedClaims(user, expiresAt),
	).SignedString(*jas.secret)

	return tokenStr, nil
}

func (jas JwtAuthService) getUnparsedToken(auth string, c *http.Cookie) string {
	if auth != "" {
		token, found := strings.CutPrefix(auth, "Bearer ")
		if found {
			return token
		}
	}
	if c != nil {
		return c.Value
	}
	return ""
}

func (jas JwtAuthService) ValidateAuth(
	authorization string, cookie *http.Cookie,
) (a.ISessionUser, error) {

	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return nil, fmt.Errorf("Missing Token")
	}

	claims := a.GetNewClaims()

	parsedToken, err := jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc)
	if err != nil || !parsedToken.Valid {
		go jas.sessionRepository.Delete(&claims.SessionUser)
		return nil, fmt.Errorf("Expired session")
	}

	return &claims.SessionUser, nil
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

	claimsUser := a.GetNewSessionUser()

	claimsUser.
		SetId(dbUser.GetId()).
		SetEmail(dbUser.GetEmail()).
		SetName(dbUser.GetName())

	token, err = jas.generateToken(claimsUser, expiresAt)
	if err != nil {
		return "", expiresAt, fmt.Errorf("Could not generate token")
	}

	jas.sessionRepository.Set(claimsUser, createdAt)

	return token, expiresAt, nil
}

func (jas JwtAuthService) Logout(authorization string, cookie *http.Cookie) error {
	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return fmt.Errorf("Missing Token")
	}

	claims := a.GetNewClaims()

	if _, err := jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc); err != nil {
		go jas.sessionRepository.Delete(&claims.SessionUser)
	}

	return nil
}
