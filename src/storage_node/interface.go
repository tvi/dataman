package storagenode

import (
	"github.com/jacksontj/dataman/src/metadata"
	"github.com/jacksontj/dataman/src/query"
)

// Interface that a storage node must implement
type StorageInterface interface {
	// Initialization, this is the "config_json" for the `storage_node`
	Init(map[string]interface{}) error

	// Get the current meta from however it is stored
	GetMeta() (*metadata.Meta, error)

	// Schema-Functions
	AddDatabase(db *metadata.Database) error
	RemoveDatabase(dbname string) error

	AddTable(dbname string, table *metadata.Table) error
	RemoveTable(dbname string, tablename string) error

	AddIndex(dbname, tablename string, index *metadata.TableIndex) error
	RemoveIndex(dbname, tablename, indexname string) error

	// Data-Functions
	// TODO: split out the various functions into grouped interfaces that make sense
	// for now we'll just have one, but eventually we could support "TransactionalStorageNode" etc.
	// TODO: more specific types for each method
	Get(query.QueryArgs) *query.Result
	Set(query.QueryArgs) *query.Result
	//Delete(query.QueryArgs) *query.Result
	Filter(query.QueryArgs) *query.Result
}
