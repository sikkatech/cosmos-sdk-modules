package addsigners

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


var (
	addrs      = []sdk.AccAddress{
		sdk.AccAddress("test1"),
		sdk.AccAddress("test2"),
	}
)


// test ValidateBasic for MsgDeposit
func TestMsgVote(t *testing.T) {
	tests := []struct {
		signers []sdk.AccAddress
	}{
		{addrs},
		{[]sdk.AccAddress{addrs[0]}},
	}

	for i, tc := range tests {
		msg := NewMsgVote(tc.voterAddr, tc.proposalID, tc.option)
		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", i)
		}
	}
}