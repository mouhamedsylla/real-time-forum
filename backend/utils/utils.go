package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"real-time-forum/orm"
	"reflect"
)

const URL = ""

type User struct {
	message string
}

// The function `RespondWithJSON` writes JSON data to an HTTP response with the specified status code.
func ResponseWithJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// The `OrmInit` function initializes a new ORM instance with a specified database name in Go.
func OrmInit(dbName, dbPath string) (*orm.ORM, error) {
	gorm := orm.NewORM()
	if err := gorm.InitDB(dbName, dbPath); err != nil {
		return nil, err
	}
	return gorm, nil
}

// The CreateDatabase function initializes a database connection using GORM and performs auto migration
// for specified models.
func CreateDatabase(dbName, dbPath string, models ...interface{}) error {
	gorm := orm.NewORM()
	if err := gorm.InitDB(dbName, dbPath); err != nil {
		return err
	}
	gorm.AutoMigrate(models...)
	return nil
}

func InitStorage(dbname, dbpath string, models ...interface{}) (*orm.ORM, error) {
	if _, err := os.Stat(dbpath + dbname); os.IsNotExist(err) {
		CreateDatabase(dbname, dbpath, models...)
	}
	return OrmInit(dbname, dbpath)
}

// The function `DecodeJSONRequestBody` decodes the JSON request body into a specified model and
// returns the decoded data along with an HTTP status code and any potential errors.
func DecodeJSONRequestBody(r *http.Request, model interface{}) (interface{}, int, error) {
	decodedData := reflect.New(reflect.TypeOf(model)).Interface()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = json.Unmarshal(data, decodedData); err != nil {
		return nil, http.StatusBadRequest, err
	}
	return decodedData, http.StatusOK, nil
}
