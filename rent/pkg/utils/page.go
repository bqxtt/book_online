package utils

import "github.com/bqxtt/book_online/user/pkg/common"

// FIXME: capture and solve integer overflow properly
func CalculateLimitOffset(pageNo int64, pageSize int64) (limit int64, offset int64) {
	if pageNo <= 0 {
		pageNo = 1
	}

	if pageSize <= 0 {
		pageSize = common.DefaultUserPageSize
	}

	return pageSize, (pageNo - 1) * pageSize
}

func IsValidPagedInfo(pageNo int64, pageSize int64) bool {
	return (pageNo > 0) && (pageSize > 0)
}
