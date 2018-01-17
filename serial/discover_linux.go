package serial

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/yookoala/realpath"
)

func readLine(fn string) string {
	f, err := os.Open(fn)
	if err == nil {
		scan := bufio.NewScanner(f)
		if scan.Scan() {
			return strings.TrimSpace(scan.Text())
		}
	}
	return ""
}

func Discover() ([]DiscoverInfo, error) {

	var devs []string

	s1, _ := filepath.Glob("/dev/ttyS*")
	s2, _ := filepath.Glob("/dev/ttyUSB*")
	devs = append(devs, s1...)
	devs = append(devs, s2...)

	var infos []DiscoverInfo

	for _, d := range devs {
		info := DiscoverInfo{
			PortName: d,
		}
		name := filepath.Base(d)
		subsystem := ""
		devpath := fmt.Sprintf("/sys/class/tty/%s/device", name)
		if _, err := os.Lstat(devpath); err == nil {
			devpath, _ = realpath.Realpath(devpath)
			subsystem = filepath.Join(devpath, "subsystem")
			subsystem, _ = realpath.Realpath(subsystem)
			subsystem = filepath.Base(subsystem)
		}

		usbpath := ""
		switch subsystem {
		case "usb-serial":
			usbpath = filepath.Dir(filepath.Dir(devpath))
			info.Type = PortTypeUsb
		case "usb":
			usbpath = filepath.Dir(devpath)
			info.Type = PortTypeUsb
		case "pnp":
			info.Type = PortTypeIntegrated
		case "pci":
			// We recognize the subsystem but we don't have anything to do
		case "platform":
			// Not a real serial port, just exposed by the driver
			continue
		default:
			fmt.Println("unknown subsystem", subsystem)
		}

		switch info.Type {
		case PortTypeUsb:
			info.Description = readLine(filepath.Join(usbpath, "product"))
			idvendor, _ := strconv.Atoi(readLine(filepath.Join(usbpath, "idVendor")))
			info.UsbVendorId = idvendor
			idproduct, _ := strconv.Atoi(readLine(filepath.Join(usbpath, "idProduct")))
			info.UsbProductId = idproduct
			info.UsbSerialNo = readLine(filepath.Join(usbpath, "serial"))
		}

		infos = append(infos, info)
	}

	return infos, nil
}
