package resend

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/resend/resend-go/v2"
)

//go:embed code.html
var CodeHTML string
var (
	Client resend.Client
)

func init() {
	apiKey := "re_cHSSXzRF_BUGw3ReXycy3abECzgA2GbS7"
	Client = *resend.NewClient(apiKey)
}

type ResendService struct{}

func (r *ResendService) Send(code string, email ...string) (*resend.SendEmailResponse, error) {

	html, err := r.Html(map[string]string{"Code": code})
	if err != nil {
		return nil, err
	}

	params := &resend.SendEmailRequest{
		From:    "cmc <cursor@cursor.jshub.top>",
		Subject: "cmc service email vaildate code",
		Html:    html,
		Bcc:     email,
		To:      []string{"a@example.com"},
	}

	sent, err := Client.Emails.Send(params)
	if err != nil {
		return nil, err
	}

	return sent, nil

}

func (r *ResendService) Html(data map[string]string) (string, error) {

	tmpl, err := template.New("code").Parse(CodeHTML)

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil

}
