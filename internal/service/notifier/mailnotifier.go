package notifier

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/mail"
)

type mailNotifier struct {
	config *config.MailConfig
}

func newMailNotifier(cfg *config.MailConfig) (Notifier, error) {
	if cfg == nil {
		return nil, errors.New("mail cfg is nil")
	}

	return &mailNotifier{
		config: cfg,
	}, nil
}

func (n *mailNotifier) Notify(user *model.User, result *model.TrackingResult) error {
	mailContentBuilder := mail.NewMailContentBuilder()

	mailAuth := smtp.PlainAuth("", n.config.Sender, n.config.SenderPassword, n.config.SMTPHost)

	mailBody, err := mailContentBuilder.
		SetUser(user).
		SetTrackingResult(result).
		Build()
	if err != nil {
		return err
	}

	sendTo := []string{user.Email}

	return smtp.SendMail(n.getMailAddr(), mailAuth, n.config.Sender, sendTo, mailBody.Bytes())
}

func (n *mailNotifier) getMailAddr() string {
	return fmt.Sprintf("%s:%s", n.config.SMTPHost, n.config.SMTPPort)
}
