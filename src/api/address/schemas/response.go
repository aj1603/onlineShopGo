package schemas

type Address struct {
	ID           int    `json:"id" validate:"required,gt=0"`
	REGION       string `json:"region" validate:"required"`
	CITY         string `json:"city" validate:"required"`
	ADDRESS_LINE string `json:"address_line" validate:"required"`
	CUSTOMERS_ID int    `json:"customers_id" validate:"required,gt=0"`
}

type Product_category struct {
	ID          int     `json:"id"`
	NAME        string  `json:"name"`
	DESCRIPTION string  `json:"description"`
	PRICE       float32 `json:"price"`
	PRODUCT_SKU string  `json:"product_sku"`
	QUANTITY    int     `json:"quantity"`
	CATEGORY_ID int     `json:"categories_id"`
	DISCOUNT_ID int     `json:"discounts_id"`
	BRAND_ID    int     `json:"brands_id"`
	Images      []Image `json:"images"`
}

type Image struct {
	ID         int    `json:"id"`
	IMG_URL    string `json:"img_url"`
	PRODUCT_ID int    `json:"products_id"`
}
