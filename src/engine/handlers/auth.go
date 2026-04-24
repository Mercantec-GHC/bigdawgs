package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "user_id"

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := userIDFromToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserID(r *http.Request) (uint, error) {
	userID, ok := r.Context().Value(userIDKey).(uint)
	if !ok || userID == 0 {
		return 0, errors.New("missing authenticated user")
	}

	return userID, nil
}

func userIDFromToken(r *http.Request) (uint, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("missing authorization header")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return 0, errors.New("invalid authorization header")
	}

	tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, bearerPrefix))
	if tokenStr == "" {
		return 0, errors.New("missing bearer token")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return 0, errors.New("JWT_SECRET is not configured")
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}

		return []byte(secret), nil
	})
	if err != nil {
		return 0, fmt.Errorf("parse token: %w", err)
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claim, ok := claims["user_id"]
	if !ok {
		return 0, errors.New("token missing user_id claim")
	}

	userID, err := parseUserIDClaim(claim)
	if err != nil {
		return 0, fmt.Errorf("invalid user_id claim: %w", err)
	}

	return userID, nil
}

func parseUserIDClaim(claim any) (uint, error) {
	switch value := claim.(type) {
	case float64:
		if value <= 0 || value != float64(uint(value)) {
			return 0, errors.New("user_id must be a positive integer")
		}

		return uint(value), nil
	case string:
		n, err := strconv.ParseUint(value, 10, 64)
		if err != nil || n == 0 {
			return 0, errors.New("user_id must be a positive integer")
		}

		return uint(n), nil
	case json.Number:
		n, err := strconv.ParseUint(value.String(), 10, 64)
		if err != nil || n == 0 {
			return 0, errors.New("user_id must be a positive integer")
		}

		return uint(n), nil
	default:
		return 0, errors.New("unsupported user_id claim type")
	}
}
