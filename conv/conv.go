package conv

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit + 459.67) * 5 / 9
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return kelvin*9/5 - 459.67
}
