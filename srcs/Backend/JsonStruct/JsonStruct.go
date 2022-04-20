package JsonStruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type JsonStruct struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

// Value Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (j JsonStruct) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (j *JsonStruct) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &j)
}

func (j *JsonStruct) FillFromFile(pathOfFile string) error {
	open, ok := os.Open(pathOfFile)
	if ok != nil {
		return ok
	}
	decoder := json.NewDecoder(open)
	ok = decoder.Decode(j)
	if ok != nil {
		return ok
	}
	ok = open.Close()
	return ok
}

func NewFromFile(pathOfFile string) (*JsonStruct, error) {
	var js JsonStruct
	ok := js.FillFromFile(pathOfFile)
	return &js, ok
}

func ParseBytes(data []byte) (js JsonStruct, ok error) {
	ok = json.Unmarshal(data, &js)
	return js, ok
}