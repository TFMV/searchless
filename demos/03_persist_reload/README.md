# Persistence: Optional, Not Required 💾

> "The managed vector database sells you scale you don't need, wrapped in complexity you can't avoid."

## The Problem

Vector databases treat persistence as a core requirement. You need to configure storage, manage backups, and handle failures. It's infrastructure you don't always need.

## The Solution

[chromem-go](https://github.com/philippgille/chromem-go) makes persistence optional. Want to save your vectors? One function call. Want to keep them in memory? That's the default. No infrastructure required.

This demo shows how chromem-go makes persistence optional, not required. We'll:

1. Create an index with some embeddings
2. Save it to disk with a single function call
3. Exit the program
4. Reload from disk
5. Query immediately

## Running the Demo

```bash
go run main.go
```

## What You'll See

The demo demonstrates persistence in action:

- Creating and saving an index
- Reloading from disk
- Immediate querying after reload
- How persistence is a feature, not a requirement

## The Big Idea

This demo shows how chromem-go handles persistence with SQLite-like simplicity:

```go
// Save to disk
collection.Save("index.db")

// Reload from disk
collection := chromem.Load("index.db")
```

No configuration. No complexity. Just save and load.

## Technical Depth

- In-memory by default
- Optional persistence when needed
- Sub-millisecond search on reload
- Zero external dependencies

## Next Steps

- Explore `04_semantic_snippets` for a practical use case
- Return to `01_hello_searchless` to see the basics
- Try `02_similarity_modes` to understand different distance metrics

## Why This Matters

The cloud taught us to scale first, optimize later. But what if we optimized first and never needed to scale? If it fits in RAM, keep it in RAM.
