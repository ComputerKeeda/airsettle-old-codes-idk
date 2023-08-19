package simulation

import (
	"math/rand"

	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/keeper"
	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddValidator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddValidator{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddValidator simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddValidator simulation not implemented"), nil, nil
	}
}
