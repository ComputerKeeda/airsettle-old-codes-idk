package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitValidatorVote = "submit_validator_vote"

var _ sdk.Msg = &MsgSubmitValidatorVote{}

func NewMsgSubmitValidatorVote(creator string, pollId string, vote bool) *MsgSubmitValidatorVote {
	return &MsgSubmitValidatorVote{
		Creator: creator,
		PollId:  pollId,
		Vote:    vote,
	}
}

func (msg *MsgSubmitValidatorVote) Route() string {
	return RouterKey
}

func (msg *MsgSubmitValidatorVote) Type() string {
	return TypeMsgSubmitValidatorVote
}

func (msg *MsgSubmitValidatorVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitValidatorVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitValidatorVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
