package bbox

import (
	"testing"
	"github.com/Sirupsen/logrus"
	"fmt"
)

var bbox *Bbox

func init() {
	bbox = New()
	bbox.Host = "192.168.40.2"
	logrus.SetLevel(logrus.DebugLevel)
}

func TestDevice(t *testing.T) {
	device, err := bbox.getDevice()
	if err != nil {
		logrus.WithError(err).Error("cannot get device")
	}
	fmt.Println(">>", device)
}

func TestWanIpStats(t *testing.T) {
	stats, err := bbox.getWanIpStats()
	if err != nil {
		logrus.WithError(err).Error("cannot get ip stats")
	}
	fmt.Println(">>", stats)
}

func TestWanXdsl(t *testing.T) {
	stats, err := bbox.getWanXdsl()
	if err != nil {
		logrus.WithError(err).Error("cannot get wan xdsl")
	}
	fmt.Println(">>", stats)
}

func TestWanIp(t *testing.T) {
	stats, err := bbox.getWanIp()
	if err != nil {
		logrus.WithError(err).Error("cannot get wan ip")
	}
	fmt.Println(">>", stats)
}
