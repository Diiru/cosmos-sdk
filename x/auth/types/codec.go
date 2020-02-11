package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

var _ AuthCodec = (*Codec)(nil)

// AuthCodec defines the interface needed to serialize x/auth state. It must be
// aware of all concrete account types.
type AuthCodec interface {
	codec.Marshaler

	MarshalAccount(acc exported.AccountI) ([]byte, error)
	UnmarshalAccount(bz []byte) (exported.AccountI, error)

	MarshalAccountJSON(acc exported.AccountI) ([]byte, error)
	UnmarshalAccountJSON(bz []byte) (exported.AccountI, error)
}

type Codec struct {
	codec.Marshaler

	// Keep reference to the amino codec to allow backwards compatibility along
	// with type, and interface registration.
	amino *codec.Codec
}

func NewCodec(amino *codec.Codec) *Codec {
	return &Codec{Marshaler: codec.NewHybridCodec(amino), amino: amino}
}

// MarshalAccount marshals an AccountI interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (c *Codec) MarshalAccount(accI exported.AccountI) ([]byte, error) {
	acc := &Account{}
	acc.SetAccountI(accI)
	return c.Marshaler.MarshalBinaryLengthPrefixed(acc)
}

// UnmarshalAccount returns an AccountI interface from raw encoded account bytes
// of a Proto-based Account type. An error is returned upon decoding failure.
func (c *Codec) UnmarshalAccount(bz []byte) (exported.AccountI, error) {
	acc := &Account{}
	if err := c.Marshaler.UnmarshalBinaryLengthPrefixed(bz, acc); err != nil {
		return nil, err
	}
	return acc.GetAccountI(), nil
}

// MarshalAccountJSON JSON encodes an account object implementing the AccountI
// interface.
func (c *Codec) MarshalAccountJSON(acc exported.AccountI) ([]byte, error) {
	return c.Marshaler.MarshalJSON(acc)
}

// UnmarshalAccountJSON returns an AccountI from JSON encoded bytes.
func (c *Codec) UnmarshalAccountJSON(bz []byte) (exported.AccountI, error) {
	acc := &Account{}
	if err := c.Marshaler.UnmarshalJSON(bz, acc); err != nil {
		return nil, err
	}

	return acc.GetAccountI(), nil
}

// ----------------------------------------------------------------------------

// ModuleCdc is the global x/auth Amino codec.
//
// NOTE: This codec is deprecated, where a codec via NewCodec without an Amino
// codec should be used.
var ModuleCdc = NewCodec(codec.New())

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.GenesisAccount)(nil), nil)
	cdc.RegisterInterface((*exported.AccountI)(nil), nil)
	cdc.RegisterConcrete(&BaseAccount{}, "cosmos-sdk/Account", nil)
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}

// RegisterAccountTypeCodec registers an external account type defined in
// another module for the internal ModuleCdc.
func RegisterAccountTypeCodec(o interface{}, name string) {
	ModuleCdc.amino.RegisterConcrete(o, name, nil)
}

func init() {
	RegisterCodec(ModuleCdc.amino)
	codec.RegisterCrypto(ModuleCdc.amino)
}
