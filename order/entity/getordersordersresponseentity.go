package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetOrdersOrdersResponseEntity struct {
	BranchNumber                string                                 `json:"branch_number"`
	TaxCode                     string                                 `json:"tax_code"`
	ExtraAttributes             string                                 `json:"extra_attributes"`
	AddressUpdatedAt            string                                 `json:"address_updated_at"`
	ShippingFee                 float32                                `json:"shipping_fee"`
	CustomerFirstName           string                                 `json:"customer_first_name"`
	PaymentMethod               string                                 `json:"payment_method"`
	Statuses                    []string                               `json:"statuses"`
	Remarks                     string                                 `json:"remarks"`
	OrderNumber                 int64                                  `json:"order_number"`
	OrderId                     int64                                  `json:"order_id"`
	Voucher                     float32                                `json:"voucher"`
	NationalRegistrationNumber  string                                 `json:"national_registration_number"`
	PromisedShippingTimes       string                                 `json:"promised_shipping_times"`
	ItemsCount                  int                                    `json:"items_count"`
	VoucherPlatform             float32                                `json:"voucher_platform"`
	VoucherSeller               float32                                `json:"voucher_seller"`
	CreatedAt                   string                                 `json:"created_at"`
	Price                       string                                 `json:"price"`
	AddressBilling              GetOrdersAddressBillingResponseEntity  `json:"address_billing"`
	WarehouseCode               string                                 `json:"warehouse_code"`
	ShippingFeeOriginal         float32                                `json:"shipping_fee_original"`
	ShippingFeeDiscountSeller   float32                                `json:"shipping_fee_discount_seller"`
	ShippingFeeDiscountPlatform float32                                `json:"shipping_fee_discount_platform"`
	AddressShipping             GetOrdersAddressShippingResponseEntity `json:"address_shipping"`
	CustomerLastName            string                                 `json:"customer_last_name"`
	GiftOption                  bool                                   `json:"gift_option"`
	VoucherCode                 string                                 `json:"voucher_code"`
	UpdatedAt                   string                                 `json:"updated_at"`
	DeliveryInfo                string                                 `json:"delivery_info"`
	GiftMessage                 string                                 `json:"gift_message"`
}

func (g GetOrdersOrdersResponseEntity) String() string {
	return lib.ObjectToString(g)
}
