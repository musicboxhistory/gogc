find . -name "*.go" | xargs grep -rl "openapi" | xargs gsed -i -e "s/openapi/model/g"


amf/Namf_Communication
amf/Namf_MBSBroadcast
lmf/Nlmf_Location
smf/Nsmf_PDUSession

curl -s localhost:8081/n5g-eir-eic/v1/equipment-status

openapi-generator generate -i openapi.yml -g go-gin-server -o ./server


brew services start mongodb-community
brew services stop mongodb-community

brew services start redis
brew services stop redis
