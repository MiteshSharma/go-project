package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"

	jwt "github.com/dgrijalva/jwt-go"
)

func (a *App) SignToken(userID int, roles []model.Role) (string, *model.AppError) {
	currentTime := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":      userID,
		"currentTime": currentTime,
		"roles":       roles,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(a.Config.AuthConfig.HmacSecret))

	if err != nil {
		a.Log.Debug("Token signing got error", logger.Error(err))
		return "", model.NewAppError("Token signing got error", http.StatusInternalServerError)
	}

	result := a.Repository.User().GetSession(userID)
	var existingUserSession model.UserSession
	if result.Err == nil {
		existingUserSession = result.Data.(model.UserSession)
	}
	session := &model.UserSession{
		UserID: userID,
		Token:  tokenString,
	}
	if existingUserSession == (model.UserSession{}) {
		a.Repository.User().CreateSession(session)
	} else {
		a.Repository.User().UpdateSession(session)
	}

	return tokenString, nil
}

func (a *App) ResetToken(userID int) {
	a.Repository.User().DeleteSession(userID)
}

func (a *App) VerifyAndParseToken(tokenString string) (*model.UserSession, *model.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.Config.AuthConfig.HmacSecret), nil
	})
	if err != nil {
		return nil, model.NewAppError(err.Error(), http.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userID"].(float64))
		session, err := a.getSession(userID)
		if err == nil {
			return session, nil
		}
		return nil, model.NewAppError("user not authorized", http.StatusUnauthorized)
	}
	return nil, model.NewAppError("incorrect token", http.StatusBadRequest)
}

func (a *App) getSession(userID int) (*model.UserSession, *model.AppError) {
	sessionResult := a.Repository.User().GetSession(userID)
	if sessionResult.Err != nil {
		return nil, sessionResult.Err
	}
	session := sessionResult.Data.(*model.UserSession)
	return session, nil
}
