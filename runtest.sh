executable=./bin/huffman-compressor

for test in `ls ./test`
do
  $executable -src ./test/$test -dst ./test/$test.huff
  $executable -src ./test/$test.huff -dst ./test/${test%%.*}.ext.${test#*.} -ext
  if diff ./test/$test ./test/${test%%.*}.ext.${test#*.}
  then
    echo "Test case "${test}" passed"
  else
    echo "Test case "${test}" failed"
  fi
  rm -rf ./test/$test.huff
  rm -rf ./test/${test%%.*}.ext.${test#*.}
done
