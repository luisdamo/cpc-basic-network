package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CPCContract struct {
}

// Definimos la estructura CPC que representa la pieza
type CPC struct {
	Fabricante  string `json:"Fabricante"`
	TipoPieza   string `json:"TipoPieza"`
	Estado      int    `json:"Estado"`
	Ensamblador string `json:"Ensamblador"`
	Cliente     string `json:"Cliente"`
}

func (p *CPCContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	msg := "todo cor'recto en la` inicializacion"
	bytearray := []byte(msg)
	return shim.Success(bytearray)
}

func (p *CPCContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("Nombre function:", fn, "Parametros:", args)
	switch fn {
	case "get":
		return functionGet(stub, args)
	case "set":
		return functionSet(stub, args)
	case "initledger":
		return initLedger(stub)
	case "getpieza":
		return getPieza(stub, args)
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
func initLedger(stub shim.ChaincodeStubInterface) peer.Response {
	piezas := []CPC{
		CPC{Fabricante: "Fundiciones A SA", TipoPieza: "CARTER001", Estado: 1, Ensamblador: "Ensamblados A SA", Cliente: "Ford"},
		CPC{Fabricante: "Fundiciones A SA", TipoPieza: "CARTER024", Estado: 1, Ensamblador: "Ensamblados A SA", Cliente: "Toyota"},
		CPC{Fabricante: "Fundiciones A SA", TipoPieza: "OILP023", Estado: 1, Ensamblador: "Ensamblados A SA", Cliente: "Audi"},
	}

	i := 0
	for i < len(piezas) {
		fmt.Println("i es ", i)
		piezaAsBytes, _ := json.Marshal(piezas[i])
		stub.PutState("PIEZA"+strconv.Itoa(i), piezaAsBytes)
		fmt.Println("Añadida pieza ", piezas[i])
		i = i + 1
	}

	return shim.Success(nil)
}
func getPieza(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Numero de argumentos incorrecto. Se espera 1")
	}

	piezaAsBytes, _ := stub.GetState(args[0])
	return shim.Success(piezaAsBytes)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(CPCContract)); err != nil {
		fmt.Printf("Error starting CPCContract chaincode: %s", err)
	}
}
