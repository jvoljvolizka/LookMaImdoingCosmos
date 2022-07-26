package blue

import (
	"math/rand"

	"blue/testutil/sample"
	bluesimulation "blue/x/blue/simulation"
	"blue/x/blue/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bluesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePool = "op_weight_msg_create_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	opWeightMsgVote = "op_weight_msg_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVote int = 100

	opWeightMsgCreatePrefPool = "op_weight_msg_create_pref_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePrefPool int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blueGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blueGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePool, &weightMsgCreatePool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePool = defaultWeightMsgCreatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePool,
		bluesimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = defaultWeightMsgVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVote,
		bluesimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePrefPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePrefPool, &weightMsgCreatePrefPool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePrefPool = defaultWeightMsgCreatePrefPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePrefPool,
		bluesimulation.SimulateMsgCreatePrefPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
