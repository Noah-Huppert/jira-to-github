.PHONY: install run \
	test test-out

ORIG_EXE_NAME=jira-to-github
EXE_NAME=j2gh
MAIN_FILE=main.go

TEST_OUT_DIR=test_out
TEST_COVER_FILE=coverage.out
TEST_COVER_PATH=${TEST_OUT_DIR}/${TEST_COVER_FILE}
TEST_COVER_MODE=count

# install: builds and saves executable in "$GOPATH/${EXE_NAME}"
install:
	go install
	mv "${GOPATH}/bin/${ORIG_EXE_NAME}" "${GOPATH}/bin/${EXE_NAME}"

# run builds and executes the tool
run:
	go run "${MAIN_FILE}"

# test runs test suite on tool
test: test-out
	go test \
		-outputdir "${TEST_OUT_DIR}" \
		-coverprofile "${TEST_COVER_FILE}" \
		-covermode "${TEST_COVER_MODE}" \
		./...

# test-out ensures the TEST_OUT_DIR directory exists
test-out:
	mkdir -p "${TEST_OUT_DIR}"

