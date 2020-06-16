module github.com/dragonheaven/pegnet-stakingd

go 1.13

require (
	github.com/Factom-Asset-Tokens/factom v0.0.0-20191121210333-481aa2c19345
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/pegnet/pegnet v0.5.0
	github.com/pegnet/pegnetd v0.5.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/spf13/viper v1.7.0
)

replace github.com/Factom-Asset-Tokens/factom => github.com/Emyrk/factom v0.0.0-20200113153851-17d98c31e1bd
