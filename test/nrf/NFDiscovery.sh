printf "\nTEST 1\n"
curl -X GET -vL "localhost:8082/nnrf-disc/v1/nf-instances?target-nf-type=AMF&requester-nf-type=AMF"
printf "\n\nTEST 2\n"
curl -X GET -vL "localhost:8082/nnrf-disc/v1/nf-instances?target-nf-type=UPF&requester-nf-type=AMF"
printf "\n\nTEST 3\n"
curl -X GET -vL localhost:8082/nnrf-disc/v1/nf-instances
printf "\n"
