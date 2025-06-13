# Hello Searchless! 🚀

> "I needed to search ten documents. The tutorial told me to install Kubernetes."

## The Problem

Traditional vector databases require complex infrastructure, API keys, and external services. They're overkill for simple semantic search needs.

## The Solution

[chromem-go](https://github.com/philippgille/chromem-go) makes semantic search as simple as SQLite makes databases. No infrastructure, no API keys, just pure Go code.

This demo proves a radical idea: semantic search doesn't need infrastructure. In this small Go application, we'll:

1. Load pre-computed embeddings for various tech concepts
2. Query them with a "database" embedding
3. Show ranked results with similarity scores
4. Time the entire operation (spoiler: it's sub-millisecond)

## Running the Demo

```bash
go run main.go
```

## What You'll See

The demo loads 10 pre-computed embeddings and performs a semantic search. You'll see:

- Ranked results with similarity scores
- The time taken for the entire operation
- How chromem-go makes semantic search feel like opening a file

## The Big Idea

This isn't just a demo - it's a proof of concept that semantic search can be as simple as:

```go
collection := chromem.NewCollection()
collection.Add(documents)
results := collection.Query(query)
```

No infrastructure. No complexity. Just search.

## Next Steps

- Try `02_similarity_modes` to see different distance metrics in action
- Check out `03_persist_reload` to learn about persistence
- Explore `04_semantic_snippets` for a practical use case

## Why This Matters

The best infrastructure disappears. This one vanishes into your codebase.
