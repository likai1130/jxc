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
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1739328942579Test)
}

func _1739328942579Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		return db.Transaction(func(tx *gorm.DB) error {
			err := tx.Debug().Migrator().AutoMigrate(
				new(models.PurchaseList),
				new(models.PurchaseSaleListGoods),
				new(models.InventoryTransaction),
			)
			if err != nil {
				return err
			}

			return tx.Create(&common.Migration{
				Version: version,
			}).Error
		})

	})
}
