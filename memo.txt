find . -name "*.go" | xargs grep -rl "openapi" | xargs gsed -i -e "s/openapi/model/g"


amf/Namf_MBSBroadcast
lmf/Nlmf_Location
smf/Nsmf_PDUSession

openapi-generator generate -i openapi.yml -g go-gin-server -o ./server


brew services start mongodb-community
brew services stop mongodb-community

brew services start redis
brew services stop redis
