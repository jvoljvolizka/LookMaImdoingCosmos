package keeper

import (
	"context"
	"strconv"

	"blue/x/blue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	poolCount := k.GetPoolCount(ctx)
	Id, err := strconv.ParseUint(msg.PoolId, 10, 64)
	if err != nil {
		return nil, err
	}
	if Id >= poolCount {
		return nil, status.Error(codes.InvalidArgument, "invalid Id")
	}

	Answer, err := strconv.ParseUint(msg.Answer, 10, 64)
	if err != nil {
		return nil, err
	}
	pool := k.GetPool(ctx, Id)
	if Answer >= pool.Answerrange {
		return nil, status.Error(codes.InvalidArgument, "invalid Answer")
	}
	pool.Answers = append(pool.Answers, Answer)
	k.EditPool(ctx, pool)

	return &types.MsgVoteResponse{Pool: &pool}, nil
}
