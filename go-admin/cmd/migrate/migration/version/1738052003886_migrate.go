package version

import (
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
	"runtime"

	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1738052003886Test)
}

func _1738052003886Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// TODO: 例如 新增表结构 使用过程中请删除此段代码
		err := tx.Debug().Migrator().AutoMigrate(
			new(models.Supplier),
		)
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
