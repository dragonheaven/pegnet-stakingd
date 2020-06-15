package config

// A list of config locations
const (
	LoggingLevel = "app.loglevel"
	SqliteDBPath = "app.dbpath"
	APIListen    = "app.APIListen"

	// DBlockSync Stuff
	DBlockSyncRetryPeriod = "dblocksync.retry"

	CustomSQLDBMode = "db.mode"
	SQLDBWalMode    = "db.wal"

	Server               = "app.Server"
	Wallet               = "app.Wallet"
	WalletUser           = "app.WalletUser"
	WalletPass           = "app.WalletPass"
	PegnetStakingd       = "app.PegnetStakingd"
	ECPrivateKey         = "app.ECPrivateKey"
	DisableHardForkCheck = "app.DisableHardForkCheck"
)
