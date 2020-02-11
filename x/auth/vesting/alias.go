package vesting

// nolint

import (
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

var (
	// functions aliases
	RegisterCodec                  = types.RegisterCodec
	NewBaseVestingAccount          = types.NewBaseVestingAccount
	NewContinuousVestingAccountRaw = types.NewContinuousVestingAccountRaw
	NewContinuousVestingAccount    = types.NewContinuousVestingAccount
	NewPeriodicVestingAccountRaw   = types.NewPeriodicVestingAccountRaw
	NewPeriodicVestingAccount      = types.NewPeriodicVestingAccount
	NewDelayedVestingAccountRaw    = types.NewDelayedVestingAccountRaw
	NewDelayedVestingAccount       = types.NewDelayedVestingAccount
	NewCodec                       = types.NewCodec

	// variable aliases
	VestingCdc = types.VestingCdc
)

type (
	BaseVestingAccount       = types.BaseVestingAccount
	ContinuousVestingAccount = types.ContinuousVestingAccount
	PeriodicVestingAccount   = types.PeriodicVestingAccount
	DelayedVestingAccount    = types.DelayedVestingAccount
	Period                   = types.Period
	Periods                  = types.Periods
	Codec                    = types.Codec
)
