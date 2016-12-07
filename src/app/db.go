package app

type DBWhere map[string]interface{}

type DBFilter struct {
	Limit   int
	Offset  int
	OrderBy string
	Reverse bool
	Preload []string
}

type DBCreator interface {
	Store(interface{}) error
}

type DBFinder interface {
	One(model interface{}, id int) error
	OneBy(model interface{}, w DBWhere) error
	FindBy(models interface{}, w DBWhere, f *DBFilter) error
	FirstOrInit(m interface{}, w DBWhere) error
}

type DBExistser interface {
	ExistsBy(b interface{}, w DBWhere) (bool, error)
}

type DBCreatorExistser interface {
	DBCreator
	DBExistser
}

// IDatabase is database interface
type IDatabase interface {
	DBCreator
	DBExistser
	DBFinder
	// Count(model interface{}, query Query) (uint, error)
	UpdateField(model interface{}, field string, value interface{}) error
	// Update(model interface{}, query Query, change Query) (affectedRows uint, err error)
	// Delete(model interface{}, query Query) (affectedRows uint, err error)
	IsNotFoundErr(error) bool
}
