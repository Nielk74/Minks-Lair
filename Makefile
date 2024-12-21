create-migration: $(name)
	echo "Creating migration file"
	migrate create -ext sql -dir migrations -seq $(name)
	echo "Migration file created"