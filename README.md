# Repositorio cpc-basic-network - inicialización
echo "# hyperledger" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/luisdamo/cpc-basic-network.git
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
peer chaincode install -n programa -p github.com -v 1.0

### Instanciamos el chaincode
peer chaincode instantiate -C main -n programa -v 1.0 -c '{"Args":[""]}'

### Invocamos el metodo set
peer chaincode invoke -C main -n programa -c '{"Args":["set", "id_1", "valor_1"]}'

### Invocamos el metodo get
peer chaincode query -C main -n programa -c '{"Args":["get", "id_1"]}'