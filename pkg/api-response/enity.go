package apiresponse

/*
   |--------------------------------------------------------------------------
   | Default Struct Response API
   |--------------------------------------------------------------------------
   |
   | You can chang this every momen you want    |
*/

type Meta struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Code    string      `json:"code"`
}

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}