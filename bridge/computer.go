package bridge

type Computer interface {
	Print()
	SetPrinter(Printer)
}
