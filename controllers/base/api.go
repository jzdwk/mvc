package base

import (
	"github.com/mvc/models"
)

var (
	PublishRequestMessageMethodFilter = []string{
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
	}
)

type ApiController struct {
	ParamBuilderController
	User *models.User
}

func (c *ApiController) Prepare() {
	//base auth token check
	/*authString := c.Ctx.Input.Header("Authorization")

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		logs.Info("AuthString invalid:", authString)
		c.CustomAbort(http.StatusUnauthorized, "Token invalid!")
	}
	tokenString := kv[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(beego.AppConfig.String("TokenSecrets")), nil
	})

	errResult := errors.ErrorResult{}
	switch err.(type) {
	case nil: // no error
		if !token.Valid { // but may still be invalid
			errResult.Code = http.StatusUnauthorized
			errResult.Msg = "Token Invalid ! "
		}
	case *jwt.ValidationError: // something was wrong during the validation
		errResult.Code = http.StatusUnauthorized
		errResult.Msg = err.Error()

	default: // something else went wrong
		errResult.Code = http.StatusInternalServerError
		errResult.Msg = err.Error()
	}

	if err != nil || !token.Valid {
		c.CustomAbort(errResult.Code, errResult.Msg)
	}
	claim := token.Claims.(jwt.MapClaims)
	aud := claim["userName"].(string)
	c.User, err = models.UserModel.GetUserDetail(aud)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}*/
}
