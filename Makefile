.PHONY: run \
	test test-out

MAIN_FILE=main.go

TEST_OUT_DIR=test_out
TEST_COVER_FILE=coverage.out
TEST_COVER_PATH=${TEST_OUT_DIR}/${TEST_COVER_FILE}
TEST_COVER_MODE=count

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

