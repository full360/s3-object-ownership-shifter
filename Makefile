MAIN				= cmd/s3copier
BIN					= main
DIR					= s3copier
OUTPUT_DIR	 		= ./build

.PHONY: help
.DEFAULT_GOAL := help

build:
	mkdir -p $(OUTPUT_DIR)/${DIR} && GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o $(OUTPUT_DIR)/$(DIR)/$(BIN) ./$(MAIN)

release: build ## Zip a linux binary as AWS Deployment archive
	cd ${OUTPUT_DIR}/${DIR} && zip $(BIN).zip $(BIN)

clean:
	$(RM) ${OUTPUT_DIR}/${DIR}/$(BIN).zip
	$(RM) ${OUTPUT_DIR}/${DIR}/$(BIN)

