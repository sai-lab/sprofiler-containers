IMAGE=docker.io/guni1192/hello-world-c
TAG=test
CONTAINER_ID_FILE=/tmp/$(IMAGE)/container_id.txt
PROFILE=/tmp/$(IMAGE)/profile.json
HOOKS=--hooks-dir /usr/share/containers/oci/hooks.d/
CID_OPT=--cidfile $(CONTAINER_ID_FILE)
ANNOTATION=--annotation "io.sprofiler.output_seccomp_profile_path=$(PROFILE)"
SECURITY_OPT=--security-opt seccomp=$(PROFILE)

LOCAL_PROFILE=$(shell pwd)/profile.json
LOCAL_SECURITY_OPT=--security-opt seccomp=$(LOCAL_PROFILE)

.PHONY: build

build:
	gcc -o hello main.c

container-build:
	sudo podman build -t $(IMAGE):$(TAG) .

run:
	rm -rf /tmp/$(IMAGE)
	mkdir -p /tmp/$(IMAGE)
	sudo podman $(HOOKS) run --rm $(CID_OPT) $(ANNOTATION) $(IMAGE):$(TAG)

test:
	sudo podman run --rm $(SECURITY_OPT) $(IMAGE):$(TAG)

test-local:
	sudo podman run --rm $(LOCAL_SECURITY_OPT) $(IMAGE):$(TAG)

clean:
	rm -rf profile.json
