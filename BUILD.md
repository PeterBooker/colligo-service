# Colligo Build Instructions

`go install -tags=dev -ldflags="-s -w -X main.version=0.1 -X main.commit=$(git rev-parse --verify HEAD) -X main.date=$((Get-Date).toString("yyyy-MM-dd"))"`