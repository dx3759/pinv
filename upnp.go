package pinv

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/NebulousLabs/go-upnp"
)

type upnpstate struct {
	LocalPort  int
	RemotePort int
	Enable     bool
	ExternalIP string
}

var upnpInfo upnpstate

func registerUpnp() {
	upnpInfo = upnpstate{
		LocalPort: GloOptions.Port,
		Enable:    false,
	}

	d, err := upnp.Discover()
	if err != nil {
		logrus.Warn("upnp ", err.Error())
		return
	}

	ip, err := d.ExternalIP()
	if err != nil {
		logrus.Warn("upnp external ip ", err.Error())
		return
	}

	logrus.Debug("upnp your external ip is ", ip)

	err = d.Forward(uint16(GloOptions.Port), GloOptions.AppName())
	if err != nil {
		logrus.Warn("upnp forward ", err.Error())
		return
	}
	upnpInfo.Enable = true
	upnpInfo.ExternalIP = ip
	upnpInfo.RemotePort = upnpInfo.LocalPort

	logrus.Errorf("%+v", upnpInfo)
}
