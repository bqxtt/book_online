package adapter

import (
	"github.com/bqxtt/book_online/rent/pkg/model"
	"github.com/bqxtt/book_online/rent/pkg/sdk/rentpb"
	"github.com/golang/protobuf/ptypes"
)

func AdaptToPbRentRecord(detail *model.RentRecordDetail) (*rentpb.RentRecord, error) {
	borrowedAtPb, err := ptypes.TimestampProto(detail.BorrowedAt)
	if err != nil {
		return nil, err
	}

	deadlinePb, err := ptypes.TimestampProto(detail.Deadline)
	if err != nil {
		return nil, err
	}

	returnedAtPb, err := ptypes.TimestampProto(detail.ReturnedAt)
	if err != nil {
		return nil, err
	}

	var rentStatus rentpb.RentStatus
	if detail.Finished {
		rentStatus = rentpb.RentStatus_FINISHED
	} else {
		rentStatus = rentpb.RentStatus_IN_PROCESS
	}

	return &rentpb.RentRecord{
		Id:         detail.ID,
		UserId:     detail.UserID,
		Book:       detail.Book,
		Status:     rentStatus,
		BorrowedAt: borrowedAtPb,
		ReturnedAt: returnedAtPb,
		Deadline:   deadlinePb,
	}, nil
}

func AdaptToPbRentRecords(details []*model.RentRecordDetail) ([]*rentpb.RentRecord, error) {
	pbRecords := make([]*rentpb.RentRecord, 0, len(details))

	for _, detail := range details {
		pbRecord, err := AdaptToPbRentRecord(detail)
		if err != nil {
			return nil, err
		}

		pbRecords = append(pbRecords, pbRecord)
	}

	return pbRecords, nil
}
