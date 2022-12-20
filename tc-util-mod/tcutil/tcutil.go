// Package tcutil provides functionality of TwinCAT systems.
package tcutil

import (
	"strconv"
	"strings"
)

type Route struct {
	Name    string `xml:"Name"`    // struct tag (metadata) used for defining name when writing to xml
	Address string `xml:"Address"` // struct tag (metadata) used for defining name when writing to xml
	NetId   string `xml:"NetId"`   // struct tag (metadata) used for defining name when writing to xml
	Type    string `xml:"Type"`    // struct tag (metadata) used for defining name when writing to xml
}

type RemoteConnections struct {
	RemoteConnections []Route `xml:"Route"` // struct tag (metadata) used for defining name when writing to xml
}

type TcConfig struct {
	XmlnsXsi string            `xml:"xmlns:xsi,attr"`    // struct tag (metadata) used for defining name when writing to xml
	TcConfig RemoteConnections `xml:"RemoteConnections"` // struct tag (metadata) used for defining name when writing to xml
}

const XmlFileName string = "StaticRoutes.xml"
const XmlsXsiAttributeValue string = "http://www.w3.org/2001/XMLSchema-instance"
const XmlFileHeader string = `<?xml version="1.0" encoding="UTF-8"?>`
const NetIdSuffix string = ".1.1"
const RouteType string = "TCP_IP"

// GetTcConfigBasedOnBaseIpAddressWithFirstThreeSegments returns the twincat configuration based on the passed
// three segment ip address (e.g. 192.168.1.). The missing segment is fully enumerated.
func GetTcConfigBasedOnBaseIpAddressWithFirstThreeSegments(baseIpAddress string) TcConfig {
	remoteConnections := GetRemoteConnectionsBasedOnBaseIpAddressWithFirstThreeSegments(baseIpAddress)
	tcConfig := TcConfig{XmlnsXsi: XmlsXsiAttributeValue, TcConfig: *remoteConnections}

	return tcConfig
}

// GetRemoteConnectionsBasedOnBaseIpAddressWithFirstThreeSegments returns the remote connections based on the passed
// three segment ip address (e.g. 192.168.1.). The missing segment is fully enumerated.
func GetRemoteConnectionsBasedOnBaseIpAddressWithFirstThreeSegments(baseIpAddress string) *RemoteConnections {
	if strings.HasSuffix(baseIpAddress, ".") == false {
		baseIpAddress = baseIpAddress + "."
	}

	remoteConnections := RemoteConnections{}

	for ii := 0; ii < 256; ii++ {
		currentIPAddress := baseIpAddress + strconv.Itoa(ii)
		currentNetId := currentIPAddress + NetIdSuffix
		currentRoute := Route{}
		currentRoute.Name = currentIPAddress
		currentRoute.Address = currentIPAddress
		currentRoute.NetId = currentNetId
		currentRoute.Type = RouteType

		remoteConnections.RemoteConnections = append(remoteConnections.RemoteConnections, currentRoute)
	}

	return &remoteConnections
}
