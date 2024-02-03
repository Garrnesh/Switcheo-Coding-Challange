package prob

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"prob/testutil/sample"
	probsimulation "prob/x/prob/simulation"
	"prob/x/prob/types"
)

// avoid unused import issue
var (
	_ = probsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateExchange = "op_weight_msg_create_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateExchange int = 100

	opWeightMsgUpdateExchange = "op_weight_msg_update_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateExchange int = 100

	opWeightMsgDeleteExchange = "op_weight_msg_delete_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteExchange int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	probGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&probGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateExchange int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateExchange, &weightMsgCreateExchange, nil,
		func(_ *rand.Rand) {
			weightMsgCreateExchange = defaultWeightMsgCreateExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateExchange,
		probsimulation.SimulateMsgCreateExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateExchange int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateExchange, &weightMsgUpdateExchange, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateExchange = defaultWeightMsgUpdateExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateExchange,
		probsimulation.SimulateMsgUpdateExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteExchange int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteExchange, &weightMsgDeleteExchange, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteExchange = defaultWeightMsgDeleteExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteExchange,
		probsimulation.SimulateMsgDeleteExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateExchange,
			defaultWeightMsgCreateExchange,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				probsimulation.SimulateMsgCreateExchange(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateExchange,
			defaultWeightMsgUpdateExchange,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				probsimulation.SimulateMsgUpdateExchange(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteExchange,
			defaultWeightMsgDeleteExchange,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				probsimulation.SimulateMsgDeleteExchange(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
