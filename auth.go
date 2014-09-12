package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func clientLoginClient(host, email, password string) *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("failed to make cookie jar: %v", err)
	}
	client := &http.Client{
		Jar: jar,
	}

	v := url.Values{}
	v.Set("Email", email)
	v.Set("Passwd", password)
	v.Set("service", "ah")
	v.Set("source", "Misc-remote_api-0.1")
	v.Set("accountType", "HOSTED_OR_GOOGLE")

	resp, err := client.PostForm("https://www.google.com/accounts/ClientLogin", v)
	if err != nil {
		log.Fatalf("could not post login: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unsuccessful request: status %d; body %q", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatalf("unable to read response: %v", err)
	}

	m := regexp.MustCompile(`Auth=(\S+)`).FindSubmatch(body)
	if m == nil {
		log.Fatalf("no auth code in response %q", body)
	}
	auth := string(m[1])

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     "/_ah/login",
		RawQuery: "continue=/&auth=" + url.QueryEscape(auth),
	}

	// Disallow redirects.
	redirectErr := errors.New("stopping redirect")
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return redirectErr
	}

	resp, err = client.Get(u.String())
	if urlErr, ok := err.(*url.Error); !ok || urlErr.Err != redirectErr {
		log.Fatalf("could not get auth cookies: %v", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusFound {
		log.Fatalf("unsuccessful request: status %d; body %q", resp.StatusCode, body)
	}

	client.CheckRedirect = nil
	return client
}
