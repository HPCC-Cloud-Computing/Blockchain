package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// MainChaincode --
type MainChaincode struct {
}

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

// Init MainChaincode --
func (t *MainChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("user Init")
	return shim.Success(nil)
}

// init User
func (t *MainChaincode) initUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string
	var queryArgs [][]byte

	if len(args) < 6 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 1")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := args[1]
	userID := args[2]
	nameUser := args[3]
	dateOfBrith := args[4]
	sexUser := args[5]
	addressUser := args[6]

	queryArgs = toChaincodeArgs(functionName, userID, nameUser, dateOfBrith, sexUser, addressUser)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	return shim.Success(nil)
}

func (t *MainChaincode) initProfile(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string
	var queryArgs [][]byte

	if len(args) < 6 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 1")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := args[1]
	userID := args[2]
	class := args[3]
	bc := args[4]

	queryArgs = toChaincodeArgs(functionName, userID, class, bc)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	return shim.Success(nil)
}

func (t *MainChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "initUser" {
		return t.initUser(stub, args)
	} else if function == "initProfile" {
		return t.initProfile(stub, args)

	return shim.Success([]byte("Invalid invoke function name. Expecting \"invoke\" \"query\""))
}

func main() {
	a := new(MainChaincode)
	stubA := shim.NewMockStub("ex05", a)
	a.Init(stubA)
	err := shim.Start(a)
	if err != nil {
		fmt.Println("Error starting Simple chaincode: ", err)
	}
	// agr := (string){"Args":["a","10"]}
	// stubA := shim.ChaincodeStubInterface()
	// a.Init(stubA)
}
