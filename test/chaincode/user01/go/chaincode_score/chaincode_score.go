package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ScoreChaincode --
type ScoreChaincode struct {
}

// Profile --
type Score struct {
	UserID     string    `json:"user_id"`
	GraduScore []Subject `json: "gradu_score"`
	AverScore  string    `json: "aver_score"`
}

// Subject --
type Subject struct {
	NameSubject  string `json:"name_subject"`
	ScoreSubject string `json:"score_subject"`
}

// Init ProfileChaincode
func (t *ScoreChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("user score Init")
	return shim.Success(nil)
}

func (t *ScoreChaincode) initScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var finalScore float64
	//   0          1
	// "userID", "class"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userID := args[0]
	scoreSubject := args[1]

	// ==== Check if user already exists ====
	profileAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if profileAsBytes != nil {
		fmt.Println("This profile already exists: " + userID)
		return shim.Error("This profile already exists: " + userID)
	}

	var listSubjectNew []Subject

	listSubject := strings.Split(scoreSubject, "&")

	for _, value := range listSubject {
		valueNew := strings.Split(value, "#")
		listSubjectNew = append(listSubjectNew, Subject{valueNew[0], valueNew[1]})
		score, _ := strconv.ParseFloat(valueNew[1], 10)
		finalScore = finalScore + score
	}

	finalScore = finalScore / float64(len(listSubject))

	score := Score{userID, listSubjectNew, strconv.FormatFloat(finalScore, 'f', 2, 64)}

	scoreJSONasBytes, err := json.Marshal(score)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save score to state ===
	err = stub.PutState(userID, scoreJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *ScoreChaincode) updateScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var finalScore float64
	//   0         1
	// "userID", "class"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userID := args[0]
	scoreSubject := args[1]

	fmt.Println("- start updateProfile ", userID)

	scoreAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("Failed to get user:" + err.Error())
	} else if scoreAsBytes == nil {
		return shim.Error("User does not exist")
	}

	var listSubjectNew []Subject

	listSubject := strings.Split(scoreSubject, "&")

	for _, value := range listSubject {
		valueNew := strings.Split(value, "#")
		listSubjectNew = append(listSubjectNew, Subject{valueNew[0], valueNew[1]})
		score, _ := strconv.ParseFloat(valueNew[1], 10)
		finalScore = finalScore + score
	}

	finalScore = finalScore / float64(len(listSubject))

	score := Score{userID, listSubjectNew, strconv.FormatFloat(finalScore, 'f', 2, 64)}

	scoreJSONasBytes, _ := json.Marshal(score)
	err = stub.PutState(userID, scoreJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end success")
	return shim.Success(nil)
}

func (t *ScoreChaincode) deleteScore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var scoreJSON Score
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]

	valAsbytes, err := stub.GetState(userID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + userID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"File score does not exist: " + userID + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &scoreJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + userID + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(userID)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

func (t *ScoreChaincode) getScoreByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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

func (t *ScoreChaincode) getValueScoreByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"user_id\":\"%s\"}}", userID)

	queryResults, err := getValueQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getValueQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getValueQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		buffer.WriteString(string(queryResponse.Value))
	}
	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
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

// Invoke --
func (t *ScoreChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("user score Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "getScoreByID" {
		// get profile by id
		return t.getScoreByID(stub, args)
	} else if function == "deleteScore" {
		// Delete profile
		return t.deleteScore(stub, args)
	} else if function == "updateScore" {
		// update profile
		return t.updateScore(stub, args)
	} else if function == "initScore" {
		// create new profile
		return t.initScore(stub, args)
	} else if function == "getValueScoreByID" {
		// create new profile
		return t.getValueScoreByID(stub, args)
	}

	return shim.Error("Invalid invoke function name")
}

func main() {
	err := shim.Start(new(ScoreChaincode))
	if err != nil {
		fmt.Printf("Error starting profile chaincode: %s", err)
	}
}
