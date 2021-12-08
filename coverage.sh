go test ./... -cover | grep coverage|awk '{print substr($5, 1, length($5) - 1)}' | awk '{if (!($1 >= 90)) { print "Coverage: " $1 "%" ", Expected threshold: " 90 "%"; exit 1 } }'

