package relay

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"

	cryptocodec "github.com/stratosnet/stratos-chain/crypto/codec"
	pottypes "github.com/stratosnet/stratos-chain/x/pot/types"
	registertypes "github.com/stratosnet/stratos-chain/x/register/types"
	sdstypes "github.com/stratosnet/stratos-chain/x/sds/types"
)

var Cdc *codec.LegacyAmino
var ProtoCdc *codec.ProtoCodec
var Ir codectypes.InterfaceRegistry

func init() {
	Ir = codectypes.NewInterfaceRegistry()
	ProtoCdc = codec.NewProtoCodec(Ir)
	registertypes.RegisterInterfaces(Ir)
	pottypes.RegisterInterfaces(Ir)
	sdstypes.RegisterInterfaces(Ir)
	cryptocodec.RegisterInterfaces(Ir)

	Cdc = codec.NewLegacyAmino()
	sdktypes.RegisterLegacyAminoCodec(Cdc)
	registertypes.RegisterLegacyAminoCodec(Cdc)
	pottypes.RegisterLegacyAminoCodec(Cdc)
	sdstypes.RegisterLegacyAminoCodec(Cdc)
	cryptocodec.RegisterCrypto(Cdc)
	Cdc.Seal()
}
