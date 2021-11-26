package mail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/static"
)

type MailContentBuilder interface {
	Build() (*bytes.Buffer, error)
	SetTemplatePath(path string) MailContentBuilder
	SetUser(user *model.User) MailContentBuilder
	SetTrackingResult(result *model.TrackingResult) MailContentBuilder
}

type mailContentBuilder struct {
	mailTemplatePath string
	user             *model.User
	trackingResult   *model.TrackingResult
}

func NewMailContentBuilder() MailContentBuilder {
	return &mailContentBuilder{}
}

func (b *mailContentBuilder) SetTemplatePath(path string) MailContentBuilder {
	b.mailTemplatePath = path
	return b
}

func (b *mailContentBuilder) SetUser(user *model.User) MailContentBuilder {
	b.user = user
	return b
}

func (b *mailContentBuilder) SetTrackingResult(result *model.TrackingResult) MailContentBuilder {
	b.trackingResult = result
	return b
}

func (b *mailContentBuilder) Build() (*bytes.Buffer, error) {

	if b.user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	if b.trackingResult == nil {
		return nil, fmt.Errorf("tracking download is nil")
	}

	var templatePath string
	if b.mailTemplatePath != "" {
		templatePath = b.mailTemplatePath
	} else {
		templatePath = fmt.Sprintf(static.MAIL_TEMPLATE_PATH, static.DEFAULT_LANGAGE)
	}

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	body.Write([]byte(createSubject(fmt.Sprintf("PriceTracker notification: %s", b.trackingResult.Name))))

	title := ""

	if b.user.Gender {
		title = static.MALE_TITLE
	} else {
		title = static.FEMALE_TITLE
	}

	err = t.Execute(&body, struct {
		Name        string
		Title       string
		ProductName string
		URL         string
		OldPrice    string
		NewPrice    string
	}{
		Name:        b.user.FullName,
		Title:       title,
		ProductName: b.trackingResult.Name,
		URL:         b.trackingResult.URL,
		OldPrice:    b.trackingResult.OldPrice,
		NewPrice:    b.trackingResult.NewPrice,
	})
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func createSubject(subject string) string {
	return fmt.Sprintf(static.SUBJECT_PLACEHOLDER, subject, static.MAIL_HEADER)
}
