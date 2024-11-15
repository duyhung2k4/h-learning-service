.PHONY: all run_backends

all: run_backends

run_backends:
	@echo "Running backends..."
	( cd h-learning-core && go run . ) & \
	( cd h-learning-encoding-hls && go run . ) & \
	( cd h-learning-upload-mp4 && go run . ) & \
	( cd h-learning-video-hls && go run . ) & \
	wait