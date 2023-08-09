package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/circlefin/noble-cctp/x/cctp/types"
)

func (k msgServer) PauseSendingAndReceivingMessages(goCtx context.Context, msg *types.MsgPauseSendingAndReceivingMessages) (*types.MsgPauseSendingAndReceivingMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, found := k.GetAuthority(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAuthorityNotSet, "authority is not set")
	}

	if owner.Address != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot pause sending and receiving messages")
	}

	paused := types.SendingAndReceivingMessagesPaused{
		Paused: true,
	}
	k.SetSendingAndReceivingMessagesPaused(ctx, paused)

	event := types.SendingAndReceivingPausedEvent{}
	err := ctx.EventManager().EmitTypedEvent(&event)
	return &types.MsgPauseSendingAndReceivingMessagesResponse{}, err
}
