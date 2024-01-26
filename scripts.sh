docker run -d --name ddgodeliv-postgres -p 5432:5432 -e POSTGRES_PASSWORD=ddgodeliv-password -v postgres:/var/lib/postgresql/data postgres:16.1-alpine3.18
