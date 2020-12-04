package main

type computer interface {
	print()
	setPrinter()
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	mac := &mac{printer: hpPrinter}
	mac.print()

	window := &window{printer: epsonPrinter}
	window.print()

	mac.setPrinter(epsonPrinter)
	mac.print()
}
