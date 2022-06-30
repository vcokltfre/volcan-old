package impl

import "github.com/sirupsen/logrus"

var Error *ErrorHandler

type ErrorHandler struct{}

func (e *ErrorHandler) Dispatch(err error) {
	logrus.Error(err)
}
