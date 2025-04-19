run-all:
	cd api-gateaway && make server &
	cd auth && make server &
	cd product && make server &
	cd order && make server &

run-services:
	cd auth && make server &
	cd product && make server &
	cd order && make server &

kill-ports:
	@for port in 50051 50052 50053; do \
		pid=$$(lsof -t -i :$$port); \
		if [ -n "$$pid" ]; then \
			echo "Killing process on port $$port (PID: $$pid)"; \
			kill -9 $$pid; \
		else \
			echo "No process found on port $$port"; \
		fi \
	done