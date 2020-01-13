package addsigners

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "bank" type messages.
func NewHandler() sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgAddSigners:
			return handleMsgAddSigners(ctx, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message type: %T", msg)
		}
	}
}

// Handle MsgSend.
func handleMsgAddSigners(ctx sdk.Context, msg MsgAddSigners) (*sdk.Result, error) {
	return &sdk.Result{}, nil
}
