func ServiceErrorToErrorResp(err error) *handlerwrap.ErrorResponse {
	switch {
	case errors.Is(err, service.ErrRecordNotFound):
		return handlerwrap.NewErrorResponse(
			err,
			http.StatusNotFound,
			"record_not_found",
			"record not found",
		)
	case errors.Is(err, service.ErrInvalidColorEnum):
		return handlerwrap.NewErrorResponse(
			errC,
			http.StatusBadRequest,
			"color_enum_invalid",
			fmt.Sprintf("the selected color %s does not exist found", errC.Color)
		)
	default:
		return handlerwrap.InternalServerError{Err: err}.ToErrorResponse()
	}
}

