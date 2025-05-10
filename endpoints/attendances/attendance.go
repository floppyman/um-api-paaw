package attendances

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	
	"github.com/floppyman/um-common/web"
	
	"github.com/floppyman/um-api-paaw/base"
)

func (re AttendanceEndPoint) List(startTime time.Time, endTime time.Time, includePending bool, limit int, offset int) (bool, AttendanceListResponse, error) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, AttendanceListResponse{}, err
	}
	
	if limit <= 0 || limit > 200 {
		limit = 200
	}
	if offset < 0 {
		offset = 0
	}
	
	qs := web.QueryStringBuilder(map[string]string{
		"start_time":     startTime.Format("2006-01-02"),
		"end_time":       endTime.Format("2006-01-02"),
		"includePending": strconv.FormatBool(includePending),
		"limit":          strconv.FormatInt(int64(limit), 10),
		"offset":         strconv.FormatInt(int64(offset), 10),
	})
	req := base.CreateRequest(base.HttpGet, fmt.Sprintf("/attendance/list?%s", qs), nil, true)
	
	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, AttendanceListResponse{}, err
	}
	
	var res AttendanceListResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, AttendanceListResponse{}, err2
	}
	return true, res, nil
}

func (re AttendanceEndPoint) Create(item AttendanceCreateItem) (bool, AttendanceCreateResponse, error) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, AttendanceCreateResponse{}, err
	}
	
	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return false, AttendanceCreateResponse{}, err
	}
	
	req := base.CreateRequest(base.HttpPost, "/attendance/create", bodyBytes, true)
	
	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, AttendanceCreateResponse{}, err
	}
	
	var res AttendanceCreateResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, AttendanceCreateResponse{}, err2
	}
	return true, res, nil
}

func (re AttendanceEndPoint) Update(id int, item AttendanceDataItem) (bool, AttendanceUpdateResponse, error) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, AttendanceUpdateResponse{}, err
	}
	
	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return false, AttendanceUpdateResponse{}, err
	}
	
	req := base.CreateRequest(base.HttpPatch, fmt.Sprintf("/attendance/update/%d", id), bodyBytes, true)
	
	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, AttendanceUpdateResponse{}, err
	}
	
	var res AttendanceUpdateResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, AttendanceUpdateResponse{}, err2
	}
	return true, res, nil
}

func (re AttendanceEndPoint) Delete(id int) (bool, AttendanceDeleteResponse, error) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, AttendanceDeleteResponse{}, err
	}
	
	req := base.CreateRequest(base.HttpDelete, fmt.Sprintf("/attendance/delete/%d", id), nil, true)
	
	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, AttendanceDeleteResponse{}, err
	}
	
	var res AttendanceDeleteResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, AttendanceDeleteResponse{}, err2
	}
	return true, res, nil
}
