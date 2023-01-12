package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuryaContractDetailsValidation(t *testing.T) {
	specs := map[string]struct {
		src    FuryaContractDetails
		expErr bool
	}{
		"all good": {
			src: FuryaContractDetailsFixture(t),
		},
		"empty callbacks": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = nil
			}),
		},
		"multiple callbacks": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{
					{Position: 1, PrivilegeType: "begin_blocker"},
					{Position: 1, PrivilegeType: "end_blocker"},
				}
			}),
		},
		"duplicate callbacks": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{
					{Position: 1, PrivilegeType: "begin_blocker"},
					{Position: 2, PrivilegeType: "begin_blocker"},
				}
			}),
			expErr: true,
		},
		"unknown callback": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{Position: 1, PrivilegeType: "unknown"}}
			}),
			expErr: true,
		},
		"empty callback": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{Position: 1}}
			}),
			expErr: true,
		},
		"invalid callback position": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{Position: math.MaxUint8 + 1, PrivilegeType: "begin_blocker"}}
			}),
			expErr: true,
		},
		"empty callback position": {
			src: FuryaContractDetailsFixture(t, func(d *FuryaContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{PrivilegeType: "begin_blocker"}}
			}),
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			gotErr := spec.src.ValidateBasic()
			if spec.expErr {
				assert.Error(t, gotErr)
				return
			}
			assert.NoError(t, gotErr)
		})
	}
}
