.PHONY: build

ginkgo := go run github.com/onsi/ginkgo/v2/ginkgo -r --race --randomize-all --fail-on-pending

test: test-unit

test-unit:
	$(ginkgo)
