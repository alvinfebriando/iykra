package migration

import (
	"github.com/alvinfebriando/costumer-test/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	customer := &entity.Customer{}

	_ = db.Migrator().DropTable(customer)
	_ = db.AutoMigrate(customer)
}
