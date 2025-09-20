package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// CreateStudent stores a new student
func (m msgServer) CreateStudent(goCtx context.Context, msg *types.MsgCreateStudent) (*types.MsgCreateStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	student := types.Student{
		Id:        msg.Creator,
		Name:      msg.Name,
		Age:       msg.Age,
		CreatedAt: ctx.BlockHeight(),
	}

	if err := student.Validate(); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if err := m.SetStudent(goCtx, student); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("create_student",
			sdk.NewAttribute("id", student.Id),
			sdk.NewAttribute("name", student.Name),
			sdk.NewAttribute("age", fmt.Sprintf("%d", student.Age)),
		),
	)

	return &types.MsgCreateStudentResponse{}, nil
}

// DeleteStudent removes a student by ID
func (m msgServer) DeleteStudent(goCtx context.Context, msg *types.MsgDeleteStudent) (*types.MsgDeleteStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := m.Keeper.GetStudent(goCtx, msg.Id)
	if err != nil {
		return nil, sdkerrors.ErrKeyNotFound.Wrap("student not found")
	}

	if err := m.Keeper.DeleteStudent(goCtx, msg.Id); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("delete_student",
			sdk.NewAttribute("id", msg.Id),
		),
	)

	return &types.MsgDeleteStudentResponse{}, nil
}
