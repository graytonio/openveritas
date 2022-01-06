# Veritas

This project is meant to be a generic value storage system for use with IT automation such as ansible. Veritas provides a central source of truth for playbooks to pull from and push out to hosts.

## Usage

**Docker Compose (Recommended)**

```yaml
version: "3"
services:
  veritas:
    image: ghcr.io/graytonio/veritas-server:latest
    container_name: veritas
    environment:
      - MONGO_DB=veritas
      - MONGO_URI=mongodb://mongo:27017
      - PORT=8080
    ports:
      - 8080:8080
    restart: unless-stopped
  mongo:
    image: mongo
    container_name: mongo
    volumes:
      - /path/on/host:/data/db
    restart: unless-stopped
```

**Docker CLI**

```bash
docker run -d \
-e 'MONGO_DB=veritas' \
-e 'MONGO_URI=mongodb://localhost:27017' \
-e 'PORT=8080' \
-p 8080:8080 \
ghcr.io/graytonio/veritas-server:latest
```

## Features

## To Do

- [ ] Implement Proper Testing
- [ ] Property change history
- [ ] User authentication
- [ ] Anislbe plugin
- [x] Create workflows
