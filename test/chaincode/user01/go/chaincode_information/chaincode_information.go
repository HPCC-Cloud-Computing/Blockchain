package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// InformationChaincode example simple Chaincode implementation
type InformationChaincode struct {
}

type User struct {
	UserID      string `json:"user_id"`
	NameUser    string `json:"name_user"`
	DateOfBrith string `json:"date_of_brith"`
	SexUser     string `json:"sex_user"`
	AddressUser string `json:"address_user"`
}

func (t *InformationChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("user information init")
	return shim.Success(nil)
}

func (t *InformationChaincode) initUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0           1       		2           3			 4
	// "userID", "nameUser", "dateOfBrith", "sexUser", "addressUser""
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	userID := args[0]
	nameUser := strings.ToLower(args[1])
	dateOfBrith := args[2]
	sexUser := strings.ToLower(args[3])
	addressUser := strings.ToLower(args[4])

	// ==== Check if user already exists ====
	userAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if userAsBytes != nil {
		fmt.Println("This user already exists: " + userID)
		return shim.Error("This user already exists: " + userID)
	}

	// ==== Create user object and marshal to JSON ====
	user := &User{userID, nameUser, dateOfBrith, sexUser, addressUser}
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save user to state ===
	err = stub.PutState(userID, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *InformationChaincode) updateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0         1			 	2			3			 4
	// "userID", "nameUser", "dateOfBrith", "sexUser", "addressUser"
	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userID := args[0]
	nameUser := strings.ToLower(args[1])
	dateOfBrith := args[2]
	sexUser := strings.ToLower(args[3])
	addressUser := strings.ToLower(args[4])

	fmt.Println("- start updateUser ", userID)

	userAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("Failed to get user:" + err.Error())
	} else if userAsBytes == nil {
		return shim.Error("User does not exist")
	}

	userOld := User{}
	err = json.Unmarshal(userAsBytes, &userOld) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	userOld.NameUser = nameUser
	userOld.DateOfBrith = dateOfBrith
	userOld.SexUser = sexUser
	userOld.AddressUser = addressUser

	userJSONasBytes, _ := json.Marshal(userOld)
	err = stub.PutState(userID, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end success")
	return shim.Success(nil)
}

func (t *InformationChaincode) getUserByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"user_id\":\"%s\"}}", userID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *InformationChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var userJSON User
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]

	// to maintain the color~name index, we need to read the marble first and get its color
	valAsbytes, err := stub.GetState(userID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + userID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"User does not exist: " + userID + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &userJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + userID + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(userID) //remove the marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

func (t *InformationChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("user_information Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "getUserByID" {
		// get user by id
		return t.getUserByID(stub, args)
	} else if function == "deleteUser" {
		// Delete user
		return t.deleteUser(stub, args)
	} else if function == "updateUser" {
		// update user
		return t.updateUser(stub, args)
	} else if function == "initUser" {
		// create new user
		fmt.Println("initUser")
		return t.initUser(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

func main() {
	err := shim.Start(new(InformationChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
