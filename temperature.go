package elgatoring

const (
	MinTemperature Temperature = 143
	MaxTemperature Temperature = 344
)

type Temperature int

func (t Temperature) Kelvin() int {
	return 1000000 / int(t)
}

func (t Temperature) Mireds() int {
	return int(t)
}

func TemperatureFromKelvin(k int) Temperature {
	return Temperature(1000000 / k)
}

func TemperatureFromMireds(mireds int) Temperature {
	return Temperature(mireds)
}
