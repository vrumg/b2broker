package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(":13111", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := pb.NewMessageServiceClient(conn)

	go clientByIDConnect(ctx, "111", client)
	go clientByIDConnect(ctx, "222", client)
	go clientByIDConnect(ctx, "333", client)

	time.Sleep(time.Second * 5)

	inCreateGroupChat := &pb.CreateGroupChatRequest{
		ChatName: "chat10",
		Username: "111",
	}
	outCreateGroupChat, err := client.CreateGroupChat(ctx, inCreateGroupChat)
	log.Printf("CreateGroupChat: in %v, out %v, err %v", inCreateGroupChat, outCreateGroupChat, err)

	inJoinGroupChat := &pb.JoinGroupChatRequest{
		ChatName: "chat10",
		Username: "222",
	}
	outJoinGroupChat, err := client.JoinGroupChat(ctx, inJoinGroupChat)
	log.Printf("CreateGroupChat: in %v, out %v, err %v", inJoinGroupChat, outJoinGroupChat, err)

	inJoinGroupChat2 := &pb.JoinGroupChatRequest{
		ChatName: "chat10",
		Username: "222",
	}
	outJoinGroupChat2, err := client.JoinGroupChat(ctx, inJoinGroupChat2)
	log.Printf("CreateGroupChat: in %v, out %v, err %v", inJoinGroupChat2, outJoinGroupChat2, err)

	inJoinGroupChat3 := &pb.JoinGroupChatRequest{
		ChatName: "chat10",
		Username: "333",
	}
	outJoinGroupChat3, err := client.JoinGroupChat(ctx, inJoinGroupChat3)
	log.Printf("CreateGroupChat: in %v, out %v, err %v", inJoinGroupChat3, outJoinGroupChat3, err)

	inLeftGroupChat := &pb.LeftGroupChatRequest{
		ChatName: "chat10",
		Username: "333",
	}
	outLeftGroupChat, err := client.LeftGroupChat(ctx, inLeftGroupChat)
	log.Printf("LeftGroupChat: in %v, out %v, err %v", inLeftGroupChat, outLeftGroupChat, err)

	inSendMessage := &pb.SendMessageRequest{
		Data: &pb.Message{
			SenderId:   "111",
			ReceiverId: "222",
			Payload:    "hello from 111 to 222",
		},
	}
	outSendMessage, err := client.SendMessage(ctx, inSendMessage)
	log.Printf("SendMessage: in %v, out %v, err %v", inSendMessage, outSendMessage, err)

	inSendMessage2 := &pb.SendMessageRequest{
		Data: &pb.Message{
			SenderId:   "111",
			ReceiverId: "chat10",
			Payload:    "hello from 111 to chan10",
		},
	}
	outSendMessage2, err := client.SendMessage(ctx, inSendMessage2)
	log.Printf("SendMessage: in %v, out %v, err %v", inSendMessage2, outSendMessage2, err)

	inListChannels := &pb.ListChannelsRequest{}
	outListChannels, err := client.ListChannels(ctx, inListChannels)
	log.Printf("SendMessage: in %v, out %v, err %v", inListChannels, outListChannels, err)

	log.Printf("finished")
}

func clientByIDConnect(ctx context.Context, clientID string, client pb.MessageServiceClient) {
	in := &pb.ConnectRequest{Username: clientID}
	stream, err := client.Connect(ctx, in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Message received from stream %s: %v", clientID, resp.Data)
		}
	}()

	<-done
}
