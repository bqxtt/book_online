package handler

import (
	"context"
	"github.com/bqxtt/book_online/rent/pkg/handler/adapter"
	"github.com/bqxtt/book_online/rent/pkg/sdk/base"
	"github.com/bqxtt/book_online/rent/pkg/sdk/rentpb"
	"github.com/bqxtt/book_online/rent/pkg/service"
	"github.com/bqxtt/book_online/rent/pkg/utils"
	"github.com/golang/protobuf/ptypes"
)

const defaultSuccessMessage = "success"

type RentHandler struct {
	rentpb.UnimplementedRentServiceServer

	rentSerivce *service.RentService
}

func NewHandler() (*RentHandler, error) {
	rs, err := service.NewRentService()
	if err != nil {
		return nil, err
	}

	return &RentHandler{
		rentSerivce: rs,
	}, nil
}

func (r *RentHandler) BorrowBook(ctx context.Context, request *rentpb.BorrowBookRequest) (*rentpb.BorrowBookReply, error) {
	if request == nil {
		return &rentpb.BorrowBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil borrow book request pb message"),
		}, nil
	}

	userId, bookId := request.GetUserId(), request.GetBookId()

	deadline, err := r.rentSerivce.BorrowBook(userId, bookId)
	if err != nil {
		return &rentpb.BorrowBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "borrow book failed: %v", err),
		}, nil
	}

	deadlinePb, err := ptypes.TimestampProto(deadline)
	if err != nil {
		return &rentpb.BorrowBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "invalid deadline: %v", err),
		}, nil
	}

	return &rentpb.BorrowBookReply{
		Reply:    utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Deadline: deadlinePb,
	}, nil
}

func (r *RentHandler) ReturnBook(ctx context.Context, request *rentpb.ReturnBookRequest) (*rentpb.ReturnBookReply, error) {
	if request == nil {
		return &rentpb.ReturnBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil return book request pb message"),
		}, nil
	}

	recordId := request.GetRecordId()

	if err := r.rentSerivce.ReturnBook(recordId); err != nil {
		return &rentpb.ReturnBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "return book failed: %v", err),
		}, nil
	}

	return &rentpb.ReturnBookReply{
		Reply: utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
	}, nil
}

func (r *RentHandler) ListBorrowedBook(ctx context.Context, request *rentpb.ListBorrowedBookRequest) (*rentpb.ListBorrowedBookReply, error) {
	if request == nil {
		return &rentpb.ListBorrowedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list borrowed book request pb message"),
		}, nil
	}

	userId := request.GetUserId()
	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()

	details, totalCount, totalPages, err := r.rentSerivce.ListBorrowedBook(userId, pageNo, pageSize)
	if err != nil {
		return &rentpb.ListBorrowedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "list borrowed book records failed: %v", err),
		}, nil
	}

	pbRecords, err := adapter.AdaptToPbRentRecords(details)
	if err != nil {
		return &rentpb.ListBorrowedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt to pb: %v", err),
		}, nil
	}

	return &rentpb.ListBorrowedBookReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Records:    pbRecords,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func (r *RentHandler) ListReturnedBook(ctx context.Context, request *rentpb.ListReturnedBookRequest) (*rentpb.ListReturnedBookReply, error) {
	if request == nil {
		return &rentpb.ListReturnedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list returned book request pb message"),
		}, nil
	}

	userId := request.GetUserId()
	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()

	details, totalCount, totalPages, err := r.rentSerivce.ListReturnedBook(userId, pageNo, pageSize)
	if err != nil {
		return &rentpb.ListReturnedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "list returned book records failed: %v", err),
		}, nil
	}

	pbRecords, err := adapter.AdaptToPbRentRecords(details)
	if err != nil {
		return &rentpb.ListReturnedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt to pb: %v", err),
		}, nil
	}

	return &rentpb.ListReturnedBookReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Records:    pbRecords,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func (r *RentHandler) ListBook(ctx context.Context, request *rentpb.ListBookRequest) (*rentpb.ListBookReply, error) {
	if request == nil {
		return &rentpb.ListBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list book request pb message"),
		}, nil
	}

	userId := request.GetUserId()
	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()

	details, totalCount, totalPages, err := r.rentSerivce.ListRecordByUserId(userId, pageNo, pageSize)
	if err != nil {
		return &rentpb.ListBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "list book records failed: %v", err),
		}, nil
	}

	pbRecords, err := adapter.AdaptToPbRentRecords(details)
	if err != nil {
		return &rentpb.ListBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt to pb: %v", err),
		}, nil
	}

	return &rentpb.ListBookReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Records:    pbRecords,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func (r *RentHandler) ListAllBorrowedBook(ctx context.Context, request *rentpb.ListAllBorrowedBookRequest) (*rentpb.ListAllBorrowedBookReply, error) {
	if request == nil {
		return &rentpb.ListAllBorrowedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list all borrowed book request pb message"),
		}, nil
	}

	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()

	details, totalCount, totalPages, err := r.rentSerivce.ListAllBorrowedBook(pageNo, pageSize)
	if err != nil {
		return &rentpb.ListAllBorrowedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "list all borrowed book records failed: %v", err),
		}, nil
	}

	pbRecords, err := adapter.AdaptToPbRentRecords(details)
	if err != nil {
		return &rentpb.ListAllBorrowedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt to pb: %v", err),
		}, nil
	}

	return &rentpb.ListAllBorrowedBookReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Records:    pbRecords,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func (r *RentHandler) ListAllReturnedBook(ctx context.Context, request *rentpb.ListAllReturnedBookRequest) (*rentpb.ListAllReturnedBookReply, error) {
	if request == nil {
		return &rentpb.ListAllReturnedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list all returned book request pb message"),
		}, nil
	}

	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()

	details, totalCount, totalPages, err := r.rentSerivce.ListAllReturnedBook(pageNo, pageSize)
	if err != nil {
		return &rentpb.ListAllReturnedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "list all returned book records failed: %v", err),
		}, nil
	}

	pbRecords, err := adapter.AdaptToPbRentRecords(details)
	if err != nil {
		return &rentpb.ListAllReturnedBookReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt to pb: %v", err),
		}, nil
	}

	return &rentpb.ListAllReturnedBookReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Records:    pbRecords,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func (r *RentHandler) ListAllBookRecords(ctx context.Context, request *rentpb.ListAllBookRecordsRequest) (*rentpb.ListAllBookRecordsReply, error) {
	if request == nil {
		return &rentpb.ListAllBookRecordsReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list all book records request pb message"),
		}, nil
	}

	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()

	details, totalCount, totalPages, err := r.rentSerivce.ListAllRecord(pageNo, pageSize)
	if err != nil {
		return &rentpb.ListAllBookRecordsReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "list all book records failed: %v", err),
		}, nil
	}

	pbRecords, err := adapter.AdaptToPbRentRecords(details)
	if err != nil {
		return &rentpb.ListAllBookRecordsReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt to pb: %v", err),
		}, nil
	}

	return &rentpb.ListAllBookRecordsReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Records:    pbRecords,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}
