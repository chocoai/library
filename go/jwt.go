package main

import (
    "fmt"
    "time"
    "io/ioutil"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
)

var secretKey []byte

func init() {
    if keyData, e := ioutil.ReadFile("SecretKey"); e == nil {
        secretKey = keyData
    } else {
        panic(e)
    }
}

func main() {
    fmt.Printf("%T\n", time.Now().Unix())
    tokenString := CreateToken()
    fmt.Println(tokenString)

    //ValidateToken(tokenString)
}

func CreateToken() string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": 18,
        "iat": time.Now().Unix(),
        "exp": time.Now().Unix() + 60,
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
       fmt.Println(err)
    }

    return tokenString
}

/*
func ValidateToken(tokenString string) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        fmt.Println("user_id:", claims["user_id"])
    } else {
        fmt.Println(err)
    }
}
*/

func ValidateToken(r *http.Request) {
    token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
        func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        })

    if err != nil {
        fmt.Println(err)
    }

    if token.Valid {
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            fmt.Println(err)
        } else {
            fmt.Println("user_id:", claims["user_id"])
        }
    } else {
        fmt.Println("Token is not valid\n")
    }
}
