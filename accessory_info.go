package elgatoring

type AccessoryInfo struct {
	ProductName         string   `json:"productName"`
	HardwareBoardType   int      `json:"hardwareBoardType"`
	HardwareRevision    float64  `json:"hardwareRevision"`
	MacAddress          string   `json:"macAddress"`
	FirmwareBuildNumber int      `json:"firmwareBuildNumber"`
	FirmwareVersion     string   `json:"firmwareVersion"`
	SerialNumber        string   `json:"serialNumber"`
	DisplayName         string   `json:"displayName"`
	Features            []string `json:"features"`
	WifiInfo            WifiInfo `json:"wifi-info"`
}
