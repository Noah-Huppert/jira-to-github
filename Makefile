.PHONY: test test-out

TEST_OUT_DIR=test_out
TEST_COVER_FILE=coverage.out
TEST_COVER_PATH=${TEST_OUT_DIR}/${TEST_COVER_FILE}
TEST_COVER_MODE=count

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

