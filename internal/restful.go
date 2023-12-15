package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseOK represents a successful response with a result field of type gin.H
type ResponseOK struct {
	Result gin.H `json:"result"`
}

// ResponseOKWithList represents a successful response with a result field of type []gin.H
type ResponseOKWithList struct {
	Result []gin.H `json:"result"`
}

// ResponseError represents an error response with a code and message fields.
type ResponseError struct {
	Code int `json:"code,omitempty"`

	Message string `json:"message"`
}

// ResponseSuccess represents a successful response.
func ResponseSuccess(c *gin.Context, result gin.H, resultList []gin.H) {
	if resultList != nil {
		c.IndentedJSON(http.StatusOK, &ResponseOKWithList{
			Result: resultList,
		})
	} else if result != nil {
		c.IndentedJSON(http.StatusOK, &ResponseOK{
			Result: result,
		})
	} else {
		c.IndentedJSON(http.StatusOK, nil)
	}
}

// ResponseFail represents an error response.
func ResponseFail(c *gin.Context, code int, message string) {
	c.IndentedJSON(code, &ResponseError{
		Code:    code,
		Message: message,
	})
}

// NewErrorMessage returns a new error message with the given message and error.
func NewErrorMessage(message string, err error) string {
	return message + ": " + err.Error()
}
