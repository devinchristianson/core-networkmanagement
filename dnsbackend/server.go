// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dnsbackend

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"core-networkmanager/dnsbackend/pb"
	"github.com/miekg/dns"
	"google.golang.org/grpc"
)

type dnsServer struct{}

func (d *dnsServer) Query(ctx context.Context, in *pb.DnsPacket) (*pb.DnsPacket, error) {
	m := new(dns.Msg)
	if err := m.Unpack(in.Msg); err != nil {
		return nil, fmt.Errorf("failed to unpack msg: %v", err)
	}
	r := new(dns.Msg)
	r.SetReply(m)
	r.Authoritative = true
	
	//do stuff here

	if len(r.Answer) == 0 {
		r.Rcode = dns.RcodeNameError
	}

	out, err := r.Pack()
	if err != nil {
		return nil, fmt.Errorf("failed to pack msg: %v", err)
	}
	return &pb.DnsPacket{Msg: out}, nil
}

func main() {
	var port = 8053
	lis, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDnsServiceServer(grpcServer, &dnsServer{})
	panic(grpcServer.Serve(lis))
}