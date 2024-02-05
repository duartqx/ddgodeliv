package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	e "ddgodeliv/common/errors"
	a "ddgodeliv/domains/auth"
	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

type customClaims struct {
	a.SessionUser
	jwt.RegisteredClaims
}

type JwtAuthService struct {
	userRepository    u.IUserRepository
	driverRepository  d.IDriverRepository
	sessionRepository a.ISessionRepository
	secret            *[]byte
}

func GetNewJwtAuthService(
	userRepository u.IUserRepository,
	driverRepository d.IDriverRepository,
	sessionRepository a.ISessionRepository,
	secret *[]byte,
) *JwtAuthService {
	return &JwtAuthService{
		userRepository:    userRepository,
		driverRepository:  driverRepository,
		sessionRepository: sessionRepository,
		secret:            secret,
	}
}

func (jas JwtAuthService) claims() *customClaims {
	return &customClaims{
		SessionUser: *a.GetNewSessionUser(), RegisteredClaims: jwt.RegisteredClaims{},
	}
}

func (jas JwtAuthService) keyFunc(t *jwt.Token) (interface{}, error) {
	return *jas.secret, nil
}

func (jas JwtAuthService) generateToken(user *a.SessionUser, expiresAt time.Time) (string, error) {

	claims := &customClaims{
		SessionUser: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(*jas.secret)
	if err != nil {
		return "", fmt.Errorf("Bad secret key")
	}

	return tokenStr, nil
}

func (jas JwtAuthService) getUnparsedToken(
	authorization string, cookie *http.Cookie,
) string {
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

func (jas JwtAuthService) ValidateAuth(
	authorization string, cookie *http.Cookie,
) (a.ISessionUser, error) {

	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return nil, fmt.Errorf("Missing Token")
	}

	claims := jas.claims()

	parsedToken, err := jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc)
	if err != nil || !parsedToken.Valid {
		jas.sessionRepository.Delete(&claims.SessionUser)

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

	driver, err := jas.driverRepository.FindByUserId(claimsUser.GetId())
	if err != nil && !errors.Is(err, e.NotFoundError) {
		return token, expiresAt, err
	}

	if driver != nil {
		claimsUser.SetDriver(driver)
	}

	token, err = jas.generateToken(claimsUser, expiresAt)
	if err != nil {
		return "", expiresAt, fmt.Errorf("Could not generate token")
	}

	if err := jas.sessionRepository.Set(claimsUser, createdAt); err != nil {
		return "", expiresAt, err
	}

	return token, expiresAt, nil
}

func (jas JwtAuthService) Logout(authorization string, cookie *http.Cookie) error {
	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return fmt.Errorf("Missing Token")
	}

	claims := jas.claims()

	jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc)
	if err := jas.sessionRepository.Delete(&claims.SessionUser); err != nil {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}
