package infra

import (
	"app/interfaces/errs"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"

	sendgrid "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func NewMail(apiKey string) (m *Mail) {
	m = new(Mail)
	m.apiKey = apiKey
	m.apiURL = "https://api.sendgrid.com/v3/mail/send"
	return
}

type Mail struct {
	templateID string
	data       map[string]string
	apiKey     string
	apiURL     string
}

func (m *Mail) SetTemplate(templateID string, data map[string]string) {
	m.templateID = templateID
	m.data = data
}

func (m *Mail) Send(to, subject, msg string) error {

	from := sendgrid.NewEmail("sender name", "noreply@example.com")

	_to, err := mail.ParseAddress(to)

	if err != nil {
		return errs.WrapMsg(err, "mail address can't parsed")
	}

	sm := sendgrid.NewV3Mail()
	sm.SetFrom(from)
	sm.Subject = subject

	p := sendgrid.NewPersonalization()
	p.AddTos(sendgrid.NewEmail(_to.Name, _to.Address))

	if m.templateID != "" {
		for key, value := range m.data {
			p.SetSubstitution(key, value)
		}
		sm.SetTemplateID(m.templateID)
	}

	sm.AddPersonalizations(p)

	c := sendgrid.NewContent("text/html", msg)

	sm.AddContent(c)

	mailbody := sendgrid.GetRequestBody(sm)

	res, err := m.send(mailbody)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 202 {
		errs.NewWithStack("mail can't send, status: %s body: %s", res.Status, string(resBody))
	}

	return nil
}

func (m *Mail) send(data []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", m.apiURL, bytes.NewBuffer(data))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	return res, errs.Wrap(err)
}
