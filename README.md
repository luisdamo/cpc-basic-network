# Repositorio cpc-basic-network - inicialización
echo "# hyperledger" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add cpc-basic-network https://github.com/luisdamo/cpc-basic-network.git
git push -u origin main
# Repositorio hyperledger - actualización
git add .
git commit -m "actualiz. repositorio"
git push -u origin main

### Descargar fabric-samples
curl -sSL https://bit.ly/2ysbOFE | bash -s -- 1.4.11 1.4.9

### Crear organizaciones de ejemplo con el template de configuración
./fabric-samples/bin/cryptogen generate --config=./crypto-config.yaml

### Creamos el bloque genesis con configtxgen
> El archivo de configuracion de ejemplo se encuentra en ./fabric-samples/basic-network/configtx.yaml

> El archivo de configuracion que hemos modificado lo encuentra automaticamente en la ruta desde la que ejecutamos el comando

mkdir config

./fabric-samples/bin/configtxgen -profile OneOrgOrdererGenesis -outputBlock ./config/genesis.block

### Creamos transaccion de configuracion del canal main
./fabric-samples/bin/configtxgen -profile OneOrgChannel -outputCreateChannelTx ./config/channel.tx -channelID main

### Creamos la transaccion anchor peer de cada organizacion
./fabric-samples/bin/configtxgen -profile OneOrgChannel -outputAnchorPeersUpdate ./config/Org1MSPanchors.tx -channelID main -asOrg Org1MSP

### Renombrar archivos de claves a key
./update_keynames.sh

### Levantamos el contenedor orderer
En la linea 64 de docker-compose.yml introducir la red de docker manualmente, ver con
docker network ls
Introducir  - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=basic-network_basic
docker-compose up -d

### Entramos a la terminal del nodo cliente
docker exec -it cli bash

### Comprobamos el estado del nodo peer
peer node status

### Leemos la lista de canales disponibles
peer channel list

### Creamos el canal main
peer channel create -c main -o orderer.example.com:7050 -f /etc/hyperledger/configtx/channel.tx

### Nos unimos al canal main
> En el paso anterior se nos ha creado el primer bloque, "main.block", del canal

peer channel join -b main.block

### Leemos las listas de chaincodes instalados e instanciados
peer chaincode list --installed

peer chaincode list -C main --instantiated

### Instalamos el chaincode
peer chaincode install -n cpccontract1 -p github.com -v 1.6

### Instanciamos el chaincode
peer chaincode instantiate -C main -n cpccontract1 -v 1.6 -c '{"Args":[""]}'
### Verificar el log
docker logs orderer.example.com
docker logs peer0.org1.example.com
### Invocamos el metodo set
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["set", "id_1", "valor_1"]}'
### Invocamos el metodo get
peer chaincode query -C main -n cpccontract1 -c '{"Args":["get", "id_1"]}'
### Invocamos el metodo initledger
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["initledger"]}'
### Invocamos el metodo crearpieza
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["crearpieza", "210312CA000007","CARTER001","1","Fundiciones B SA","Ensamblados A SA","Seat"]}'
### Invocamos el metodo leerpieza
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["leerpieza", "210312CA000007"]}'
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["leerpieza", "210312CA000001"]}'
### Invocamos el metodo leerpiezas
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["leerpiezas"]}'
### Invocamos el metodo actualizarestadopieza
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["actualizarestadopieza","210312CA000001","2"]}'
### Invocamos el metodo version
Informa de la vesión del contrato
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["version"]}'

## Codigo del contrato
### Codigo del contrato: Estructura de datos
Utilizamos como punto de partida el contrato programa utilizado en las practicas de clase  de fabric-samples para escribir
el contrato cpccontract (car parts contract), representado por la clase Cpc.
Definimos el array assets para contener la base de datos de activos utilizados por el contrato
En nuestro caso, representamos piezas de aluminio para automóvil que representaremos por la siguiente estructura:
- DMC: Código Datamatrix identificativo de la pieza
- TYPE: Identificador del tipo de pieza
- ST: Valor numérico que representa el estado actual de la pieza (integer)
- IDMAN: Identificador del fabricante (string)
- IDASS: Identificador del ensamblador (string)
- IDCUS: Identificador del cliente (string)'
### Codigo del contrato: función initledger
Genera 6 piezas para pruebas
Inicializa el ledger con piezas para pruebas
### Invocar función initledger
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["initledger"]}'
### Invocamos el metodo leerpieza
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["leerpieza", "210312CA000002"]}'
### Codigo del contrato: función crearpieza
Genera una pieza con la estructura CPC con los argumentos pasados como parámetros
pieza = CPC{DMC: args[0], TYPE: args[1], ST: intstate, IDMAN: args[3], IDASS: args[4], IDCUS: args[5]}
### Invocar función crearpieza
peer chaincode invoke -C main -n cpccontract1 -c '{"Args":["crearpieza", "210312CA000007","CARTER001","1","Fundiciones B SA","Ensamblados A SA","Seat"]}'

## Compilación del programa del contrato
En la carpeta: cpc-basic-network/chaincode/src/github.com/sc
Ejecutamos el comando: 
go build cpccontract1.go

### Instalación de node-red
npm install -g --unsafe-perm node-red@1.2
Una vez instalado, podemos acceder a la interfase vía web:
http://localhost:1880
Instalar componentes para fabric
npm i node-red-contrib-fabric