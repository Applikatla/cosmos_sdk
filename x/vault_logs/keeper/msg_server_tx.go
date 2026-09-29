package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pb "github.com/cosmos/cosmos-sdk/x/vault_logs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc msgServer) RegAgent(goCtx context.Context, msg *pb.RegAgentRequest) (*pb.RegAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()

	agent_id := strings.ToLower(msg.AgentId)
	agent_type := strings.ToLower(msg.AgentType)

	// check if store exist with agent id
	ok, err := svc.Keeper.RegAgent.Has(ctx, agent_id)
	if err != nil {
		logger.Error("🚫 error to check agent reg")
		return &pb.RegAgentResponse{}, status.Errorf(codes.Internal, "🚫 error to check agent reg")
	}
	if ok {
		logger.Error("🚫 agent already exist")
		return &pb.RegAgentResponse{}, status.Errorf(codes.AlreadyExists, "🚫 agent already exist")
	}
	// setuping for store
	reg := pb.AgentReg{
		AgentId:   agent_id,
		AgentType: agent_type,
	}

	if err := svc.Keeper.RegAgent.Set(ctx, agent_id, reg); err != nil {
		logger.Error("🚫 error to store agent details", "error", err)
		return &pb.RegAgentResponse{}, status.Errorf(codes.Internal, "🚫 error to store agent details: %v", err)
	}
	logger.Info("✅ Successfully reg agent", "agent id", agent_id)
	return &pb.RegAgentResponse{
		AgentId:   agent_id,
		CreatedAt: ctx.BlockTime().String(),
	}, nil
}
