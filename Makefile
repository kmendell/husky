BINARY_NAME=testing/husky
 
build:
    go build -o ${BINARY_NAME}
 
run:
    go build -o ${BINARY_NAME}
    ${BINARY_NAME} testing/test.husky
 
clean:
    go clean
    rm ${BINARY_NAME}