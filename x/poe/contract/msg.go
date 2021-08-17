package contract

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

// TG4GroupInitMsg contract init message
//See https://github.com/confio/tgrade-contracts/blob/main/contracts/tg4-group/schema/instantiate_msg.json
type TG4GroupInitMsg struct {
	Admin    string      `json:"admin,omitempty"`
	Members  []TG4Member `json:"members"`
	Preauths uint64      `json:"preauths,omitempty"`
}

func (m TG4GroupInitMsg) Json(t *testing.T) string {
	return asJson(t, m)
}

type TG4Member struct {
	Addr   string `json:"addr"`
	Weight uint64 `json:"weight"`
}

func SortByWeightDesc(s []TG4Member) []TG4Member {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Weight > s[j].Weight || s[i].Weight == s[j].Weight && s[i].Addr < s[j].Addr
	})
	return s
}

// TG4UpdateMembersMsg contract execute message to update members
// See https://github.com/CosmWasm/cosmwasm-plus/tree/main/contracts/cw4-group
// https://github.com/confio/tgrade-contracts/blob/main/contracts/tg4-group/schema/execute_msg.json
type TG4UpdateMembersMsg struct {
	Add    []TG4Member `json:"add"`
	Remove []string    `json:"remove"`
}

func (m *TG4UpdateMembersMsg) Json(t *testing.T) string {
	switch {
	case m.Add == nil:
		m.Add = make([]TG4Member, 0)
	case m.Remove == nil:
		m.Remove = make([]string, 0)
	}
	x := map[string]interface{}{
		"update_members": m,
	}
	return asJson(t, x)
}

// TG4MixerInitMsg contract init message
//See https://github.com/confio/tgrade-contracts/blob/main/contracts/tg4-mixer/schema/instantiate_msg.json
type TG4MixerInitMsg struct {
	//Admin      string `json:"admin,omitempty"`
	LeftGroup  string `json:"left_group"`
	RightGroup string `json:"right_group"`
	Preauths   uint64 `json:"preauths,omitempty"`
}

func (m TG4MixerInitMsg) Json(t *testing.T) string {
	return asJson(t, m)
}

type TG4StakeInitMsg struct {
	Admin           string         `json:"admin,omitempty"`
	Denom           Denom          `json:"denom"`
	MinBond         string         `json:"min_bond"`
	TokensPerWeight string         `json:"tokens_per_weight"`
	UnbondingPeriod UnbodingPeriod `json:"unbonding_period"`
	Preauths        uint64         `json:"preauths,omitempty"`
}

func (m TG4StakeInitMsg) Json(t *testing.T) string {
	return asJson(t, m)
}

type Denom struct {
	Native string `json:"native,omitempty"`
	CW20   string `json:"cw20,omitempty"`
}

type UnbodingPeriod struct {
	Height    uint64 `json:"height,omitempty"`
	TimeInSec uint64 `json:"time,omitempty"`
}

type TG4StakeExecute struct {
	Bond *struct{} `json:"bond,omitempty"`
}

func (m TG4StakeExecute) Json(t *testing.T) string {
	return asJson(t, m)
}

// ValsetInitMsg Valset contract init message
// See https://github.com/confio/tgrade-contracts/tree/main/contracts/tgrade-valset
type ValsetInitMsg struct {
	Membership    string      `json:"membership"`
	MinWeight     int         `json:"min_weight"`
	MaxValidators int         `json:"max_validators"`
	EpochLength   int         `json:"epoch_length"`
	EpochReward   sdk.Coin    `json:"epoch_reward"`
	InitialKeys   []Validator `json:"initial_keys"`
	Scaling       int         `json:"scaling,omitempty"`
}

func (m ValsetInitMsg) Json(t *testing.T) string {
	return asJson(t, m)
}

type Validator struct {
	Operator        string          `json:"operator"`
	ValidatorPubkey ValidatorPubkey `json:"validator_pubkey"`
}

func asJson(t *testing.T, m interface{}) string {
	t.Helper()
	r, err := json.Marshal(&m)
	require.NoError(t, err)
	return string(r)
}

// TG4ValsetExecute Valset contract validator key registration
// See https://github.com/confio/tgrade-contracts/blob/main/contracts/tgrade-valset/schema/execute_msg.json
type TG4ValsetExecute struct {
	RegisterValidatorKey *RegisterValidatorKey `json:"register_validator_key,omitempty"`
	UpdateMetadata       *ValidatorMetadata    `json:"update_metadata,omitempty"`
}

type RegisterValidatorKey struct {
	PubKey   ValidatorPubkey   `json:"pubkey"`
	Metadata ValidatorMetadata `json:"metadata"`
}

type ValidatorMetadata struct {
	// moniker defines a human-readable name for the validator.
	Moniker string `json:"moniker"`
	// identity defines an optional identity signature (ex. UPort or Keybase).
	Identity string `json:"identity,omitempty"`
	// website defines an optional website link.
	Website string `json:"website,omitempty"`
	// security_contact defines an optional email for security contact.
	SecurityContact string `json:"security_contact,omitempty"`
	// details define other optional details.
	Details string `json:"details,omitempty"`
}

func MetadataFromDescription(description stakingtypes.Description) ValidatorMetadata {
	return ValidatorMetadata{
		Moniker:         description.Moniker,
		Identity:        description.Identity,
		Website:         description.Website,
		SecurityContact: description.SecurityContact,
		Details:         description.Details,
	}
}

func (m ValidatorMetadata) ToDescription() stakingtypes.Description {
	return stakingtypes.Description{
		Moniker:         m.Moniker,
		Identity:        m.Identity,
		Website:         m.Website,
		SecurityContact: m.SecurityContact,
		Details:         m.Details,
	}

}