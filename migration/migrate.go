package migration

import (
	"github.com/alvinfebriando/costumer-test/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	user := &entity.User{}

	_ = db.Migrator().DropTable(user)
	_ = db.AutoMigrate(user)
}
