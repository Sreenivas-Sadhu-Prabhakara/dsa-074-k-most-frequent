# K Most Frequent Elements

> Striver SDE Sheet 074 · Day 12 Heaps — counts are bounded by n, so bucket-by-frequency beats the heap: linear time, no ordering structure at all. Go, stdlib only, tested.

## Problem

Return the k values that occur most often in an array. The usual pitch is "count, then heap by frequency" — O(n log k).

## Approach

Count occurrences, then exploit a bound the heap ignores: a frequency is always between 1 and n, so values can be dropped into an array of buckets indexed *by frequency*. Walking buckets from n downwards collects the top k in linear time — frequency is a small-universe key, and small-universe keys sort by addressing, not comparing. Ties inside a bucket are sorted so the output is deterministic (and testable). The heap version remains the right call when the data streams in or n is unbounded; for a materialized array, buckets win.

**Complexity:** O(n + d log d) where d is distinct values in the emitted buckets — effectively O(n).

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
