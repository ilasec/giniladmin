package resp

import (
	"fmt"
	"giniladmin/internal/models"
)

// for the fast return success result
func Success() models.CommonResp {
	return models.CommonResp{
		Message: "success",
		Status:  200,
	}
}

// for the fast return failed result
func Failed(message string, args ...interface{}) models.CommonResp {
	return models.CommonResp{
		Message: fmt.Sprintf(message, args...),
		Status:  400,
	}
}

// for the fast return result with custom data
func Data(data interface{}) models.CommonResp {
	return models.CommonResp{
		Message: "success",
		Status:  200,
		Result:  data,
	}
}

// for the fast return success result
func Result(status int, message string, data any) models.CommonResp {
	return models.CommonResp{
		Message: message,
		Status:  int64(status),
		Result:  data,
	}
}
