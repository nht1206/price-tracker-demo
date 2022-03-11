package static

const (
	ApplicationStatusSuccess = iota
	ApplicationStatusLoggerInitError
	ApplicationStatusContextInitError
)

const (
	ProductStatusDraft = iota + 1
	ProductStatusWaiting
	ProductStatusTracked
	ProductStatusOnTracking
	ProductStatusOnStop
	ProductStatusTrackingFailed = 9
)

const (
	DeleteFlagFalse = 0
	DeleteFlagTrue  = 1
)

const (
	NoTarget           = 0
	MinimumRowAffected = 1
)

const (
	MAIL_TEMPLATE_PATH = "template/mail_%s.tpl"
	DEFAULT_LANGAGE    = "vi"
)

const (
	MaleTitleEn        = "Mr/Mrs"
	FemaleTitleEn      = "Ms/Mrs"
	MaleTitleVi        = "Ông"
	FemaleTitleVi      = "Bà"
	MailHeader         = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	SubjectPlaceholder = "Subject: %s \n%s\n\n"
	SubjectEn          = "PriceTracker notification"
	SubjectVi          = "PriceTracker thông báo"
)
