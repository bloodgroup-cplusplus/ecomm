package storer

type MySQLStorer struct {

	db *sqlx.DB
}

func NewMySQLStorer(db *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{
		db: db,
	}
}


func (ms *MySQLStorer) CreateProduct (ctx context.Context, p *Product)


