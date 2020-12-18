package api

type customerRequest struct {
	CustomerID int32 `uri:"id"`
}

type listCustomerRequest struct {
	PageSize int32 `form:"page_size"`
}
