package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

)

func logWrapper(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
            log.Print(req.RemoteAddr + " " + req.URL.Host + req.URL.Path + "?" + req.URL.RawQuery)
            fmt.Println(req.URL.Path)

			wrappedHandler.ServeHTTP(res, req)
		})
}

func redirectHelper(res http.ResponseWriter, req *http.Request) {
	log.Print("Redirecting...")
	http.Redirect(res, req, "https://nomad-jiujitsu.com" + req.RequestURI, http.StatusMovedPermanently)
}

func redirectToHTTPS() {
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirectHelper)))
}

func forwardToCaravan(msg string, first string, last string, email string) error {
	body := []string{
		"Subject: Form Submission from nomad-jiujitsu.com",
		"To: <caravancollective@outlook.com>",
		"From: kswnin@gmail.com",
		fmt.Sprintf("Reply-To: %s %s <%s>", first, last, email),
		"",
		fmt.Sprintf("Replying to: %s %s <%s>", first, last, email),
		"",
	}

	body = append(body, strings.Split(msg, "\n")...)
	return smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", FROM, GMAIL_KEY, "smtp.gmail.com"),
		FROM,
		[]string{"caravancollective@outlook.com"},
		[]byte(strings.Join(body, "\r\n")),
	)
}
