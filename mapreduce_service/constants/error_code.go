package constants

import "errors"

const (
	CALL_ERPLY_API_ERROR     = -101
	NORMAL_RESPONSE_STATUS   = 0
	ERROR_RESPONSE_STATUS    = -1
	ERROR_STATUS_PARAM_WRONG = -2
	TOKEN_EXPIRE             = -200
	HEADER_AUTH_EMPTY        = -300
	DATA_BASE_CALL_ERROR     = -400
)

var (
	MoveMysqlEmptyError = errors.New("remove not records")
)
