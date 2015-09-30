package serial

import "path/filepath"

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
		infos = append(infos, info)
	}

	return infos, nil
}
