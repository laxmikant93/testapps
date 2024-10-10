package handler

import (
    "encoding/json"
    
    "net/http"
    "goauth/user"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("your-secret-key")

// In-memory user store
var users = map[string]string{}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func SignUp(w http.ResponseWriter, r *http.Request) {
    var newUser user.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if _, exists := users[newUser.Username]; exists {
        http.Error(w, "User already exists", http.StatusConflict)
        return
    }

    hashedPassword, err := user.HashPassword(newUser.Password)
    if err != nil {
        http.Error(w, "Error processing password", http.StatusInternalServerError)
        return
    }

    users[newUser.Username] = hashedPassword

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("User created successfully")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
    var creds user.User
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    storedPassword, ok := users[creds.Username]
    if !ok || !user.CheckPasswordHash(creds.Password, storedPassword) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: creds.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(JwtKey)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
