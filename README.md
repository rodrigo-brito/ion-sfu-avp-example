# Simple example of usage

- SFU + AVP

- SFU: https://github.com/pion/ion-sfu
- AVP: https://github.com/pion/ion-avp

## How to start sever

- Start SFU GRPC+Websocket: `cd ion-sfu && go run cmd/signal/allrpc/main.go -c config.toml -jaddr ":7000" -gaddr ":50051"`
- Start AVP Server: `cd ion-avp && go run examples/save-to-webm/server/main.go -c examples/save-to-webm/server/config.toml`

## How to start Client

- `sudo npm i -g serve`
- `serve ion-sfu/examples/echotest`

## Recorder

SFU display your connection ID on terminal:

```
[2021-01-06 15:22:14.932] [DEBUG] [68][webrtctransport.go][func1] => Peer ckjlr0y8g0001ml85667bhl1f got remote track id: 92e29ac3-2ad9-4682-be38-71524df3f84c mediaSSRC: 1432367151 rid : streamID: eCKAxQ1Ga3hEXGjNWxqmCCyz1jrKWv76CYxa
```

`test` - ID of client connection, defined on echotest project.

Now, to recorder a stream: `cd ion-avp && go run examples/save-to-webm/client/main.go "test"`

Place your stream ID in stdout
