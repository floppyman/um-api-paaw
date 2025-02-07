package attendances

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/umbrella-sh/um-api-paaw/base"
	"github.com/umbrella-sh/um-common/web"
)

func (re AttendanceEndPoint) List(startTime time.Time, endTime time.Time, includePending bool, limit int, offset int) (bool, error, AttendanceListResponse) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, err, AttendanceListResponse{}
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
		return false, err, AttendanceListResponse{}
	}

	var res AttendanceListResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err2, AttendanceListResponse{}
	}
	return true, nil, res
}

func (re AttendanceEndPoint) Create(item AttendanceCreateItem) (bool, error, AttendanceCreateResponse) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, err, AttendanceCreateResponse{}
	}

	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return false, err, AttendanceCreateResponse{}
	}

	req := base.CreateRequest(base.HttpPost, "/attendance/create", bodyBytes, true)

	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, err, AttendanceCreateResponse{}
	}

	var res AttendanceCreateResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err2, AttendanceCreateResponse{}
	}
	return true, nil, res
}

func (re AttendanceEndPoint) Update(id int, item AttendanceDataItem) (bool, error, AttendanceUpdateResponse) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, err, AttendanceUpdateResponse{}
	}

	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return false, err, AttendanceUpdateResponse{}
	}

	req := base.CreateRequest(base.HttpPatch, fmt.Sprintf("/attendance/update/%d", id), bodyBytes, true)

	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, err, AttendanceUpdateResponse{}
	}

	var res AttendanceUpdateResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err2, AttendanceUpdateResponse{}
	}
	return true, nil, res
}

func (re AttendanceEndPoint) Delete(id int) (bool, error, AttendanceDeleteResponse) {
	if ok, err := base.ValidateOrGetToken(); !ok {
		return false, err, AttendanceDeleteResponse{}
	}

	req := base.CreateRequest(base.HttpDelete, fmt.Sprintf("/attendance/delete/%d", id), nil, true)

	ok1, body, err := base.DoRequest(req)
	if !ok1 {
		return false, err, AttendanceDeleteResponse{}
	}

	var res AttendanceDeleteResponse
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err2, AttendanceDeleteResponse{}
	}
	return true, nil, res
}
