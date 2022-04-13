package models

// The product struct
type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func searchAllProducts() []Product {
	db := DbConnection()

	allProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	// With the function Next, all the columns and lines from the query will be checked
	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		// Scan copies the columns in the current row into the values pointed at by dest.
		// The number of values in dest must be the same as the number of columns in Rows.
		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	// CLosing the database
	defer db.Close()

	return products
}
