package keeper

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(k Keeper) Migrator {
	return Migrator{keeper: k.(*keeper)}
}

// // Migrate1to2 migrates from version 1 to 2.
// func (m Migrator) Migrate1to2(ctx sdk.Context) error {
// 	return v015.MigrateStore(ctx, m.keeper.skey)
// }
