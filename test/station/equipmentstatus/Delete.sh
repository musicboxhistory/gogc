printf "\nTEST 1\n"
curl -X DELETE -vL localhost:8888/station/v1/equipmentstatus/5geir/pei-11223344
printf "\n\nTEST 2\n"
curl -X DELETE -vL localhost:8888/station/v1/equipmentstatus/5geir/pei-55667788
printf "\n\nTest 3\n"
curl -X DELETE -vL localhost:8888/station/v1/equipmentstatus/5geir/pei-12345678
printf "\n\n"
