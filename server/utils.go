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

func forwardToCaravan(msg string) error {
	body := []string{
		"Subject: Form Submission from nomad-jiujitsu.com",
		"To: caravancollective@outlook.com",
		"From: kswnin@gmail.com",
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

/*type MailGunInput struct {
    To string `url:"to,omitempty"`
    From string `url:"from,omitempty"`
    ReplyTo string `url:"h:Reply-To,omitempty"`
    Subject string `url:"subject,omitempty"`
    Body string `url:"text,omitempty"`
}

func mailgun(in MailGunInput) error {
    v, err := query.Values(in)
    if err != nil {
        return err
    }

    client := &http.Client{}
    req, _ := http.NewRequest("POST",
        "https://api.mailgun.net/v3/mg.nomad-jiujitsu.com/messages",
        strings.NewReader(v.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.SetBasicAuth("api", MAILGUN_KEY)

    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    if resp.StatusCode > 399 {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }

        return errors.New(string(body))
    }
    return nil
}*/
