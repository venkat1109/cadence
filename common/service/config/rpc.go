// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package config

import (
	"fmt"
	"net"

	"github.com/uber-common/bark"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

// RPCFactory is an implementation of service.RPCFactory interface
type RPCFactory struct {
	config      *RPC
	serviceName string
	ch          *tchannel.ChannelTransport
	logger      bark.Logger
}

// NewFactory builds a new RPCFactory
// conforming to the underlying configuration
func (cfg *RPC) NewFactory(sName string, logger bark.Logger) *RPCFactory {
	return newRPCFactory(cfg, sName, logger)
}

func newRPCFactory(cfg *RPC, sName string, logger bark.Logger) *RPCFactory {
	factory := &RPCFactory{config: cfg, serviceName: sName, logger: logger}
	return factory
}

// CreateDispatcher creates a dispatcher for inbound
func (d *RPCFactory) CreateDispatcher() *yarpc.Dispatcher {
	// Setup dispatcher for onebox
	var err error
	hostAddress := fmt.Sprintf("%v:%v", d.getListenIP(), d.config.Port)
	d.ch, err = tchannel.NewChannelTransport(
		tchannel.ServiceName(d.serviceName),
		tchannel.ListenAddr(hostAddress))
	if err != nil {
		d.logger.WithField("error", err).Fatal("Failed to create transport channel")
	}
	d.logger.Infof("Created RPC dispatcher for '%v' and listening at '%v'",
		d.serviceName, hostAddress)
	return yarpc.NewDispatcher(yarpc.Config{
		Name:     d.serviceName,
		Inbounds: yarpc.Inbounds{d.ch.NewInbound()},
	})
}

// CreateDispatcherForOutbound creates a dispatcher for outbound connection
func (d *RPCFactory) CreateDispatcherForOutbound(
	callerName, serviceName, hostName string) *yarpc.Dispatcher {
	// Setup dispatcher(outbound) for onebox
	d.logger.Infof("Created RPC dispatcher outbound for service '%v' for host '%v'",
		serviceName, hostName)
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: callerName,
		Outbounds: yarpc.Outbounds{
			serviceName: {Unary: d.ch.NewSingleOutbound(hostName)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		d.logger.WithField("error", err).Fatal("Failed to create outbound transport channel")
	}
	return dispatcher
}

func (d *RPCFactory) getListenIP() net.IP {
	if d.config.BindOnLocalHost && len(d.config.BindOnIP) > 0 {
		d.logger.Fatalf("ListenIP failed, bindOnLocalHost and bindOnIP are mutually exclussive")
	}

	if d.config.BindOnLocalHost {
		return net.IPv4(127, 0, 0, 1)
	}

	if len(d.config.BindOnIP) > 0 {
		ip = net.ParseIP(d.config.BindOnIP)
		if ip != nil && ip.To4() != nil {
			return ip.To4()
		}
		d.logger.Fatalf("ListenIP failed, unable to parse bindOnIP value %q or it is not IPv4 address", d.config.BindOnIP)
	}
	ip, err := ListenIP()
	if err != nil {
		d.logger.Fatalf("ListenIP failed, err=%v", err)
	}
	return ip
}
