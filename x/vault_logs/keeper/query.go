package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pb "github.com/cosmos/cosmos-sdk/x/vault_logs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAgentById(goCtx context.Context, req *pb.GetAgentByIdRequest) (*pb.GetAgentByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()
	agent_id := req.GetAgentId()
	if agent_id == "" {
		return &pb.GetAgentByIdResponse{}, status.Error(codes.InvalidArgument, "🚫 agent_id can't be empty")
	}

	ok, err := k.RegAgent.Has(ctx, req.AgentId)
	if err != nil {
		logger.Error("🚫 error find agent existency: %v", err)
		return &pb.GetAgentByIdResponse{}, status.Error(codes.Internal, fmt.Sprintf("🚫 error find agent existency: %v", err))
	}
	if !ok {
		return &pb.GetAgentByIdResponse{}, status.Error(codes.NotFound, "🚫 agent not found")
	}

	agent, err := k.RegAgent.Get(ctx, agent_id)
	if err != nil {
		logger.Error("🚫 error to find agent: %v", err)
		return &pb.GetAgentByIdResponse{}, status.Error(codes.NotFound, fmt.Sprintf("🚫 error to find agent: %v", err))
	}
	logger.Info("✅ Sucessfully get agent: %s", agent_id)
	return &pb.GetAgentByIdResponse{
		Agent: &agent,
	}, nil
}
