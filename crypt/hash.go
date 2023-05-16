package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var secretKey = readKey()

func readKey() []byte {
	data, err := os.ReadFile("C:/Users/1/Proetkirovanie/crypt/key.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	key, err := hex.DecodeString(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println(key)
	return key
}

func SecretKey() []byte {
	return secretKey
}

func Hash(str string) string {
	encoder := hmac.New(sha256.New, secretKey)
	encoder.Write([]byte(str))
	return hex.EncodeToString(encoder.Sum(nil))
}

type Header struct {
	Alg string
	Typ string
}
type Payload struct {
	Iss string
	Exp int64
	Jti int
}

func SetCookies(w http.ResponseWriter, id int) {
	header, _ := json.Marshal(Header{"HS256", "JWT"})
	payload, _ := json.Marshal(Payload{"UIRS", (time.Now().Unix() + 1200), id})
	unsigned := base64.StdEncoding.EncodeToString(header) + "." + base64.StdEncoding.EncodeToString(payload)
	sign := base64.StdEncoding.EncodeToString([]byte(Hash(unsigned)))
	token := unsigned + "." + sign
	tokenCookie := http.Cookie{
		Name:     "Token",
		Value:    token,
		MaxAge:   3000,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, &tokenCookie)
	header, _ = json.Marshal(Header{"HS256", "JWT"})
	payload, _ = json.Marshal(Payload{"UIRS", (time.Now().Unix() + 2592000), id})
	unsigned = base64.StdEncoding.EncodeToString(header) + "." + base64.StdEncoding.EncodeToString(payload)
	sign = base64.StdEncoding.EncodeToString([]byte(Hash(unsigned)))
	token = unsigned + "." + sign
	refreshCookie := http.Cookie{
		Name:     "Refresh",
		Value:    token,
		HttpOnly: true,
		MaxAge:   2592000,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, &refreshCookie)
}

func CookieIsValid(cookies []*http.Cookie, name string) int {
	for _, cookie := range cookies {
		//fmt.Println(cookie.Name)
		if cookie.Name == name { //tbh you need to check header for decoding algorithm header, _ := base64.StdEncoding.DecodeString(parts[0])
			//fmt.Println("Cookie found")
			token := cookie.Value
			//fmt.Printf("Token value: %s\n", token)
			parts := strings.Split(token, ".")
			fmt.Println(len(parts))
			if len(parts) != 3 {
				return 0
			}
			//fmt.Printf("Parts %v", parts)
			sign := base64.StdEncoding.EncodeToString([]byte(Hash(parts[0] + "." + parts[1])))
			//fmt.Println("What server produced", sign)
			//fmt.Println("What we got", parts[2])
			if sign != parts[2] {
				return 0
			}
			json_str, _ := base64.StdEncoding.DecodeString(parts[1])
			var payload Payload
			json.Unmarshal([]byte(json_str), &payload)
			if payload.Exp < time.Now().Unix() {
				return 0
			}
			return payload.Jti
		}
	}
	fmt.Println("Cookie not found")
	return 0
}
