package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

var PATH string

// The `InitDB` function is responsible for initializing the database connection and creating the
// necessary files and directories for database migration.
func (o *ORM) InitDB(name string, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0775)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	_, err := os.Stat(path + name)
	if os.IsNotExist(err) {
		file, err := os.Create(path + name)
		if err != nil {
			return fmt.Errorf("failed to create database file: %v", err)
		}
		file.Close()
	}

	if _, err := os.Stat(path + "migrates"); os.IsNotExist(err) {
		err := os.Mkdir(path+"migrates", 0755)
		if err != nil {
			return fmt.Errorf("failed to create migrates directory: %v", err)
		}
	}

	o.Db, err = sql.Open("sqlite3", path+name)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	PATH = path
	return nil
}

// The CreateTable function creates a SQL table with the given name and fields.
func CreateTable(name string, fields ...*Field) string {
	sqlTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", name)
	var all []string
	for _, field := range fields {
		all = append(all, "\t"+TableField(field))
	}
	sqlTable += strings.Join(all, ",\n") + "\n)"
	return sqlTable
}

// The `AutoMigrate` function is responsible for automatically creating database tables based on the
// provided struct definitions. It takes in a variadic parameter `tables` which represents the struct
// definitions of the tables to be created.
func (o *ORM) AutoMigrate(tables ...interface{}) error {
	for _, table := range tables {
		v, _table := InitTable(table)

		createTableSQL := CreateTable(v.Name(), _table.AllFields...)
		if len(_table.ForeignKey) > 0 {
			createTableSQL = strings.TrimSuffix(createTableSQL, "\n)")
			createTableSQL += ",\n" + "\t" + strings.Join(_table.ForeignKey, ",\n") + "\n)"
		}

		o.AddTable(_table)
		_, err := o.Db.Exec(createTableSQL)
		if err != nil {
			return fmt.Errorf("failed to execute create table SQL for table %s: %v", v.Name(), err)
		}

		currentTime := time.Now()
		fileName := fmt.Sprintf("%smigrates/%s-create-table-%s.sql", PATH, currentTime.Format("2006-01-02-15-04-05"), v.Name())
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create migration file %s: %v", fileName, err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("failed to close migration file %s: %v", fileName, err)
			}
		}()

		_, err = file.WriteString(createTableSQL)
		if err != nil {
			return fmt.Errorf("failed to write to migration file %s: %v", fileName, err)
		}
	}
	return nil
}

// The function `InitTable` initializes a table by extracting field information from a given struct
// type and creating corresponding fields in the table.
func InitTable(table interface{}) (reflect.Type, *Table) {
	v := reflect.TypeOf(table)
	_table := NewTable(v.Name())
	var foreignKeys []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Type

		if fieldType.Kind() == reflect.Struct {
			// Gérer les sous-structs récursivement.
			for j := 0; j < fieldType.NumField(); j++ {
				structField := fieldType.Field(j)
				ormgoTag, fk := GetTags(structField)
				foreignKeys = append(foreignKeys, fk...)
				_table.AddField(NewField(structField.Name, structField.Type, ormgoTag))
			}
		} else {
			// Gérer les champs normaux.
			ormgoTag, fk := GetTags(field)
			foreignKeys = append(foreignKeys, fk...)
			_table.AddField(NewField(field.Name, fieldType, ormgoTag))
		}
		_table.ForeignKey = foreignKeys
	}

	return v, _table
}
