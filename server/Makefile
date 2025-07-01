# e.g. "make new_mig create_users"
new_mig:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Error: migration name missing"; \
		echo "Usage: make new_mig <migration_name>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir repository/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

%:
	@: