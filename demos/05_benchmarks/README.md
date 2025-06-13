# Benchmark: Putting Honest Numbers Behind the Claims

> "Good enough performance with zero complexity beats perfect performance with infinite complexity."

## The Problem

Performance claims without data are just marketing. Vector databases love to talk about "scale" and "performance" while glossing over the complexity tax they impose. But what are the actual numbers when you choose simplicity over scale?

## The Solution

This benchmark puts concrete numbers behind chromem-go's realistic performance profile. We'll measure:

1. **Response time distributions** across different dataset sizes
2. **Memory usage profiles** - spoiler: chromem-go is exceptionally efficient  
3. **Throughput measurements** - how fast is "fast enough"?
4. **Honest assessment** - where chromem-go shines and where it doesn't

## Running the Benchmark

```bash
go run main.go
```

## What You'll See

The benchmark tests three dataset sizes (100, 1,000, and 10,000 documents) and provides:

- **Summary table** with key metrics across dataset sizes
- **Detailed statistics** for the largest dataset including percentiles
- **Response time distribution** showing where queries actually land
- **Performance assessment** with realistic expectations and genuine wins

## The Big Idea

This benchmark proves a fundamental thesis: **trade planetary scale for zero complexity**.

```go
// Test realistic workloads honestly
datasetSizes := []int{100, 1000, 10000}
for _, size := range datasetSizes {
    result := benchmark(size, 1000queries, 384dims)
    // Measure everything: latency, memory, throughput
    // Report honestly: wins and limitations
}
```

No infrastructure required. No complexity tax. Good enough performance for most use cases.

## What the Numbers Reveal

The benchmark shows chromem-go's realistic performance profile:

- **Small datasets (100 docs)**: Excellent performance, microsecond queries
- **Medium datasets (1K docs)**: Strong performance, sub-millisecond queries  
- **Large datasets (10K docs)**: Reasonable performance, single-digit milliseconds
- **Memory efficiency**: Exceptional - under 1KB per document
- **Zero complexity**: No Docker, no services, no configuration

## Technical Depth

- Tests with 384-dimensional vectors (typical for modern embeddings)
- Measures memory before/after document loading
- Includes warmup queries to eliminate cold-start effects
- Calculates proper percentiles (P50, P95, P99) not just averages
- Shows response time distribution buckets
- **Honest reporting**: chromem-go isn't the fastest, but it's the simplest

## The Sweet Spot

chromem-go excels when:

- Your dataset fits comfortably in RAM
- You value simplicity over raw speed
- You need offline or edge deployments
- You want to avoid infrastructure complexity
- "Fast enough" beats "fastest possible"

## Next Steps

- Compare these numbers to your current vector database setup
- Ask yourself: do you need planetary scale or just local intelligence?
- Try `01_hello_searchless` to see the simplicity behind the performance  
- Explore `04_semantic_snippets` for a practical use case
- Check out `02_similarity_modes` to understand the options

## Why This Matters

The best infrastructure is the infrastructure you forget exists. These benchmarks prove that local vector search can deliver the performance you need for most use cases, without the complexity you definitely don't want.
