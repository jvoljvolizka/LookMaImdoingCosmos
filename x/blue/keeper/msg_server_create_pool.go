package keeper

import (
	"context"

	"blue/x/blue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var NewPool = types.Pool{
		Owner:       msg.Creator,
		Title:       msg.Title,
		Question:    msg.Question,
		Answerrange: msg.AnswerRange,
	}
	id := k.AppendPool(ctx, NewPool)
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatePoolResponse{Id: id}, nil
}
