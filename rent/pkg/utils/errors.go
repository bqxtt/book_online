package utils

import (
	"errors"
	"fmt"
	"github.com/bqxtt/book_online/rent/pkg/sdk/base"
)

func PbReplyf(status base.REPLY_STATUS, format string, args ...interface{}) *base.BaseReply {
	return &base.BaseReply{
		Status:  status,
		Message: fmt.Sprintf(format, args...),
	}
}

func ParsePbReplyError(reply *base.BaseReply) error {
	if reply.Status != base.REPLY_STATUS_SUCCESS {
		return errors.New(reply.GetMessage())
	}

	return nil
}
