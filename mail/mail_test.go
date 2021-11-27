package mail

import (
	"testing"

	"github.com/nht1206/pricetracker/internal/model"
)

const (
	TestUserID     = uint64(1)
	TestFullName   = "dummy_full_name"
	TestEmail      = "dummy_email"
	TestGender     = true
	TestFollowType = uint(1)

	TestProductID   = uint64(1)
	TestProductName = "dummy_product_name"
	TestURL         = "dummy_url"
	TestOldPrice    = "15000"
	TestNewPrice    = "10000"
)

func TestBuild(t *testing.T) {
	user := &model.User{
		ID:         TestUserID,
		FullName:   TestFullName,
		Email:      TestEmail,
		Gender:     TestGender,
		FollowType: TestFollowType,
	}
	result := &model.TrackingResult{
		ProductId: TestProductID,
		Name:      TestProductName,
		URL:       TestURL,
		OldPrice:  TestOldPrice,
		NewPrice:  TestNewPrice,
	}
	mailContentBuilder := NewMailContentBuilder()

	_, err := mailContentBuilder.
		SetTemplatePath("../template/mail.tpl").
		SetUser(user).
		SetTrackingResult(result).
		Build()
	if err != nil {
		t.Errorf("failed to build mail content. %v", err)
	}
}

func TestNGBuild(t *testing.T) {
	user := &model.User{
		ID:         TestUserID,
		FullName:   TestFullName,
		Email:      TestEmail,
		Gender:     TestGender,
		FollowType: TestFollowType,
	}
	result := &model.TrackingResult{
		ProductId: TestProductID,
		Name:      TestProductName,
		URL:       TestURL,
		OldPrice:  TestOldPrice,
		NewPrice:  TestNewPrice,
	}
	mailContentBuilder := NewMailContentBuilder()

	t.Run("Case 1: Invalid template path", func(t *testing.T) {
		_, err := mailContentBuilder.
			SetTemplatePath("dummy").
			SetUser(user).
			SetTrackingResult(result).
			Build()
		if err == nil {
			t.Errorf("err is nil")
		}
	})

	t.Run("Case 2: user is nil", func(t *testing.T) {
		_, err := mailContentBuilder.
			SetTemplatePath("dummy").
			SetUser(nil).
			SetTrackingResult(result).
			Build()
		if err == nil {
			t.Errorf("err is nil")
		}
	})

	t.Run("Case 2: tracking result is nil", func(t *testing.T) {
		_, err := mailContentBuilder.
			SetTemplatePath("dummy").
			SetUser(user).
			SetTrackingResult(nil).
			Build()
		if err == nil {
			t.Errorf("err is nil")
		}
	})
}
