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

Ensure, that logs look similar to the example below.

For server:
```
2022/11/10 15:45:08 server listening at [::]:13111
2022/11/10 15:45:18 client id 333 connected
2022/11/10 15:45:18 client id 111 connected
2022/11/10 15:45:18 client id 222 connected
2022/11/10 15:45:23 failed to register listener: RegisterListener: client is already present in group
2022/11/10 15:45:23 sendGroupMessages: group 222 not found, skip to users: GetListeners: listeners not found
2022/11/10 15:45:23 sendUserMessage: user chat10 not found: SendMessage: client's data was not found
```

For client:
```
2022/11/10 15:45:23 CreateGroupChat: in chat_name:"chat10" username:"111", out , err <nil>
2022/11/10 15:45:23 CreateGroupChat: in chat_name:"chat10" username:"222", out , err <nil>
2022/11/10 15:45:23 CreateGroupChat: in chat_name:"chat10" username:"222", out <nil>, err rpc error: code = Internal desc = internal error
2022/11/10 15:45:23 CreateGroupChat: in chat_name:"chat10" username:"333", out , err <nil>
2022/11/10 15:45:23 LeftGroupChat: in chat_name:"chat10" username:"333", out , err <nil>
2022/11/10 15:45:23 SendMessage: in data:{sender_id:"111" receiver_id:"222" payload:"hello from 111 to 222"}, out , err <nil>
2022/11/10 15:45:23 Message received from stream 222: sender_id:"111" receiver_id:"222" payload:"hello from 111 to 222"
2022/11/10 15:45:23 Message received from stream 222: sender_id:"chat10" receiver_id:"222" payload:"hello from 111 to chan10"
2022/11/10 15:45:23 SendMessage: in data:{sender_id:"111" receiver_id:"chat10" payload:"hello from 111 to chan10"}, out , err <nil>
2022/11/10 15:45:23 Message received from stream 111: sender_id:"chat10" receiver_id:"111" payload:"hello from 111 to chan10"
2022/11/10 15:45:23 SendMessage: in , out groups:{group_id:"chat10" usernames:"111" usernames:"222"}, err <nil>
2022/11/10 15:45:23 finished
```

To re-generate protobuf schema type command
```
make generate  
```