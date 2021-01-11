package utils

import (
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/sdk/base"
)

func PbReplyf(status base.REPLY_STATUS, format string, args ...interface{}) *base.BaseReply {
	return &base.BaseReply{
		Status:  status,
		Message: fmt.Sprintf(format, args...),
	}
}
