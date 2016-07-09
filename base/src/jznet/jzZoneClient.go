package jznet

type ZoneClient struct {
	*NetClient
}

func NewZoneClient(subNetService INetClientInterface) *ZoneClient {
	return &ZoneClient {
		NetClient:NewNetClient(subNetService),
	}
} 
