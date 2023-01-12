package v2

import (
	"github.com/furyanrasta/furya/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Furya v2 upgrade.
const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
}
