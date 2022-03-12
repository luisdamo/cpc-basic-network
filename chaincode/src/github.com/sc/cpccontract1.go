package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CPCContract1 struct {
}

type CPC struct {
	DMC   string `json:"DMC"`   // DMC: Código Datamatrix identificativo de la pieza
	TYPE  string `json:"TYPE"`  // TYPE: Identificador del tipo de pieza
	ST    int    `json:"ST"`    // ST: Valor numérico que representa el estado actual de la pieza (integer)
	IDMAN string `json:"IDMAN"` // IDMAN: Identificador del fabricante (string)
	IDASS string `json:"IDASS"` // IDASS: Identificador del ensamblador (string)
	IDCUS string `json:"IDCUS"` // IDCUS: Identificador del cliente (string)'
}

func (p *CPCContract1) Init(stub shim.ChaincodeStubInterface) peer.Response {
	msg := "todo cor'recto en la` inicializacion"
	bytearray := []byte(msg)
	return shim.Success(bytearray)
}

func (p *CPCContract1) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("Nombre function:", fn, "Parametros:", args)
	switch fn {
	case "get":
		return functionGet(stub, args)
	case "set":
		return functionSet(stub, args)
	case "initledger":
		return initLedger(stub)
	case "leerpieza":
		return leerPieza(stub, args)
	case "crearpieza":
		return crearPieza(stub, args)
	case "version":
		nversion := "V 1.2"
		return shim.Success([]byte(nversion))
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
	if len(args) != 2 {
		return shim.Error("Numero de argumentos incorrecto. Se espera 2")
	}

	key := args[0]
	value := args[1]

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
		CPC{DMC: "210312CA000001", TYPE: "CARTER001", ST: 1, IDMAN: "Fundiciones A SA", IDASS: "Ensamblados A SA", IDCUS: "Ford"},
		CPC{DMC: "210312CA000002", TYPE: "CARTER001", ST: 1, IDMAN: "Fundiciones A SA", IDASS: "Ensamblados A SA", IDCUS: "Toyota"},
		CPC{DMC: "210312OP000003", TYPE: "OILP001", ST: 1, IDMAN: "Fundiciones A SA", IDASS: "Ensamblados A SA", IDCUS: "Audi"},
		CPC{DMC: "210312OP000004", TYPE: "OILP001", ST: 1, IDMAN: "Fundiciones A SA", IDASS: "Ensamblados A SA", IDCUS: "Ford"},
		CPC{DMC: "210312GE000005", TYPE: "GEAR001", ST: 1, IDMAN: "Fundiciones A SA", IDASS: "Ensamblados A SA", IDCUS: "Ford"},
		CPC{DMC: "210312GE000006", TYPE: "GEAR005", ST: 1, IDMAN: "Fundiciones A SA", IDASS: "Ensamblados A SA", IDCUS: "Audi"},
	}

	i := 0
	for i < len(piezas) {
		fmt.Println("i es ", i)
		piezaAsBytes, _ := json.Marshal(piezas[i])
		stub.PutState(piezas[i].DMC, piezaAsBytes)
		fmt.Println("Añadida pieza ", piezas[i].DMC)
		i = i + 1
	}

	return shim.Success(nil)
}
func leerPieza(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Numero de argumentos incorrecto. Se espera 1")
	}

	piezaAsBytes, _ := stub.GetState(args[0])
	return shim.Success(piezaAsBytes)
}
func crearPieza(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 6 {
		return shim.Error("Numero incorrecto de argumentos. Se esperaban 6")
	}
	intstate, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Error de conversión de tipos, se esperaba entero")
	}
	var pieza = CPC{DMC: args[0], TYPE: args[1], ST: intstate, IDMAN: args[3], IDASS: args[4], IDCUS: args[5]}

	piezaAsBytes, _ := json.Marshal(pieza)
	stub.PutState(args[0], piezaAsBytes)

	return shim.Success(nil)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(CPCContract1)); err != nil {
		fmt.Printf("Error starting CPCContract1 chaincode: %s", err)
	}
}
