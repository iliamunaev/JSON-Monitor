# JSON-Monitor

## (In progress)
A small Go service that polls a JSON endpoint and detects real data change

## Usage
```bash
git clone https://github.com/iliamunaev/JSON-Monitor
cd JSON-Monitor
export URL="https://example.com/data.json" SECONDS_TIMEOUT=60 # setup brake between requests
go run ./cmd/monitor
```

## Use only with proper authorization.
No DoS/DDoS attacks or activity that violates laws, company policy, or service terms.
You are solely responsible for any misuse.
