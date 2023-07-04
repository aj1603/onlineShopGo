package schemas

type Category struct {
	ID       int                `json:"id"`
	NAME     string             `json:"name"`
	PRODUCTS []Product_category `json:"product_category"`
}

type Product_category struct {
	ID          int                       `json:"id"`
	NAME        string                    `json:"name"`
	DESCRIPTION string                    `json:"description"`
	PRICE       float32                   `json:"price"`
	PRODUCT_SKU string                    `json:"product_sku"`
	QUANTITY    int                       `json:"quantity"`
	CATEGORY_ID int                       `json:"categories_id"`
	DISCOUNT_ID int                       `json:"discounts_id"`
	BRAND_ID    int                       `json:"brands_id"`
	PRODUCT_IMG []Product_images_category `json:"products_images"`
}

type Product_images_category struct {
	ID         int    `json:"id"`
	IMG_URL    string `json:"img_url"`
	PRODUCT_ID int    `json:"products_id"`
}
