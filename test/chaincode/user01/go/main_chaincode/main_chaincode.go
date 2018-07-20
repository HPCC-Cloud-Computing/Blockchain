package main

import (
	"fmt"
	"strings"

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
		return shim.Error("Incorrect number of arguments. Expecting atleast 6")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := "initUser"
	userID := args[1]
	nameUser := args[2]
	dateOfBrith := args[3]
	sexUser := args[4]
	addressUser := args[5]

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

	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 4")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := "initProfile"
	userID := args[1]
	class := args[2]
	bc := args[3]

	queryArgsA := toChaincodeArgs("getUserByID", userID)

	responseA := stub.InvokeChaincode("aaa1", queryArgsA, channelName)
	if responseA.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", responseA.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	if responseA.Payload == nil {
		return shim.Error("User does not exist")
	}

	queryArgs = toChaincodeArgs(functionName, userID, class, bc)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(nil)
}

func (t *MainChaincode) updateProfile(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string
	var queryArgs [][]byte

	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 5")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := "updateProfile"
	userID := args[1]
	class := args[2]
	bc := args[3]
	level := args[4]

	queryArgs = toChaincodeArgs(functionName, userID, class, bc, level)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	return shim.Success(nil)
}

func (t *MainChaincode) updateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string
	var queryArgs [][]byte

	if len(args) < 6 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 6")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := "updateUser"
	userID := args[1]
	nameUser := strings.ToLower(args[2])
	dateOfBrith := args[3]
	sexUser := strings.ToLower(args[4])
	addressUser := strings.ToLower(args[5])

	queryArgs = toChaincodeArgs(functionName, userID, nameUser, dateOfBrith, sexUser, addressUser)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	return shim.Success(nil)
}

func (t *MainChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 3")
	}

	chaincodeNameUser := args[0]
	chaincodeNameProfile := args[1]

	channelName = ""

	// Query new chaincode
	userID := args[2]

	queryArgsUser := toChaincodeArgs("deleteUser", userID)

	responseA := stub.InvokeChaincode(chaincodeNameUser, queryArgsUser, channelName)
	if responseA.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode user. Got error: %s", responseA.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	queryArgsProfile := toChaincodeArgs("deleteProfile", userID)

	responseB := stub.InvokeChaincode(chaincodeNameProfile, queryArgsProfile, channelName)
	if responseB.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", responseB.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(nil)
}

func (t *MainChaincode) getUserByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	userID := args[1]

	queryArgs := toChaincodeArgs("getUserByID", userID)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode user. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(response.Payload)
}

func (t *MainChaincode) getProfileByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	userID := args[1]

	queryArgs := toChaincodeArgs("getProfileByID", userID)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode user. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(response.Payload)
}

func (t *MainChaincode) getListProfileOfClass(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 3")
	}

	chaincodeName := args[0]
	level := args[1]
	schoolYear := args[2]
	className := args[3]

	channelName = ""

	// Query new chaincode
	queryArgs := toChaincodeArgs("getListProfileOfClass", level, schoolYear, className)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(response.Payload)
}

func (t *MainChaincode) checkScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0]
	userID := args[1]

	channelName = ""

	// Query new chaincode
	queryArgs := toChaincodeArgs("checkScore", userID)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(response.Payload)
}

func (t *MainChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "initUser" {
		return t.initUser(stub, args)
	} else if function == "initProfile" {
		return t.initProfile(stub, args)
	} else if function == "updateUser" {
		return t.updateUser(stub, args)
	} else if function == "updateProfile" {
		return t.updateProfile(stub, args)
	} else if function == "deleteUser" {
		return t.deleteUser(stub, args)
	} else if function == "getUserByID" {
		return t.getUserByID(stub, args)
	} else if function == "getProfileByID" {
		return t.getProfileByID(stub, args)
	} else if function == "getListProfileOfClass" {
		return t.getListProfileOfClass(stub, args)
	} else if function == "checkScore" {
		return t.checkScore(stub, args)
	}
	return shim.Success([]byte("Invalid invoke function name"))
}

func main() {

	err := shim.Start(new(MainChaincode))
	if err != nil {
		fmt.Println("Error starting Simple chaincode: ", err)
	}
}
