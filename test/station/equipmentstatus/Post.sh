printf "\nTEST 1\n"
curl -X POST -H "Content-Type: application/json" -d '[{"key":"imei-11223344", "status":"WHITELISTED"}, {"key":"imsi-55667788", "status":"BLACKLISTED"}]' -vL localhost:8888/station/v1/equipmentstatus/5geir
printf "\n\n"
