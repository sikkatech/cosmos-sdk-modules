package addsigners

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Governance message types and routes
const (
	ModuleName = "addsigners"
)

var _ sdk.Msg = MsgAddSigners{}

// MsgAddSigners defines a message that just returns some signers when calling `GetSigners()`
// This can be added to a Tx to for example, change the fee payer by adding this as the first msg.
type MsgAddSigners struct {
	Signers []sdk.AccAddress `json:"signers" yaml:"signers"` //  Address of additional signers to add
}

// NewMsgAddSigners creates a new MsgAddSigners instance
func NewMsgAddSigners(signers []sdk.AccAddress) MsgAddSigners {
	return MsgAddSigners{Signers: signers}
}

// Route implements Msg
func (msg MsgAddSigners) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgAddSigners) Type() string { return ModuleName }

// ValidateBasic implements Msg
func (msg MsgAddSigners) ValidateBasic() error {
	if len(msg.Signers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "addsigners msg has no addresses")
	}
	for _, signer := range msg.Signers {
		if signer.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, signer.String())
		}
	}
	return nil
}

// GetSignBytes implements Msg
func (msg MsgAddSigners) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgAddSigners) GetSigners() []sdk.AccAddress {
	return msg.Signers
}
