go test ./... -cover \
    | grep coverage \
    | awk '{print substr($5, 1, length($5) - 1) " " $2 }' \
    | awk -v THRESHOLD=60 '{ \
    if (!($1 >= THRESHOLD)) { \
        print " Coverage:\t" $2 "\t" $1 "% \033[31m FAIL\033[0m" ", Expected threshold: " THRESHOLD "% ";\
        exit 1\
    } else { \
        print " Coverage:\t" $2 "\t" $1 "% \033[32m OK \033[0m" \
    } }'