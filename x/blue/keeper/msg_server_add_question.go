package keeper

import (
	"context"

	"blue/x/blue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddQuestion(goCtx context.Context, msg *types.MsgAddQuestion) (*types.MsgAddQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, ok := k.GetPrefPool(ctx, msg.PoolId)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if msg.Creator != pool.Creator {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	var newQuestion = types.Question{}
	newQuestion.Body = msg.Body
	newQuestion.Options = msg.Options
	newQuestion.Id = uint64(len(pool.Questions))
	pool.Questions = append(pool.Questions, &newQuestion)
	k.SetPrefPool(ctx, pool)

	return &types.MsgAddQuestionResponse{Id: newQuestion.Id}, nil
}
