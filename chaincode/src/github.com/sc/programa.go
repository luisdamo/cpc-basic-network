package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ProgramaContract struct {
}

func (p *ProgramaContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	msg := "todo cor'recto en la` inicializacion"
	bytearray := []byte(msg)
	return shim.Success(bytearray)
}

func (p *ProgramaContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("Nombre function:", fn, "Parametros:", args)
	switch fn {
	case "get":
		return functionGet(stub, args)
	case "set":
		return functionSet(stub, args)
	default:
		return shim.Error("la funcion solicitada no existe")
	}
	return shim.Success(nil)
}

func functionGet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var key string
	// leemos el valor del usuario
	key = args[0]
	// pregunta: que pasa si args es nulo?
	// pregunta: que pasa si args, existe pero esta vacio?
	// tarea: modifica el codigo para devolver error en los casos anteriores
	resp, err := stub.GetState(key)
	if err != nil {
		// codigo asociado a la gestion del error
		return shim.Error(err.Error())
	}
	// codigo asociado a la respuesta correcta
	return shim.Success(resp)
}

func functionSet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// leemos los valores del usuario
	key := args[0]
	value := args[1]
	// pregunta: que pasa si args es nulo?
	// pregunta: que pasa si args, existe pero esta vacio?
	// pregunta: que pasa si args no tiene los 2 elementos necesarios?
	// tarea: modifica el codigo para devolver error en los casos anteriores
	err := stub.PutState(key, []byte(value))
	if err != nil {
		// codigo asociado a la gestion del error
		return shim.Error(err.Error())
	}
	// codigo asociado a la respuesta correcta
	return shim.Success(nil)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(ProgramaContract)); err != nil {
		fmt.Printf("Error starting ProgramaContract chaincode: %s", err)
	}
}
