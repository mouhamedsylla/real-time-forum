package utils

import orm "real-time-forum/orm"

type Storage struct {
	Gorm *orm.ORM
}

func NewStorage(dbName, dbPath string) *Storage {
	return &Storage{
		Gorm: OrmInit(dbName, dbPath),
	}
}

func (s *Storage) Insert(model interface{}) error {
	return s.Gorm.Insert(model)
}

func (s *Storage) Select(model interface{}, byField string, value interface{}) interface{} {
	s.Gorm.Custom.Where(byField, value)
	_, table := orm.InitTable(model)
	fields := table.GetFieldName()
	result := s.Gorm.Scan(model, fields...)
	s.Gorm.Custom.Clear()
	return result
}

func (s *Storage) SelectAll(model interface{}) interface{} {
	_, table := orm.InitTable(model)
	fields := table.GetFieldName()
	return s.Gorm.Scan(model, fields...)
}

