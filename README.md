# JSON-Monitor

A small Go service that polls a JSON endpoint and detects real data change
with [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)

## What is checked
After each fetch, the service unmarshals the JSON response into Go data structures
(`map[string]any`, `[]any`, etc.) and compares it to the previous result using
[`reflect.DeepEqual`](https://pkg.go.dev/reflect#DeepEqual).

It will detect changes if:
- Any value changes (string, number, bool).
- Any key is added, removed, or renamed.
- Any array item changes or appears in a different order.
- A `null` value changes to a non-null value (or vice versa).
- An empty array/map changes to `null` (and vice versa).

Formatting differences (whitespace, indentation) are ignored because they are removed before comparison.

## Usage
```bash
git clone https://github.com/iliamunaev/JSON-Monitor
cd JSON-Monitor
export URL="https://example.com/data.json" SECONDS_TIMEOUT=60 # setup break between requests
go run ./cmd/monitor
```

## Use only with proper authorization.
No DoS/DDoS attacks or activity that violates laws, company policy, or service terms.
You are solely responsible for any misuse.
