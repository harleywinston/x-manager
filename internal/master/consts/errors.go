package consts

import (
	"fmt"
)

type CustomError struct {
	Message string
	Code    int
	Detail  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf(`%d: %s\n%v`, e.Code, e.Message, e.Detail)
}

var (
	INVALID_IP_ERROR            = &CustomError{Message: "Invalid Ip Address!", Code: 500}
	INVALID_DOMAIN_ERROR        = &CustomError{Message: "Invalid Domain Name!", Code: 500}
	INVALID_GROUP_MODE_ERROR    = &CustomError{Message: "Invalid group mode!", Code: 500}
	INVALID_EMAIL_ERROR         = &CustomError{Message: "Invalid email address!", Code: 500}
	GROUP_LIMIT_ERROR           = &CustomError{Message: "All groups limit exceeded!", Code: 500}
	RECOURSE_ID_NOT_VALID_ERROR = &CustomError{Message: "Resource id is not valid!", Code: 500}

	ADD_DB_ERROR          = &CustomError{Message: "Add to DB had error!", Code: 500}
	GET_DB_ERROR          = &CustomError{Message: "Get from DB had error!", Code: 500}
	DELETE_DB_ERROR       = &CustomError{Message: "Delete from DB had error!", Code: 500}
	AUTO_MIGRATE_DB_ERROR = &CustomError{Message: "AutoMigrate to DB had error!", Code: 500}

	PARSE_INT_ERROR = &CustomError{
		Message: "Internal strconv  parseINT error!",
		Code:    500,
	}

	BIND_JSON_ERROR = &CustomError{Message: "Bind json failed!", Code: 500}

	ADD_SUCCESS    = &CustomError{Message: "Add succeed.", Code: 200}
	DELETE_SUCCESS = &CustomError{Message: "Delete succeed.", Code: 200}

	INVALID_RESOURCE_IP_ERROR = &CustomError{Message: "Invalid resource ip! Fuck you :)", Code: 502}
	INVALID_GROUP_ID_ERROR    = &CustomError{Message: "Invalid group id!", Code: 500}
)
