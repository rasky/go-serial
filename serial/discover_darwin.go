package serial

import (
	"bytes"
	"os/exec"

	"github.com/DHowett/go-plist"
)

type cIOObjectClass struct {
	Class         string `plist:"IOObjectClass"`
	ClassOverride string `plist:"IOClassNameOverride"`
	Dialin        string `plist:"IODialinDevice"`
	Callout       string `plist:"IOCalloutDevice"`
	TtySuffix     string `plist:"IOTTYSuffix"`
	TtyDevice     string `plist:"IOTTYDevice"`

	UsbProductName  string `plist:"USB Product Name"`
	UsbIdVendor     int    `plist:"idVendor"`
	UsbIdProduct    int    `plist:"idProduct"`
	UsbSerialNumber string `plist:"kUSBSerialNumberString"`

	Child *cIOObjectClass `plist:"IORegistryEntryChildren"`
}

func Discover() ([]DiscoverInfo, error) {
	out, err := exec.Command("ioreg", "-artlc", "IOSerialBSDClient").Output()
	if err != nil {
		return []DiscoverInfo{}, err
	}

	var data []cIOObjectClass

	dec := plist.NewDecoder(bytes.NewReader(out))
	err = dec.Decode(&data)
	if err != nil {
		return []DiscoverInfo{}, err
	}

	var ports []DiscoverInfo

	for i := range data {
		dev := &data[i]
		info := DiscoverInfo{}
		for dev.Child != nil {
			if dev.Class == "IOBluetoothSerialClient" {
				info.Type = PortTypeBluetooth
			} else if dev.ClassOverride == "IOUSBDevice" {
				info.Type = PortTypeUsb
				info.Description = dev.UsbProductName
				info.UsbVendorId = dev.UsbIdVendor
				info.UsbProductId = dev.UsbIdProduct
				info.UsbSerialNo = dev.UsbSerialNumber
			}
			dev = dev.Child
		}

		info.PortName = dev.Dialin
		if info.Description == "" {
			info.Description = dev.TtyDevice
		}

		ports = append(ports, info)
	}

	return ports, nil
}
