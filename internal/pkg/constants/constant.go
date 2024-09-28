package constants

const (
	EmailRegex  = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	MobileRegex = "^([+]\\d{2})?\\d{10}$"
)

var (
	PostRequestColumns   = []string{"name", "email", "mobile", "created_at", "updated_at"}
	ListRequestColumns   = []string{"id", "name", "email", "mobile", "created_at", "updated_at"}
	UpdateRequestColumns = []string{"name", "email", "mobile", "updated_at"}
)
