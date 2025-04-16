run-all:
	cd api-gateaway && make run-server &
	cd auth && make run-server &
	cd product && make server &

kill-ports:
	@for port in 8000 50051 50052 50053; do \
		pid=$$(lsof -t -i :$$port); \
		if [ -n "$$pid" ]; then \
			echo "Killing process on port $$port (PID: $$pid)"; \
			kill -9 $$pid; \
		else \
			echo "No process found on port $$port"; \
		fi \
	done