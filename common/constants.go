package common

import "os"


const BatchSize = 25
const BlockDelay = 5
const ExecutionClientRPC = "http://35.244.25.153/jsonrpc"
const SettlementClientRPC = "http://192.168.1.105:8080"
const KeyringDirectory = "./account/keys"
var DaClientRPC = os.Getenv("DA_CLIENT_RPC")
// const DaClientRPC = "http://localhost:5050/celestia"
