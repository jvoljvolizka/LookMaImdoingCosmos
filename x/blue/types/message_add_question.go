package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddQuestion = "add_question"

var _ sdk.Msg = &MsgAddQuestion{}

func NewMsgAddQuestion(creator string, body string) *MsgAddQuestion {
	return &MsgAddQuestion{
		Creator: creator,
		Body:    body,
	}
}

func (msg *MsgAddQuestion) Route() string {
	return RouterKey
}

func (msg *MsgAddQuestion) Type() string {
	return TypeMsgAddQuestion
}

func (msg *MsgAddQuestion) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddQuestion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddQuestion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
