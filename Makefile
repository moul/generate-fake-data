GOPKG ?=	moul.io/generate-fake-data
DOCKER_IMAGE ?=	moul/generate-fake-data
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	mkdir -p .tmp
	GO111MODULE=off go get github.com/campoy/embedmd
	generate-fake-data -h 2>.tmp/usage.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=random      > .tmp/random.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=address     > .tmp/address.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=apache-log  > .tmp/apache-log.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=beer        > .tmp/beer.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=hacker      > .tmp/hacker.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=hipster     > .tmp/hipster.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=lorem-ipsum > .tmp/lorem-ipsum.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=phrase      > .tmp/phrase.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=question    > .tmp/question.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=quote       > .tmp/quote.txt
	generate-fake-data --no-stderr --lines=10 --sleep-max=0 -seed=42 --dict=uuid        > .tmp/uuid.txt
	embedmd -w README.md
	rm -rf .tmp
