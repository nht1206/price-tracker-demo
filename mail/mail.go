package mail

import (
	"bytes"
	"errors"
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
		return nil, errors.New("user is nil")
	}

	if b.trackingResult == nil {
		return nil, errors.New("tracking download is nil")
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

	body.Write([]byte(createSubject(b.user.Lang)))

	err = t.Execute(&body, struct {
		Name        string
		Title       string
		ProductName string
		URL         string
		OldPrice    string
		NewPrice    string
	}{
		Name:        b.user.FullName,
		Title:       getUserTitle(b.user.Gender, b.user.Lang),
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

func createSubject(lang string) string {
	switch lang {
	case "en":
		return fmt.Sprintf(static.SubjectPlaceholder, static.SubjectEn, static.MailHeader)
	default:
		return fmt.Sprintf(static.SubjectPlaceholder, static.SubjectVi, static.MailHeader)
	}
}

func getUserTitle(gender bool, lang string) string {
	title := ""

	switch lang {
	case "en":
		if gender {
			title = static.MaleTitleEn
		} else {
			title = static.FemaleTitleEn
		}
		return title
	default:
		if gender {
			title = static.MaleTitleVi
		} else {
			title = static.FemaleTitleVi
		}
		return title
	}
}
