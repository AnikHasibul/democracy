package parser

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestSignIn(t *testing.T) {
	res, err := http.Get("https://neironix.io/user/auth/sign-in")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(GetForm("#login-form", res.Body))
}

func TestSignUp(t *testing.T) {
	res, err := http.Get("https://neironix.io/user/auth/sign-up")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(GetForm("#form-signup", res.Body))
}

func TestCSRF(t *testing.T) {
	res, err := http.Get("https://neironix.io/user/auth/sign-up")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(GetCSRFToken(res.Body))
}
