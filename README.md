<H1>Simple chatting service with groups.</H1>

Server and Client located under ./cmd

Start server with command
```
make run-server
```

Then open the second terminal and start client with command
```
make run-client
```

Check logs to look similar to the example below

For server:
```
2022/11/10 16:05:34 server listening at [::]:13111
2022/11/10 16:05:41 client id 111 connected
2022/11/10 16:05:41 client id 222 connected
2022/11/10 16:05:41 client id 333 connected
2022/11/10 16:05:46 failed to register listener: RegisterListener: client is already present in group
2022/11/10 16:05:46 sendGroupMessages: group 222 not found, skip to users: GetListeners: listeners not found
2022/11/10 16:05:46 sendUserMessage: user chat10 not found: SendMessage: client's data was not found
2022/11/10 16:05:46 UnregisterListenerFromAllGroups: client is not present in group chat10
```

For client:
```
2022/11/10 16:05:46 CreateGroupChat: in chat_name:"chat10" username:"111", out , err <nil>
2022/11/10 16:05:46 CreateGroupChat: in chat_name:"chat10" username:"222", out , err <nil>
2022/11/10 16:05:46 CreateGroupChat: in chat_name:"chat10" username:"222", out <nil>, err rpc error: code = Internal desc = internal error
2022/11/10 16:05:46 CreateGroupChat: in chat_name:"chat10" username:"333", out , err <nil>
2022/11/10 16:05:46 LeftGroupChat: in chat_name:"chat10" username:"333", out , err <nil>
2022/11/10 16:05:46 SendMessage: in data:{sender_id:"111" receiver_id:"222" payload:"hello from 111 to 222"}, out , err <nil>
2022/11/10 16:05:46 Message received from stream 222: sender_id:"111" receiver_id:"222" payload:"hello from 111 to 222"
2022/11/10 16:05:46 Message received from stream 222: sender_id:"chat10" receiver_id:"222" payload:"hello from 111 to chan10"
2022/11/10 16:05:46 SendMessage: in data:{sender_id:"111" receiver_id:"chat10" payload:"hello from 111 to chan10"}, out , err <nil>
2022/11/10 16:05:46 Message received from stream 111: sender_id:"chat10" receiver_id:"111" payload:"hello from 111 to chan10"
2022/11/10 16:05:46 SendMessage: in , out groups:{group_id:"chat10" usernames:"111" usernames:"222"}, err <nil>
2022/11/10 16:05:46 finished
```

To re-generate protobuf schema type command
```
make generate  
```