// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package driver

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"
	"os"

	"github.com/kelda-inc/hotrod-driver/driver/thrift-gen/driver"
	"time"
)

// Client is a remote client that implements driver.Interface
type Client struct {
	ch     *tchannel.Channel
	client driver.TChanDriver
}

// NewClient creates a new driver.Client
func NewClient() *Client {
	channelOpts := &tchannel.ChannelOptions{
		//Tracer: tracer,
	}
	ch, err := tchannel.NewChannel("driver-client", channelOpts)
	if err != nil {
		log.WithError(err).Fatal("Cannot create TChannel")
	}

	driverHost := os.Getenv("HOTROD_DRIVER_HOST")
	if driverHost == "" {
		driverHost = "hotrod-driver"
	}
	driverHost += ":8082"

	clientOpts := &thrift.ClientOptions{
		HostPort: driverHost,
	}
	thriftClient := thrift.NewClient(ch, "driver", clientOpts)
	client := driver.NewTChanDriverClient(thriftClient)

	return &Client{
		ch:     ch,
		client: client,
	}
}

// FindNearest implements driver.Interface#FindNearest as an RPC
func (c *Client) FindNearest(ctx context.Context, location string) ([]Driver, error) {
	log.WithField("location", location).Info("Finding nearest drivers")
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	results, err := c.client.FindNearest(thrift.Wrap(ctx), location)
	if err != nil {
		return nil, err
	}
	return fromThrift(results), nil
}

func fromThrift(results []*driver.DriverLocation) []Driver {
	retMe := make([]Driver, len(results))
	for i, result := range results {
		retMe[i] = Driver{
			DriverID: result.DriverID,
			Location: result.Location,
		}
	}
	return retMe
}

// Lock places a lock on the ID and is used to alter balances
func (c *Client) Lock(ctx context.Context, id string) {
	log.WithField("id", id).Info("Securing lock")
	ctx, cancel := context.WithTimeout(ctx, 100*time.Minute)
	defer cancel()
	_, err := c.client.Lock(thrift.Wrap(ctx), id)
	if err != nil {
		log.Error(err.Error())
	}
}

func (c *Client) Unlock(ctx context.Context, id string) {
	log.WithField("id", id).Info("releasing lock")
	ctx, cancel := context.WithTimeout(ctx, 100*time.Minute)
	defer cancel()
	_, err := c.client.Unlock(thrift.Wrap(ctx), id)
	if err != nil {
		log.Error(err.Error())
	}
}