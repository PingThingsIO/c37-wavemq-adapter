package main

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/PingThingsIO/c37wavemq/c37pb"
	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/proto"
	"github.com/immesys/wavemq/mqpb"
	logging "github.com/op/go-logging"
	"google.golang.org/grpc"
)

const EntityFile = "c37.ent"
const Namespace = "GyCetklhSNcgsCKVKXxSuCUZP4M80z9NRxU1pwfb2XwGhg=="
const SiteRouter = "127.0.0.1:4516"
const URI = "upmu/*"

var lg *logging.Logger

func init() {
	logging.SetBackend(logging.NewLogBackend(os.Stdout, "", 0))
	logging.SetFormatter(logging.MustStringFormatter("[%{level:-8s}]%{time:2006-01-02T15:04:05.000000} %{shortfile:18s} > %{message}"))
	lg = logging.MustGetLogger("log")
}

func main() {
	//Initiate waveMQ connection
	entityFile, err := ioutil.ReadFile(EntityFile)
	if err != nil {
		lg.Fatalf("could not read entity file: %v", err)
	}

	namespaceBytes, err := base64.URLEncoding.DecodeString(Namespace)
	if err != nil {
		lg.Fatalf("failed to decode namespace: %v", err)
	}

	// Establish a GRPC connection to the site router.
	conn, err := grpc.Dial(SiteRouter, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		lg.Fatalf("could not connect to the site router: %v", err)
	}

	// Create the WAVEMQ client
	client := mqpb.NewWAVEMQClient(conn)

	sub, err := client.Subscribe(context.Background(), &mqpb.SubscribeParams{
		Perspective: &mqpb.Perspective{
			EntitySecret: &mqpb.EntitySecret{
				DER: entityFile,
			},
		},
		Namespace: namespaceBytes,
		Uri:       URI,
		Expiry:    120,
	})
	if err != nil {
		lg.Fatalf("could not subscribe: grpc error: %v", err)
	}
	for {
		m, err := sub.Recv()
		if err != nil {
			lg.Fatalf("could not subscribe: grpc error: %v", err)
		}
		if m.Error != nil {
			lg.Fatalf("subscriptio nerror: %v", m.Error.Message)
		}
		frame := m.Message.Tbs.Payload[0]
		c37 := c37pb.C37DataFrame{}
		err = proto.Unmarshal(frame.Content, &c37)
		if err != nil {
			lg.Fatalf("could not unmarshal c37 protobuf: %v", err)
		}
		spew.Dump(c37)
	}
}
