package schemas

type ProductCreateRequest struct {
	Name  string `json:"name" binding:"required" example:"iPhone 15"`
	Price int    `json:"price" binding:"required" example:"999"`
}

type ProductUpdateRequest struct {
	Name  string `json:"name" example:"iPhone 15 Pro"`
	Price int    `json:"price" example:"1099"`
}
