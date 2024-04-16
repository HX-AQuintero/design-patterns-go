package main

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// ok if is necessary a new struct
type MultiFunctionPrinter struct {
	//...
}

// MultiFunctionPrinter implements Machine. No problem! 👍
func (m *MultiFunctionPrinter) Print(d Document) {
	//...
}

func (m *MultiFunctionPrinter) Fax(d Document) {
	//...
}

func (m *MultiFunctionPrinter) Scan(d Document) {
	//...
}

// let's add an old fasioned printer struct which doesn't really
// have any scanning or faxing capabilities
type OldFashionedPrinter struct {
	//...
}

// OldFashionPrinter uses Print but not Fax or Scan
func (o *OldFashionedPrinter) Print(d Document) {
	//...
}

// so, OldFashionedPrinter has to be forced to implement Machine❓
func (o *OldFashionedPrinter) Fax(d Document) {
	panic("Operation not supported")
}

func (o *OldFashionedPrinter) Scan(d Document) {
	panic("Operation not supported")
}

// that's not the way ❌
// the problem comes from putting too much into an interface
// breaking the ISP ⚠️

// a better approach: Splitting up into separate interfaces
type Printer interface {
	Print(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// MyPrinter implements Printer interface only
type MyPrinter struct {
	//...
}

func (m *MyPrinter) Print(d Document) {
	//...
}

// Photocopier implements both Printer and Scanner interfaces
type Photocopier struct {
	//...
}

func (p *Photocopier) Print(d Document) {
	//...
}

func (p *Photocopier) Scan(d Document) {
	//...
}

// combining interfaces
type MultiFunctionDevice interface {
	Printer
	Scanner
	//Faxer (just in case)
}

// interfaces + decorator pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m *MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m *MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
