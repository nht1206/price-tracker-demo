package static

const (
	APPLICATION_STATUS_SUCCESS = iota
	APPLICATION_STATUS_LOGGER_INIT_ERROR
	APPLICATION_STATUS_CONTEXT_INIT_ERROR
)

const (
	PRODUCT_STATUS_DRAFT = iota + 1
	PRODUCT_STATUS_WAITING
	PRODUCT_STATUS_TRACKED
	PRODUCT_STATUS_ON_TRACKING
	PRODUCT_STATUS_ON_STOP
	PRODUCT_STATUS_TRACKING_FAILED = 9
)

const (
	DELETE_FLAG_FALSE = 0
	DELETE_FLAG_TRUE  = 1
)

const (
	NO_TARGET            = 0
	MINIMUM_ROW_AFFECTED = 1
)

const (
	MAIL_TEMPLATE_PATH = "template/mail_%s.tpl"
	DEFAULT_LANGAGE    = "vi"
)

const (
	MALE_TITLE_EN       = "Mr/Mrs"
	FEMALE_TITLE_EN     = "Ms/Mrs"
	MALE_TITLE_VI       = "Ông"
	FEMALE_TITLE_VI     = "Bà"
	MAIL_HEADER         = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	SUBJECT_PLACEHOLDER = "Subject: %s \n%s\n\n"
	SUBJECT_EN          = "PriceTracker notification"
	SUBJECT_VI          = "PriceTracker thông báo"
)
