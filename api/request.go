package api

type listApplicantRequest struct {
	status string `uri:"status"`
}

type listCustomerRequest struct {
	PageSize int32 `form:"page_size"`
}
