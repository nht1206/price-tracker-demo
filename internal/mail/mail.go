package mail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/static"
)

type MailContentBuilder interface {
	Build() (bytes.Buffer, error)
}

type mailContentBuilder struct {
	user   *model.User
	result *model.TrackingResult
}

func NewMailContentBuilder(user *model.User, result *model.TrackingResult) MailContentBuilder {
	return &mailContentBuilder{
		user:   user,
		result: result,
	}
}

func (b *mailContentBuilder) Build() (bytes.Buffer, error) {
	var body bytes.Buffer

	t, err := template.ParseFiles(static.MAIL_TEMPLATE_PATH)
	if err != nil {
		return body, err
	}

	body.Write([]byte(createSubject(fmt.Sprintf("PriceTracker notification: %s", b.result.Name))))

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
		ProductName: b.result.Name,
		URL:         b.result.URL,
		OldPrice:    b.result.OldPrice,
		NewPrice:    b.result.NewPrice,
	})
	if err != nil {
		return body, err
	}

	return body, nil
}

func createSubject(subject string) string {
	return fmt.Sprintf(static.SUBJECT_PLACEHOLDER, subject, static.MAIL_HEADER)
}
