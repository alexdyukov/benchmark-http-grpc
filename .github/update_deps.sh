#!/usr/bin/env bash

# 1. prepare
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes -keyout example.key -out example.crt -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost"

# 2. run tests before update
go test -bench=. -benchmem -benchtime=100000x > old.txt

# 3. update dependencies
go get -u ./... && go mod tidy

# 4. configure git
git config --global user.name 'update_deps robot'
git config --global user.email 'noreply@example.com'
git remote set-url origin https://x-access-token:${GH_TOKEN}@github.com/${REPO_NAME}
git checkout -b ${BRANCH_NAME}

# 5. commit changes or fast exit
git commit -am "Fix: $(date +%F) update dependencies" || exit 0
git push --set-upstream origin ${BRANCH_NAME} -f || exit 0

# 6. run tests after update
go test -bench=. -benchmem -benchtime=100000x > new.txt

# 7. create PR with benchmark difference
gh pr create -a ${REPO_OWNER} -b "$(benchcmp old.txt new.txt)" --fill || exit 0