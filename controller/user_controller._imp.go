package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang_jwt/helper"
	"golang_jwt/model/web"
	"golang_jwt/service"
	"strconv"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

func (controller *userControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Register(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userLoginResponse := controller.UserService.Login(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userLoginResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLogoutRequest := web.UserLogoutRequest{}
	helper.ReadFromRequestBody(request, &userLogoutRequest)

	userLogoutResponse := controller.UserService.Logout(request.Context(), userLogoutRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userLogoutResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) RenewAccessToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	renewAccessTokenRequest := web.RenewAccessTokenRequest{}
	helper.ReadFromRequestBody(request, &renewAccessTokenRequest)

	renewAccessTokenResponse := controller.UserService.RenewAccessToken(request.Context(), renewAccessTokenRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   renewAccessTokenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) RevokeSession(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRevokeSessionRequest := web.UserRevokeSessionRequest{}
	helper.ReadFromRequestBody(request, &userRevokeSessionRequest)

	userRevokeSessionResponse := controller.UserService.RevokeSession(request.Context(), userRevokeSessionRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data:   userRevokeSessionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("userId"))
	helper.ErrorConditionCheck(err)

	userResponse := controller.UserService.FindById(request.Context(), userId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}