GO := go

OUT_DIR := ./bin

run:
	${GO} run ./cmd/server

build:
	${GO} build -o ${OUT_DIR}/server ./cmd/server

tidy:
	${GO} mod tidy