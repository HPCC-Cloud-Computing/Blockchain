rm latency.txt
rm throughput.txt

ID1=$1
LOOP=$2
for i in 0..$LOOP; do
    node invoke.js -u user97 --channel mychannel --chaincode mycc1 -m initResult  -a "$ID1" -a "2222" -a "Nam" -a "3333" -a "toan" -a "4444" -a "hung" -a "Gioi" -a "20172"
    $ID1 = $ID1 + "0";
done
