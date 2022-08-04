package simulation

import (
	"math/rand"

	"blue/x/blue/keeper"
	"blue/x/blue/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddQuestion(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddQuestion{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddQuestion simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddQuestion simulation not implemented"), nil, nil
	}
}
