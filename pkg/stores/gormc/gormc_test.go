package gormc

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	Id    int64           `json:"id"`
	Name  string          `json:"name"`
	Attr1 datatypes.JSON  `json:"attr1"`
	Attr2 json.RawMessage `json:"attr2"`
}

func TestNewEngine(t *testing.T) {
	db := NewEngine(Config{
		DSN:   "host=localhost user=postgres password=123456 dbname=zero port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		Debug: true,
	})

	var product = Product{}
	err := db.Where("attr1 ->> 'code' = ?", "100011").First(&product).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(jsonx.MarshalToString(product))
}

func TestJsonQuery(t *testing.T) {
	db := NewEngine(Config{
		DSN:   "host=localhost user=postgres password=123456 dbname=zero port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		Debug: true,
	})

	var products []*Product
	err := db.Scopes(jsonLike("attr1", "code", "100%")).Find(&products).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Println(jsonx.MarshalToString(products))
}

func jsonLike(column, key, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s ->> '%s' like '%s'", column, key, value))
	}
}
