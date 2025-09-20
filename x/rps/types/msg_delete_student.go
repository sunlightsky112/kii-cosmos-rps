package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteStudent = "delete_student"

var _ sdk.Msg = &MsgDeleteStudent{}

func NewMsgDeleteStudent(creator, id string) *MsgDeleteStudent {
	return &MsgDeleteStudent{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteStudent) Route() string { return RouterKey }

func (msg *MsgDeleteStudent) Type() string { return TypeMsgDeleteStudent }

func (msg *MsgDeleteStudent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStudent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap("invalid creator address" + err.Error())
	}
	if msg.Id == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("id cannot be empty")
	}
	return nil
}
