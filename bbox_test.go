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
	device, err := bbox.GetDevice()
	if err != nil {
		logrus.WithError(err).Error("cannot get device")
	}
	fmt.Println(">>", device)
}

func TestWanIpStats(t *testing.T) {
	stats, err := bbox.GetWanIpStats()
	if err != nil {
		logrus.WithError(err).Error("cannot get ip stats")
	}
	fmt.Println(">>", stats)
}

func TestWanXdsl(t *testing.T) {
	stats, err := bbox.GetWanXdsl()
	if err != nil {
		logrus.WithError(err).Error("cannot get wan xdsl")
	}
	fmt.Println(">>", stats)
}

func TestWanIp(t *testing.T) {
	stats, err := bbox.GetWanIp()
	if err != nil {
		logrus.WithError(err).Error("cannot get wan ip")
	}
	fmt.Println(">>", stats)
}
