package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateStudent = "create_student"

var _ sdk.Msg = &MsgCreateStudent{}

func NewMsgCreateStudent(creator, name string, age uint64) *MsgCreateStudent {
	return &MsgCreateStudent{
		Creator: creator,
		Name:    name,
		Age:     age,
	}
}

func (msg *MsgCreateStudent) Route() string { return RouterKey }

func (msg *MsgCreateStudent) Type() string { return TypeMsgCreateStudent }

func (msg *MsgCreateStudent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStudent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap("invalid creator address: " + err.Error())
	}
	if msg.Name == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("name cannot be empty")
	}
	if msg.Age == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("age must be greater than zero")
	}
	return nil
}
