package delivery

import (
	"log"
	"net/http"
	"pvg/domain"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(us domain.UserUseCase) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("cannot parse data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "data not found",
			})
		}

		dataUser := tmp.ToModel()
		row, _ := uh.userUsecase.AddUser(dataUser)
		if row == -1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    500,
				"message": "please make sure all fields are filled in correctly",
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "register success",
		})
	}
}

func (uh *userHandler) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idFromParam := c.Param("id")
		id, _ := strconv.Atoi(idFromParam)

		row, errDel := uh.userUsecase.DeleteCase(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "failed to delete data user"})
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "failed to delete data user"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "delete success",
		})
	}
}

func (uh *userHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitcnv, _ := strconv.Atoi(limit)
		offsetcnv, _ := strconv.Atoi(offset)
		result, err := uh.userUsecase.GetAllUserCase(limitcnv, offsetcnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ParsePUToArr2(result),
			"code":    200,
			"message": "get data success",
		})
	}
}

func (uh *userHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    FromModel(data),
			"code":    200,
			"message": "get data success",
		})
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		id := c.Param("id")
		idParam, _ := strconv.Atoi(id)
		err := c.Bind(&tmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "failed to bind data, check your input"})
		}
		row, _ := uh.userUsecase.UpdateCase(tmp.ToModel(), idParam)
		if row == 404 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "nothing to update data"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    400,
			"message": "success update data"})
	}
}
