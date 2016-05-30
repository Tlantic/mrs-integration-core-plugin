package domain

import "gopkg.in/olivere/elastic.v2"

type IIndexer interface {
	Find(index string, query elastic.Query, table string, params ...int) ([]interface{}, int64, error)
	FindById(index string, id string, table string, typeOf interface{}) (interface{}, int64, error)
	Insert(index string, table string, id string, model interface{}) error
	InsertByID(index string, table, id string, model interface{}) error
	Delete(index string, table string, query elastic.Query) error
	DeleteID(index string, table string, id string) error
	Update(index string, table string, id string, data map[string]interface{}) error
	NewBulk() *elastic.BulkService
	AddToBulk(index string, bulk *elastic.BulkService, table string, model interface{}, id string)
	SendBulk(bulk *elastic.BulkService)
}
