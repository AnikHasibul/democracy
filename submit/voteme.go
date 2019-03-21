package submit

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/anikhasibul/democracy/parser"
	logger "github.com/anikhasibul/log"
)

var log = logger.New(os.Stdout)

func NewVote(email, password, projectID string) {
	log.Info("Signing up:", email)
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar, Timeout: 3 * time.Minute}
	res, err := client.Get("https://neironix.io/user/auth/sign-up")
	if err != nil {
		log.Error(err)
		return
	}
	signupData := parser.GetForm("#form-signup", res.Body)
	res.Body.Close()
	params := url.Values{}
	for k, v := range signupData {
		if k == "SignupForm[email]" {
			v = email
		}
		if k == "SignupForm[password]" {
			v = password
		}

		params.Add(k, v)
	}
	resp, err := client.PostForm("https://neironix.io/user/auth/sign-up", params)
	if err != nil {
		log.Error(err)
		return
	}
	resp.Body.Close()
	log.Info("Signed up:", email)
	res, err = client.Get("https://neironix.io/user/auth/sign-in")
	if err != nil {
		log.Error(err)
		return
	}
	signinData := parser.GetForm("#login-form", res.Body)
	res.Body.Close()
	params = url.Values{}
	for k, v := range signinData {
		if k == "LoginForm[email]" {
			v = email
		}
		if k == "LoginForm[password]" {
			v = password
		}
		params.Add(k, v)
	}
	resp, err = client.PostForm("https://neironix.io/user/auth/sign-in", params)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Signed in:", email)
	resp.Body.Close()

	res, err = client.Get("https://neironix.io/market-challenge")
	if err != nil {
		log.Error(err)
		return
	}
	csrfToken := parser.GetCSRFToken(res.Body)
	log.Info("Voting as:", email)
	params = url.Values{}
	params.Add("project_id", projectID)
	req, err := http.NewRequest("POST", "https://neironix.io/application/market-challenge/vote", strings.NewReader(params.Encode()))
	if err != nil {
		log.Error(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-CSRF-Token", csrfToken)
	resp, err = client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Read body: ", err)
		return
	}

	log.Info("vote result:", string(data))
	resp.Body.Close()
}
