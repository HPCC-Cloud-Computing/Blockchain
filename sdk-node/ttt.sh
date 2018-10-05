test=$1
LOOP=$2
for i in `seq 1 $LOOP`; do
test="$test$i"
echo $test
done