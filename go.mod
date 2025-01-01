module galycherrygame

go 1.22

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
	github.com/mattn/go-sqlite3 v1.14.24
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.7
)

replace gorm.io/driver/sqlite => github.com/go-gorm/sqlite v1.5.5
