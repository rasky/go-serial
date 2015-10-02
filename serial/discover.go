package serial

//go:generate stringer -type PortType

// The type of the physical device that exposes the serial port
type PortType int

const (
	PortTypeUnknown PortType = iota
	PortTypeUsb
	PortTypeBluetooth
	PortTypePci
	PortTypeIntegrated
)

type DiscoverInfo struct {
	// Name of the serial port, e.g. "/dev/tty.usbserial-A8008HlV".
	PortName string

	// A human-readable description of the port. Empty if the
	// discover process wasn't able to find a description which
	// is not more useful than the PortName itself.
	Description string

	// The type of the port
	Type PortType

	// Additional USB informations. These fields are filled in
	// only if Type == PortTypeUsb
	UsbVendorId  int
	UsbProductId int
	UsbSerialNo  string
}
