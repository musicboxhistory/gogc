printf "\nTEST 1\n"
curl -X PUT -vL "localhost:8481/nudr-dr/v1/subscription-data/imsi-11223344/context-data/amf-3gpp-access" -d @Amf3GppAccessRegistration.json -H "Content-Type: application/json"
printf "\n"
