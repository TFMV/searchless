# Searchless

**Local Intelligence Patterns**

This repository contains the code and demos for the article "SQLite for Embeddings". The article explores how [chromem-go](https://github.com/philippgille/chromem-go) is making semantic search as simple as SQLite made databases.

## Semantic Search Without the Infra

> *"I needed to search ten documents. The tutorial told me to install Kubernetes."*

## What Is This?

This repository demonstrates **chromem-go**, the SQLite moment for vector embeddings. It's the companion code for the Medium article ["SQLite for Embeddings: chromem-go and the Return of Local Intelligence"](https://medium.com) - a manifesto for local-first semantic search.

Every vector database tutorial starts the same way: install Docker, configure persistence, set up authentication, expose ports, monitor resources. By the time you're ready to search your first embedding, you've architected a distributed system.

**But what if you didn't have to?**

## The Philosophy

chromem-go is having its "SQLite moment" - the realization that not everything needs to be a service. SQLite didn't win because it scaled to petabytes. It won because it disappeared into applications. No setup. No tuning. No 3am pages about connection pools. Just include it and it works.

chromem-go has that same quality: **it vanishes into your binary**.

## The Problem: Embedding Inflation

We've coined a term for what's happening in vector search: **Embedding Inflation** - when simple vector operations get wrapped in enterprise infrastructure.

The managed vector database sells you scale you don't need, wrapped in complexity you can't avoid:

- Rate limits on your own data
- Network latency on local queries  
- Authentication tokens for embeddings you computed
- Meanwhile, your use case fits in memory and runs faster locally

## The Solution: Local Intelligence Patterns

**If it fits in RAM, keep it in RAM. Skip the network hop.**

This repository showcases practical applications where local beats hosted:

- CLI tools that need offline semantic search
- Edge deployments with intermittent connectivity
- Developer tools doing code similarity matching
- Internal apps with embedding-based recommendations
- Observability systems correlating traces semantically

## Demo Progression

Each demo builds on the previous one, proving that semantic search can be as simple as opening a file:

### 🔍 [01_hello_searchless](./demos/01_hello_searchless/)

**"The 50-Line Proof"**

Prove the central thesis in 50 lines of Go. Load 10 tech concepts, query for "database", see ranked results. Time the entire operation and watch infrastructure complexity disappear.

*Key insight: Semantic search doesn't need services - it needs functions.*

### ⚖️ [02_similarity_modes](./demos/02_similarity_modes/)

**"The SQLite Moment"**

Show how chromem-go makes similarity metrics as simple as a parameter change. Same data, three different distance functions (cosine, dot product, Euclidean). No reconfiguration required.

*Key insight: Flexibility without infrastructure overhead.*

### 💾 [03_persist_reload](./demos/03_persist_reload/)

**"Persistence Without Complexity"**

Create an index, save to disk, exit. Reload and query immediately. Demonstrate how chromem-go makes persistence optional, not required - just like SQLite.

*Key insight: State management without distributed systems.*

### 📖 [04_semantic_snippets](./demos/04_semantic_snippets/)

**"Real-World Intelligence"**

Search over documentation snippets with actual semantic understanding. See how local intelligence beats keyword matching for practical use cases.

*Key insight: Context awareness without network calls.*

### 📊 [benchmarks](./demos/05_benchmarks/)

**"Honest Numbers"**

Put concrete data behind the performance claims. Measure response times, memory usage, and throughput across different dataset sizes. See where chromem-go shines and where it doesn't.

*Key insight: Good enough performance with zero complexity.*

## Running the Demos

Each demo is self-contained with its own README and can be run independently:

```bash
# Try them in order for the full story
cd demos/01_hello_searchless && go run main.go
cd ../02_similarity_modes && go run main.go  
cd ../03_persist_reload && go run main.go
cd ../04_semantic_snippets && go run main.go
cd ../benchmarks && go run main.go
```

## Key Insights

### Technical Wins

- **Exceptional memory efficiency**: Under 1KB per document
- **Predictable performance**: Millisecond search on thousands of documents
- **Zero dependencies**: Pure Go, no external services
- **Optional persistence**: Save state when you want, skip when you don't

### Philosophical Wins

- **Disappearing infrastructure**: Like `json.Marshal()` but for similarity
- **Local-first intelligence**: Your data, your machine, your control
- **Simplicity over scale**: Optimize for the 99% of use cases that fit in RAM

## Contributing

Found an issue? Have a use case that would make a great demo? Want to share your own local intelligence pattern? [Open an issue](https://github.com/user/searchless/issues) or submit a pull request.

## Credits

- **chromem-go**: The brilliant work of [Philipp Gille](https://github.com/philippgille/chromem-go)

---
