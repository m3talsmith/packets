package api

import (
	"context"

	"github.com/google/gopacket/pcap"
)

type Service struct {
	UnimplementedApiServer
}

func (s *Service) GetDevices(ctx context.Context, r *GetDevicesRequest) (*GetDevicesResponse, error) {
	devs, err := pcap.FindAllDevs()
	if err != nil {
		return nil, err
	}

	var devices []*Device
	for _, d := range devs {
		dev := &Device{
			Name:        d.Name,
			Description: d.Description,
			Flags:       int32(d.Flags),
		}
		for _, a := range d.Addresses {
			dev.Addresses = append(dev.Addresses, &Address{
				Ip:        a.IP.String(),
				Netmask:   a.Netmask.String(),
				Broadcast: a.Broadaddr.String(),
				P2P:       a.P2P.String(),
			})
		}
		devices = append(devices, dev)
	}
	return &GetDevicesResponse{Devices: devices}, nil
}
