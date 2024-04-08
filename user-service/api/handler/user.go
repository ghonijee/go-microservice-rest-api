package handler

import (
	"encoding/json"
	"go-microservice/user-service/api/dto"
	"go-microservice/user-service/internal/common"
	"go-microservice/user-service/internal/user"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dataRequest dto.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&dataRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = common.Validate.Struct(dataRequest)
	if err != nil {
		errors := common.GetValidationErrors(err)
		common.ErrorInvalidData(w, common.ErrorResponse{
			Errors: errors,
		})
		return
	}

	hashPass, err := common.HashPassword(dataRequest.Password)
	common.LogIfError(err)

	err = user.CreateUser(user.User{
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: hashPass,
	})

	if err != nil {
		common.ErrorBadRequest(w, err.Error())
		return
	}

	common.LogIfError(err)

	common.SuccessCreateData(w, common.SuccessResponse{})
}

func GetDetailUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(userId, 0, 32)
	common.LogIfError(err)
	user, err := user.FindById(int(id))

	if err != nil {
		common.ErrorInternalServer(w, err.Error())
		return
	}

	common.SuccessJson(w, common.SuccessResponse{
		Data: user,
	})
}
