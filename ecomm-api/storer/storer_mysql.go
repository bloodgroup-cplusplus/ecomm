package storer

type MySQLStorer struct {

	db *sqlx.DB
}

func NewMySQLStorer(db *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{
		db: db,
	}
}


func (ms *MySQLStorer) CreateProduct (ctx context.Context, p *Product) (*Product,error) {

res,err := ns.db.NameExecContext (ctx,"INSERT INTO PRODUCTS (name,image,category,description,rating,num_reviews,price,count_in_stock) VALUES (:name,:image,:category,:description,:rating, :num_reviews, :price,:count_in_sock)")

	if err !=nil {
		return nil, fmt.Error("Error inserting prodcut %w",err)
	}

	id , err := res.LastInsertId()
	if err !=nil {
		return nil, fmt.Errorf("Error getting last insert id: %w",err)
	}


	p.ID = id 

	return p, nil 
}


func (ms *MySQLStorer) GetProduct (ctx context.Context,id int64) (*Product,error){
	var p Product 
	err := ms.db.GetContext(ctx,&p,"SELECT * FROM products where id=?",id)
	if err !=nil {
		return nil, fmt.Errorf("Error getting product: %w",err)
	}

	return &p,nil 
}

func (ms *MYSQLStorer) ListProducts (ctx context.Context) ([] *Product,error) {
	var products []*Product 
	err := ms.db.SelectContext(ctx, &products, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error listing products:%w", err)
	}

	return products, nil 
}


