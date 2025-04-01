package player

import (
	"errors"

	"gitlab.casinovip.tech/minigame_backend/c_engine/pkg/utils/u_jwt"
)

func TokenValid(c Player, body LoginReqBody) (*u_jwt.Claims, error) {
	if body.Token == "" {
		return nil, errors.New("token is empty")
	}

	oClaims, err := u_jwt.ParseToken(body.Token, constant.JwtSalt)
	if err != nil {
		return nil, err
	}

	if av, ok := c.(*Player); ok {
		av.Id = oClaims.Id
		av.SetDevice(body.Device)
	}

	return oClaims, nil
}
