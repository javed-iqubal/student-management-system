package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/javed-iqubal/student-management-system/internal/types"
	"github.com/javed-iqubal/student-management-system/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		// request validation

		if err := validator.New().Struct(student); err != nil {
			validatorErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatorErrors))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
