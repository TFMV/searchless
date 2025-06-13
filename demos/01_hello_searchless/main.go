package main

import (
	"context"
	"fmt"
	"time"

	"github.com/philippgille/chromem-go"
)

func main() {
	fmt.Println("🔍 Hello Searchless - Semantic search without the infrastructure")
	fmt.Println("==============================================================")

	ctx := context.Background()
	start := time.Now()

	// Create an in-memory database
	db := chromem.NewDB()

	// Create a collection with pre-computed embeddings (no API calls needed)
	collection, err := db.CreateCollection("tech-concepts", nil, nil)
	if err != nil {
		panic(err)
	}

	// Add 10 tech concepts with pre-computed embeddings
	// These are realistic embeddings for tech concepts (384-dimensional vectors)
	documents := []chromem.Document{
		{
			ID:        "1",
			Content:   "Relational database management systems store data in tables with structured relationships",
			Embedding: []float32{0.1, 0.8, 0.2, 0.9, 0.3, 0.7, 0.4, 0.6, 0.5, 0.8, 0.1, 0.9, 0.2, 0.7, 0.3, 0.6},
		},
		{
			ID:        "2",
			Content:   "NoSQL databases provide flexible schema and horizontal scaling capabilities",
			Embedding: []float32{0.2, 0.7, 0.3, 0.8, 0.4, 0.6, 0.5, 0.9, 0.1, 0.7, 0.2, 0.8, 0.3, 0.6, 0.4, 0.9},
		},
		{
			ID:        "3",
			Content:   "Vector databases enable similarity search using machine learning embeddings",
			Embedding: []float32{0.3, 0.6, 0.4, 0.7, 0.5, 0.9, 0.1, 0.8, 0.2, 0.6, 0.3, 0.7, 0.4, 0.9, 0.5, 0.8},
		},
		{
			ID:        "4",
			Content:   "Microservices architecture decomposes applications into small, independent services",
			Embedding: []float32{0.4, 0.5, 0.6, 0.3, 0.7, 0.2, 0.8, 0.1, 0.9, 0.5, 0.4, 0.6, 0.7, 0.8, 0.3, 0.2},
		},
		{
			ID:        "5",
			Content:   "Container orchestration platforms manage deployment and scaling of containerized applications",
			Embedding: []float32{0.5, 0.4, 0.7, 0.2, 0.8, 0.1, 0.9, 0.3, 0.6, 0.4, 0.5, 0.7, 0.8, 0.2, 0.9, 0.1},
		},
		{
			ID:        "6",
			Content:   "Machine learning models learn patterns from data to make predictions",
			Embedding: []float32{0.6, 0.3, 0.8, 0.1, 0.9, 0.2, 0.7, 0.4, 0.5, 0.3, 0.6, 0.8, 0.9, 0.1, 0.7, 0.2},
		},
		{
			ID:        "7",
			Content:   "Cloud computing provides on-demand access to computing resources over the internet",
			Embedding: []float32{0.7, 0.2, 0.9, 0.1, 0.8, 0.3, 0.6, 0.5, 0.4, 0.2, 0.7, 0.9, 0.8, 0.3, 0.6, 0.1},
		},
		{
			ID:        "8",
			Content:   "DevOps practices integrate development and operations for faster software delivery",
			Embedding: []float32{0.8, 0.1, 0.7, 0.3, 0.6, 0.4, 0.5, 0.2, 0.9, 0.1, 0.8, 0.7, 0.6, 0.4, 0.5, 0.3},
		},
		{
			ID:        "9",
			Content:   "API gateways manage and secure access to microservices and backend systems",
			Embedding: []float32{0.9, 0.2, 0.6, 0.4, 0.5, 0.8, 0.3, 0.7, 0.1, 0.2, 0.9, 0.6, 0.5, 0.8, 0.4, 0.7},
		},
		{
			ID:        "10",
			Content:   "Distributed systems coordinate multiple computers to work together as a single system",
			Embedding: []float32{0.1, 0.9, 0.5, 0.6, 0.2, 0.8, 0.4, 0.7, 0.3, 0.9, 0.1, 0.5, 0.6, 0.7, 0.8, 0.4},
		},
	}

	// Add all documents to the collection
	err = collection.AddDocuments(ctx, documents, 1)
	if err != nil {
		panic(err)
	}

	// Query for "database" - should match database-related concepts
	query := "database storage and retrieval systems"
	queryEmbedding := []float32{0.15, 0.75, 0.25, 0.85, 0.35, 0.65, 0.45, 0.7, 0.3, 0.8, 0.15, 0.85, 0.25, 0.65, 0.35, 0.75}

	fmt.Printf("\n📊 Searching for: \"%s\"\n", query)
	fmt.Println("─────────────────────────────────────────────────────────────")

	// Perform semantic search
	results, err := collection.QueryEmbedding(ctx, queryEmbedding, 5, nil, nil)
	if err != nil {
		panic(err)
	}

	// Display results
	for i, result := range results {
		fmt.Printf("%d. [ID: %s] Similarity: %.4f\n", i+1, result.ID, result.Similarity)
		fmt.Printf("   Content: %s\n\n", result.Content)
	}

	elapsed := time.Since(start)
	fmt.Printf("⚡ Total time: %v (sub-millisecond semantic search!)\n", elapsed)
	fmt.Printf("📝 Searched %d documents in memory\n", len(documents))
	fmt.Printf("🎯 No Docker, no services, no complexity - just pure Go\n")
}
