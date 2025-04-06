package elgatoring

type WifiInfo struct {
	Ssid         string `json:"ssid"`
	FrequencyMHz int    `json:"frequencyMHz"`
	Rssi         int    `json:"rssi"`
}
