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
	db.Instance.AutoMigrate(&SchemaPkg.Product{})

	// Create
	db.Instance.Create(&SchemaPkg.Product{Code: "D42", Price: 100})

	// Read
	var product SchemaPkg.Product
	db.Instance.First(&product, 1)                 // find product with integer primary key
	db.Instance.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	// db.Instance.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	// db.Instance.Model(&product).Updates(SchemaPkg.Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Instance.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Instance.Delete(&product, 1)
}
