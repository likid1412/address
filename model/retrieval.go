package model

// AddressInfoRow represents an address and property information.
// It can be serialized into JSON and used for CSV operations.
type AddressInfoRow struct {
	// Unique identifier for the address
	ID string `json:"id,omitempty"`
	// Complete address
	FullAddress string `json:"full_address,omitempty"`
	// Prefecture (Japanese: 都道府県)
	Prefecture string `json:"prefecture,omitempty" csv:"都道府県"`
	// City (Japanese: 市区町村)
	City string `json:"city,omitempty" csv:"市区町村"`
	// Town (Japanese: 町名)
	Town string `json:"town,omitempty" csv:"町名"`
	// Block number (Japanese: 丁目)
	Chome string `json:"chome,omitempty" csv:"丁目"`
	// Lot number (Japanese: 番地)
	Banchi string `json:"banchi,omitempty" csv:"番地"`
	// Number (Japanese: 号)
	Go string `json:"go,omitempty" csv:"号"`
	// Building name (Japanese: 建物名)
	Building string `json:"building,omitempty" csv:"建物名"`
	// Price (Japanese: 価格)
	Price string `json:"price,omitempty" csv:"価格"`
	// Nearest station (Japanese: 最寄駅)
	NearestStation string `json:"nearest_station,omitempty" csv:"最寄駅"`
	// Property type (Japanese: 物件タイプ)
	PropertyType string `json:"property_type,omitempty" csv:"物件タイプ"`
	// Land area (Japanese: 敷地面積)
	LandArea string `json:"land_area,omitempty" csv:"敷地面積"`
}
