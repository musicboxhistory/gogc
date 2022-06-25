printf "\nTEST 1\n"
curl -X GET -vL localhost:8081/n5g-eir-eic/v1/equipment-status?pei=11223344
printf "\n\nTEST 2\n"
curl -X GET -vL localhost:8081/n5g-eir-eic/v1/equipment-status?pei=55667788
printf "\n\nTEST 3\n"
curl -X GET -vL localhost:8081/n5g-eir-eic/v1/equipment-status?pei=12345678
printf "\n\nTEST 4\n"
curl -X GET -vL localhost:8081/n5g-eir-eic/v1/equipment-status
printf "\n"
