package nct

import (
	"dna"
	"encoding/json"
)

// APIDeviceInfo represents deviceinfo param from url builder.
type APIDeviceInfo struct {
	DeviceId     dna.String `json:"DeviceID"`
	OSName       dna.String `json:"OsName"`
	OSVersion    dna.String `json:"OsVersion"`
	AppName      dna.String `json:"AppName"`
	AppVersion   dna.String `json:"AppVersion"`
	UserInfo     dna.String `json:"UserInfo"`
	LocationInfo dna.String `json:"LocationInfo"`
}

// NewAPIDeviceInfo returns new APIDeviceInfo.
func NewAPIDeviceInfo() *APIDeviceInfo {
	device := new(APIDeviceInfo)
	device.DeviceId = "90c18c4cb3c37d442e8386631d46b46f"
	device.OSName = "ANDROID"
	device.OSVersion = "10"
	device.AppName = "NhacCuaTui"
	device.AppVersion = "5.0.1"
	device.UserInfo = ""
	device.LocationInfo = ""
	return device
}

// String converts device struct to JSON string.
func (device *APIDeviceInfo) String() string {
	bytes, err := json.Marshal(*device)
	if err == nil {
		return string(bytes)
	} else {
		return `{"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}`
	}
}
