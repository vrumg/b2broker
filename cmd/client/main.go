package main

import (
	"context"
	"io"
	"log"

	pb "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	// dial server
	conn, err := grpc.Dial(":13111", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create client
	client := pb.NewMessageServiceClient(conn)

	// set connection stream for users
	go clientByIDConnect(ctx, "111", client)
	go clientByIDConnect(ctx, "222", client)
	go clientByIDConnect(ctx, "333", client)

	inCreateGroupChat := &pb.CreateGroupChatRequest{
		ChatName: "chat10",
		Username: "111",
	}
	outCreateGroupChat, err := client.CreateGroupChat(ctx, inCreateGroupChat)
	log.Printf("CreateGroupChat: in %+v, out %+v", inCreateGroupChat, outCreateGroupChat)

	inJoinGroupChat := &pb.JoinGroupChatRequest{
		ChatName: "chat10",
		Username: "222",
	}
	outJoinGroupChat, err := client.JoinGroupChat(ctx, inJoinGroupChat)
	log.Printf("CreateGroupChat: in %+v, out %+v", inJoinGroupChat, outJoinGroupChat)

	inJoinGroupChat2 := &pb.JoinGroupChatRequest{
		ChatName: "chat10",
		Username: "222",
	}
	outJoinGroupChat2, err := client.JoinGroupChat(ctx, inJoinGroupChat2)
	log.Printf("CreateGroupChat: in %+v, out %+v", inJoinGroupChat2, outJoinGroupChat2)

	inJoinGroupChat3 := &pb.JoinGroupChatRequest{
		ChatName: "chat10",
		Username: "333",
	}
	outJoinGroupChat3, err := client.JoinGroupChat(ctx, inJoinGroupChat3)
	log.Printf("CreateGroupChat: in %+v, out %+v", inJoinGroupChat3, outJoinGroupChat3)

	inLeftGroupChat := &pb.LeftGroupChatRequest{
		ChatName: "chat10",
		Username: "333",
	}
	outLeftGroupChat, err := client.LeftGroupChat(ctx, inLeftGroupChat)
	log.Printf("LeftGroupChat: in %+v, out %+v", inLeftGroupChat, outLeftGroupChat)

	inSendMessage := &pb.SendMessageRequest{
		Data: &pb.Message{
			SenderId:   "111",
			ReceiverId: "222",
			Payload:    "hello from 111 to 222",
		},
	}
	outSendMessage, err := client.SendMessage(ctx, inSendMessage)
	log.Printf("SendMessage: in %+v, out %+v", inSendMessage, outSendMessage)

	inSendMessage2 := &pb.SendMessageRequest{
		Data: &pb.Message{
			SenderId:   "111",
			ReceiverId: "chan10",
			Payload:    "hello from 111 to 222",
		},
	}
	outSendMessage2, err := client.SendMessage(ctx, inSendMessage2)
	log.Printf("SendMessage: in %+v, out %+v", inSendMessage2, outSendMessage2)

	inListChannels := &pb.ListChannelsRequest{}
	outListChannels, err := client.ListChannels(ctx, inListChannels)
	log.Printf("SendMessage: in %+v, out %+v", inListChannels, outListChannels)

	log.Printf("finished")
}

func clientByIDConnect(ctx context.Context, clientID string, client pb.MessageServiceClient) {
	in := &pb.ConnectRequest{Username: "111"}
	stream, err := client.Connect(ctx, in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Message received from stream %s: %v", clientID, resp.Data)
		}
	}()

	<-done //we will wait until all response is received
}
