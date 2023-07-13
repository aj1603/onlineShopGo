package schemas

type Cart_res struct {
	ID          int              `json:"id"`
	CART_SKU    string           `json:"carts_sku"`
	CUSTOMER_ID int              `json:"customers_id"`
	CART_ITEMS  []Cart_items_res `json:"cart_items"`
}

type Cart_items_res struct {
	ID         int       `json:"id"`
	QUANTITY   int       `json:"quantity"`
	PRODUCT_ID int       `json:"products_id"`
	CARD_ID    int       `json:"carts_id"`
	PRODUCTS   []Product `json:"products"`
}

type Product struct {
	ID          int              `json:"id"`
	NAME        string           `json:"name"`
	DESCRIPTION string           `json:"description"`
	PRICE       float32          `json:"price"`
	PRODUCT_SKU string           `json:"product_sku"`
	QUANTITY    int              `json:"quantity"`
	CATEGORY_ID int              `json:"categories_id"`
	DISCOUNT_ID int              `json:"discounts_id"`
	BRAND_ID    int              `json:"brands_id"`
	PRODUCT_IMG []Product_images `json:"products_images"`
}

type Product_images struct {
	ID         int    `json:"id"`
	IMG_URL    string `json:"img_url"`
	PRODUCT_ID int    `json:"products_id"`
}
