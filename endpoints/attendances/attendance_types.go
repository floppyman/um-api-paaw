package attendances

import "time"

type AttendanceEndPoint struct {
}

type AttendanceDataItem struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Break     int    `json:"break"`
	Comment   string `json:"comment"`
}

type AttendanceCreateItem struct {
	Attendances []AttendanceDataItem `json:"attendances"`
}

type AttendanceListResponse struct {
	Success  bool                           `json:"success"`
	Metadata AttendanceListResponseMetadata `json:"metadata"`
	Data     []AttendanceListResponseData   `json:"data"`
	Offset   string                         `json:"offset"`
	Limit    string                         `json:"limit"`
}
type AttendanceListResponseMetadata struct {
	TotalElements int `json:"total_elements"`
	CurrentPage   int `json:"current_page"`
	TotalPages    int `json:"total_pages"`
}
type AttendanceListResponseData struct {
	Id         int                                  `json:"id"`
	Type       string                               `json:"type"`
	Attributes AttendanceListResponseDataAttributes `json:"attributes"`
}
type AttendanceListResponseDataAttributes struct {
	Employee    int         `json:"employee"`
	Date        string      `json:"date"`
	StartTime   string      `json:"start_time"`
	EndTime     string      `json:"end_time"`
	Break       int         `json:"break"`
	Comment     string      `json:"comment"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Status      string      `json:"status"`
	Project     interface{} `json:"project"`
	IsHoliday   bool        `json:"is_holiday"`
	IsOnTimeOff bool        `json:"is_on_time_off"`
}

type AttendanceCreateResponse struct {
	Success bool                           `json:"success"`
	Data    *AttendanceCreateResponseData  `json:"data"`
	Error   *AttendanceCreateResponseError `json:"error"`
}
type AttendanceCreateResponseData struct {
	Id      []int  `json:"id"`
	Message string `json:"message"`
}
type AttendanceCreateResponseError struct {
	Code    []int  `json:"code"`
	Message string `json:"message"`
}

type AttendanceUpdateResponse struct {
	Success bool          `json:"success"`
	Data    []interface{} `json:"data"`
}

type AttendanceDeleteResponse struct {
	Success bool                           `json:"success"`
	Data    *AttendanceDeleteResponseData  `json:"data"`
	Error   *AttendanceDeleteResponseError `json:"error"`
}
type AttendanceDeleteResponseData struct {
	Message string `json:"message"`
}
type AttendanceDeleteResponseError struct {
	Code    []int  `json:"code"`
	Message string `json:"message"`
}
