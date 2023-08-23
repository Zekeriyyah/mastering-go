package main

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) farenheit() farenheit {
	return farenheit((1.8 * c) + 32)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) farenheit() farenheit {
	v := k.celsius()
	return v.farenheit()
}

func (f farenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

func (f farenheit) kelvin() kelvin {
	v := f.celsius()
	return v.kelvin()
}
