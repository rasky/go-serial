// generated by stringer -type PortType; DO NOT EDIT

package serial

import "fmt"

const _PortType_name = "PortTypeUnknownPortTypeUsbPortTypeBluetoothPortTypePciPortTypeIntegrated"

var _PortType_index = [...]uint8{0, 15, 26, 43, 54, 72}

func (i PortType) String() string {
	if i < 0 || i >= PortType(len(_PortType_index)-1) {
		return fmt.Sprintf("PortType(%d)", i)
	}
	return _PortType_name[_PortType_index[i]:_PortType_index[i+1]]
}
