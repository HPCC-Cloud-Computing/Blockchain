package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// MainChaincode --
type MainChaincode struct {
}

type User struct {
	UserID      string `json:"user_id"`
	NameUser    string `json:"name_user"`
	DateOfBrith string `json:"date_of_brith"`
	SexUser     string `json:"sex_user"`
	AddressUser string `json:"address_user"`
}

// Profile --
type Profile struct {
	UserID  string   `json:"user_id"`
	Class10 Class    `json:"class_10"`
	Class11 Class    `json:"class_11"`
	Class12 Class    `json:"class_12"`
	BC      []string `json:"bc"`
}

// Class --
type Class struct {
	ClassName  string    `json:"class_name"`
	NameSchool string    `json:"name_school"`
	SchoolYear string    `json:"school_year"`
	NameHT     string    `json:"name_HT"`
	NameGVCN   string    `json:"name_GVCN"`
	Subjects   []Subject `ison:"subjects"`
	FinalScore string    `json:"final_score"`
	HK         string    `json:"hk"`
	DH         []string  `json:"dh"`
}

type ResponseCheck struct {
	UserID      string `json:"user_id"`
	NameUser    string `json:"name_user"`
	DateOfBrith string `json:"date_of_brith"`
	SexUser     string `json:"sex_user"`
	AddressUser string `json:"address_user"`
	BCString    string `json:"bc_string"`
}

type BCUser struct {
	BCString string `json:"bc_string"`
}

// Subject --
type Subject struct {
	NameSubject  string `json:"name_subject"`
	ScoreSubject string `json:"score_subject"`
}

type Score struct {
	UserID     string    `json:"user_id"`
	GraduScore []Subject `json: "gradu_score"`
	AverScore  string    `json: "aver_score"`
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

	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 4")
	}

	chaincodeName := args[0]

	channelName = ""

	// Query new chaincode
	functionName := "updateProfile"
	userID := args[1]
	class := args[2]
	level := args[3]

	queryArgs = toChaincodeArgs(functionName, userID, class, level)

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

func (t *MainChaincode) initScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 3")
	}

	chaincodeName := args[0]
	userID := args[1]
	score := args[2]

	channelName = ""

	// Query new chaincode
	queryArgs := toChaincodeArgs("initScore", userID, score)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(nil)
}

func (t *MainChaincode) updateScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 3")
	}

	chaincodeName := args[0]
	userID := args[1]
	score := args[2]

	channelName = ""

	// Query new chaincode
	queryArgs := toChaincodeArgs("updateScore", userID, score)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(nil)
}

func (t *MainChaincode) deleteScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0]
	userID := args[1]

	channelName = ""

	// Query new chaincode
	queryArgs := toChaincodeArgs("deleteScore", userID)

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	return shim.Success(nil)
}

func (t *MainChaincode) getScoreByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0]
	userID := args[1]

	channelName = ""

	// Query new chaincode
	queryArgs := toChaincodeArgs("getScoreByID", userID)

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
	var err error

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 3")
	}

	chaincodeNameUser := args[0]
	chaincodeNameProfile := args[1]
	chaincodeNameScore := args[2]
	userID := args[3]

	channelName = ""

	// Query new chaincode
	queryArgsScore := toChaincodeArgs("getValueScoreByID", userID)

	responseA := stub.InvokeChaincode(chaincodeNameScore, queryArgsScore, channelName)
	if responseA.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", responseA.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	score := &Score{}

	err = json.Unmarshal(responseA.Payload, &score)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryArgsProfile := toChaincodeArgs("checkScore", userID, score.AverScore)

	responseB := stub.InvokeChaincode(chaincodeNameProfile, queryArgsProfile, channelName)
	if responseB.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", responseB.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	bcString := &BCUser{}

	err = json.Unmarshal(responseB.Payload, bcString)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to decode JSONaaa of: " + userID + "\"}"
		return shim.Error(jsonResp)
	}

	queryArgsUser := toChaincodeArgs("getValueUserByID", userID)

	responseC := stub.InvokeChaincode(chaincodeNameUser, queryArgsUser, channelName)
	if responseC.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode profile. Got error: %s", responseC.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	user := &User{}

	err = json.Unmarshal(responseC.Payload, &user)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to decode JSON of: " + userID + "\"}"
		return shim.Error(jsonResp)
	}

	checkScore := ResponseCheck{user.UserID, user.NameUser, user.DateOfBrith, user.SexUser, user.AddressUser, bcString.BCString}

	dataJSONasBytes, err := json.Marshal(checkScore)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(dataJSONasBytes)
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
	} else if function == "initScore" {
		return t.initScore(stub, args)
	} else if function == "updateScore" {
		return t.updateScore(stub, args)
	} else if function == "getScoreByID" {
		return t.getScoreByID(stub, args)
	} else if function == "deleteScore" {
		return t.deleteScore(stub, args)
	}
	return shim.Success([]byte("Invalid invoke function name"))
}

func main() {

	err := shim.Start(new(MainChaincode))
	if err != nil {
		fmt.Println("Error starting Simple chaincode: ", err)
	}
}
