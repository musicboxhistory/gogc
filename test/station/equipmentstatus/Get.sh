printf "\nTEST 1\n"
curl -X GET -vL localhost:8888/station/v1/equipmentstatus/5geir/imei-11223344
printf "\n\nTEST 2\n"
curl -X GET -vL localhost:8888/station/v1/equipmentstatus/5geir
printf "\n\n"
