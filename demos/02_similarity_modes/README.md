# Similarity Modes: One Library, Many Metrics 🔄

> "SQLite didn't win because it scaled to petabytes. It won because it disappeared into applications."

## The Problem

Vector databases often lock you into a single similarity metric. Want to try a different one? That's a different database, different API, different infrastructure.

## The Solution

[chromem-go](https://github.com/philippgille/chromem-go) makes switching similarity metrics as simple as changing a function parameter. No new infrastructure, no new API keys, just a different function call.

This demo shows how chromem-go makes similarity metrics as simple as a function parameter. We'll:

1. Load the same set of embeddings
2. Query them using three different distance metrics:
   - Cosine similarity
   - Dot product
   - Euclidean distance
3. Compare how results rank differently
4. Show that changing metrics is just a parameter change

## Running the Demo

```bash
go run main.go
```

## What You'll See

The demo performs the same semantic search using different similarity metrics. You'll observe:

- How each metric ranks results differently
- The relative strengths of each approach
- That changing metrics is as simple as changing a parameter

## The Big Idea

This demo embodies chromem-go's "SQLite moment" - the realization that not everything needs to be a service. Just like SQLite made databases disappear into applications, chromem-go makes vector search disappear into your code:

```go
// Change similarity mode with a parameter
results := collection.Query(query, chromem.WithSimilarityMode(mode))
```

## Technical Depth

- Cosine similarity: Best for normalized embeddings
- Dot product: Faster but sensitive to vector magnitude
- Euclidean distance: Intuitive but computationally more expensive

## Next Steps

- Try `03_persist_reload` to see how state persists across restarts
- Explore `04_semantic_snippets` for a practical use case
- Return to `01_hello_searchless` to see the basics

## Why This Matters

The best infrastructure is the infrastructure you forget exists. chromem-go doesn't want to be your vector platform - it wants to be a function call that happens to remember things.
