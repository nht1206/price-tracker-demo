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