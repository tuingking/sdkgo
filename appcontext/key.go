package appcontext

type key string

const (
	RequestID     key = `x-request-id`
	AppLang       key = `x-app-lang`
	AppName       key = "x-app-name"
	AppDebug      key = "x-app-debug"
	UserID        key = `x-user-id`
	RequestMethod key = `x-request-method`
	PackageName   key = `x-package-name`
)
