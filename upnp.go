package ymfile

import (
	"fmt"

	"github.com/prestonTao/upnp"
)

type upnpstate struct {
	LocalPort  int
	RemotePort int
	Enable     bool
	ExternalIP string
}

var upnpState upnpstate

func initUpnp() {
	upnpState = upnpstate{
		LocalPort: GloOptions.Port,
		Enable:    false,
	}

	mapping := new(upnp.Upnp)
	if err := mapping.AddPortMapping(8080, 8080, "TCP"); err != nil {
		fmt.Println("fail !")
	}

	fmt.Println("success !")
	upnpState.Enable = true
}
