package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "server internal error")
	InvalidParams             = NewError(10000001, "input params error")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "delegated failed,can't find appKey and appSecret")
	UnauthorizedTokenError    = NewError(10000004, "delegated failed,Token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "delegated failed,Token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "delegated failed,Token generate failed")
	TooManyRequests           = NewError(10000007, "too many requests")
)
