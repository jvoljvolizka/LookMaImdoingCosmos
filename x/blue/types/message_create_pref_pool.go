package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePrefPool = "create_pref_pool"

var _ sdk.Msg = &MsgCreatePrefPool{}

func NewMsgCreatePrefPool(creator string, title string) *MsgCreatePrefPool {
  return &MsgCreatePrefPool{
		Creator: creator,
    Title: title,
	}
}

func (msg *MsgCreatePrefPool) Route() string {
  return RouterKey
}

func (msg *MsgCreatePrefPool) Type() string {
  return TypeMsgCreatePrefPool
}

func (msg *MsgCreatePrefPool) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePrefPool) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePrefPool) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

