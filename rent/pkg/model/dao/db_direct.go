package dao

import (
	"github.com/bqxtt/book_online/rent/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBDirectDAO struct {
	DB *gorm.DB
}

func (d *DBDirectDAO) TransactionRent(errChan <-chan error, doneChan chan<- bool, record *model.RentRecord) (*model.RentRecord, error) {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		_, err := d.CreateRentRecord(record)
		if err != nil {
			doneChan <- false
			return err
		}
		doneChan <- true

		select {
		case err := <-errChan:
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (d *DBDirectDAO) TransactionReturn(errChan <-chan error, doneChan chan<- bool, record *model.RentRecord) (*model.RentRecord, error) {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		_, err := d.UpdateRentRecord(record)
		if err != nil {
			doneChan <- false
			return err
		}
		doneChan <- true

		select {
		case err := <-errChan:
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (d *DBDirectDAO) CreateRentRecord(record *model.RentRecord) (*model.RentRecord, error) {
	result := d.DB.Create(record)
	if result.Error != nil {
		return nil, result.Error
	}

	return record, nil
}

func (d *DBDirectDAO) GetRentRecordById(recordId int64) (*model.RentRecord, error) {
	record := &model.RentRecord{}

	result := d.DB.Where("id = ?", recordId).Find(record)
	if result.Error != nil {
		return nil, result.Error
	}

	return record, nil
}

func (d *DBDirectDAO) UpdateRentRecord(record *model.RentRecord) (*model.RentRecord, error) {
	result := d.DB.Model(record).Update(record)
	if result.Error != nil {
		return nil, result.Error
	}

	return record, nil
}

func (d *DBDirectDAO) ListBorrowedRecordByUserId(userId int64, limit int64, offset int64) ([]*model.RentRecord, error) {
	var records []*model.RentRecord

	result := d.DB.Where("user_id = ? and finished = ?", userId, false).Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (d *DBDirectDAO) ListReturnedRecordByUserId(userId int64, limit int64, offset int64) ([]*model.RentRecord, error) {
	var records []*model.RentRecord

	result := d.DB.Where("user_id = ? and finished = ?", userId, true).Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (d *DBDirectDAO) ListRecordByUserId(userId int64, limit int64, offset int64) ([]*model.RentRecord, error) {
	var records []*model.RentRecord

	result := d.DB.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (d *DBDirectDAO) ListAllBorrowedRecord(limit int64, offset int64) ([]*model.RentRecord, error) {
	var records []*model.RentRecord

	result := d.DB.Where("finished = ?", false).Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (d *DBDirectDAO) ListAllReturnedRecord(limit int64, offset int64) ([]*model.RentRecord, error) {
	var records []*model.RentRecord

	result := d.DB.Where("finished = ?", true).Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (d *DBDirectDAO) ListAllRecord(limit int64, offset int64) ([]*model.RentRecord, error) {
	var records []*model.RentRecord

	result := d.DB.Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func NewDBDirectDAO() (*DBDirectDAO, error) {
	db, err := gorm.Open("mysql", "root:19981108@tcp(139.9.140.136:3307)/book_online?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		return nil, err
	}
	return &DBDirectDAO{
		DB: db,
	}, nil
}