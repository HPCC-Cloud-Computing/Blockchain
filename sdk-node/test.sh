rm latency.txt
rm throughput.txt

LOOP = $1
node invoke.js -u user97 --channel mychannel --chaincode mycc1 -m initResult  -a "112" -a "2222" -a "Nam" -a "3333" -a "toan" -a "4444" -a "hung" -a "Gioi" -a "20172" -l "$LOOP"