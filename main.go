package main

import (
	ConfigPkg "demo/src/config"
	DBPkg "demo/src/db"
	SchemaPkg "demo/src/schema"

	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootPath   = filepath.Dir(b)
)

func main() {
	config := ConfigPkg.Config{}
	config.Init(rootPath)

	db := DBPkg.DB{}
	db.Init(&config.DB)

	// Migrate the schema
	// Create table `rename_product` with struct Product
	db.Instance.Table("rename_product").AutoMigrate(&SchemaPkg.Product{})

	// Create
	db.Instance.Create(&SchemaPkg.Product{Code: "D43", Price: 100})

	// Read
	var product SchemaPkg.Product
	db.Instance.First(&product) // find product with integer primary key
	// log.Printf("demo: %+v\n", product)
	// db.Instance.First(&product, "code = ?", "D43") // find product with code D42
	// log.Printf("demo: %+v\n", product)

	product.Price = 400
	// Update - update product's price to 200
	db.Instance.Model(&product).Updates(&product)
	// Update - update multiple fields
	// db.Instance.Model(&product).Updates(SchemaPkg.Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Instance.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Instance.Delete(&product, "code = ?", product.Code)

	db.Instance.Exec("DROP TABLE products")
}
