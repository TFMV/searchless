package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/philippgille/chromem-go"
)

func main() {
	fmt.Println("💾 Persist & Reload Demo - SQLite-like Persistence")
	fmt.Println("=================================================")

	ctx := context.Background()
	dbPath := "./chromem-data"

	// Check if database already exists
	if _, err := os.Stat(dbPath); err == nil {
		fmt.Println("📁 Found existing database, loading...")
		loadAndQuery(ctx, dbPath)
	} else {
		fmt.Println("🆕 No existing database found, creating new one...")
		createAndSave(ctx, dbPath)
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("💡 Run this program again to see the persistence in action!")
		fmt.Println("   The database will reload instantly from disk")
	}
}

func createAndSave(ctx context.Context, dbPath string) {
	fmt.Println("\n📝 Creating persistent database...")

	// Create persistent database
	db, err := chromem.NewPersistentDB(dbPath, false)
	if err != nil {
		panic(err)
	}

	// Create collection
	collection, err := db.CreateCollection("knowledge-base",
		map[string]string{
			"description": "Technical documentation snippets",
			"version":     "1.0",
		}, nil)
	if err != nil {
		panic(err)
	}

	// Add documents with embeddings
	documents := []chromem.Document{
		{
			ID:        "doc-001",
			Content:   "Docker containers provide lightweight, portable application packaging",
			Embedding: []float32{0.1, 0.9, 0.3, 0.7, 0.5, 0.8, 0.2, 0.6, 0.4, 0.9, 0.1, 0.8, 0.3, 0.7, 0.5, 0.6},
			Metadata: map[string]string{
				"category":   "containerization",
				"difficulty": "beginner",
			},
		},
		{
			ID:        "doc-002",
			Content:   "Kubernetes orchestrates containers across clusters with automated scaling",
			Embedding: []float32{0.2, 0.8, 0.4, 0.6, 0.3, 0.9, 0.1, 0.7, 0.5, 0.8, 0.2, 0.6, 0.4, 0.9, 0.3, 0.7},
			Metadata: map[string]string{
				"category":   "orchestration",
				"difficulty": "advanced",
			},
		},
		{
			ID:        "doc-003",
			Content:   "Microservice architecture breaks applications into independent, deployable services",
			Embedding: []float32{0.3, 0.7, 0.5, 0.9, 0.1, 0.6, 0.2, 0.8, 0.4, 0.7, 0.3, 0.9, 0.5, 0.6, 0.1, 0.8},
			Metadata: map[string]string{
				"category":   "architecture",
				"difficulty": "intermediate",
			},
		},
		{
			ID:        "doc-004",
			Content:   "REST APIs enable communication between services using HTTP protocols",
			Embedding: []float32{0.4, 0.6, 0.2, 0.8, 0.3, 0.7, 0.5, 0.9, 0.1, 0.6, 0.4, 0.8, 0.2, 0.7, 0.3, 0.9},
			Metadata: map[string]string{
				"category":   "api",
				"difficulty": "beginner",
			},
		},
		{
			ID:        "doc-005",
			Content:   "Database sharding distributes data across multiple database instances",
			Embedding: []float32{0.5, 0.9, 0.1, 0.7, 0.3, 0.8, 0.4, 0.6, 0.2, 0.9, 0.5, 0.7, 0.1, 0.8, 0.3, 0.6},
			Metadata: map[string]string{
				"category":   "database",
				"difficulty": "advanced",
			},
		},
	}

	fmt.Printf("   Adding %d documents to collection...\n", len(documents))
	err = collection.AddDocuments(ctx, documents, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("   ✅ Documents added and persisted to disk")

	// Show what was saved
	fmt.Println("\n📂 Database structure on disk:")
	err = filepath.Walk(dbPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(dbPath, path)
		if relPath == "." {
			fmt.Printf("   %s/\n", filepath.Base(dbPath))
		} else {
			indent := "   "
			for i := 0; i < len(filepath.SplitList(filepath.Dir(relPath))); i++ {
				indent += "  "
			}
			if info.IsDir() {
				fmt.Printf("%s%s/\n", indent, info.Name())
			} else {
				fmt.Printf("%s%s (%d bytes)\n", indent, info.Name(), info.Size())
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("   Error reading directory: %v\n", err)
	}

	// Perform a test query
	fmt.Println("\n🔍 Testing query before exit...")
	queryEmbedding := []float32{0.15, 0.85, 0.35, 0.65, 0.45, 0.75, 0.25, 0.55, 0.4, 0.8, 0.15, 0.7, 0.35, 0.6, 0.45, 0.65}
	results, err := collection.QueryEmbedding(ctx, queryEmbedding, 2, nil, nil)
	if err != nil {
		panic(err)
	}

	for i, result := range results {
		fmt.Printf("   %d. [%s] Similarity: %.4f\n", i+1, result.ID, result.Similarity)
		fmt.Printf("      %s\n", result.Content)
	}
}

func loadAndQuery(ctx context.Context, dbPath string) {
	fmt.Println("\n🔄 Loading database from disk...")

	// Load existing persistent database
	db, err := chromem.NewPersistentDB(dbPath, false)
	if err != nil {
		panic(err)
	}

	// List collections
	collections := db.ListCollections()
	fmt.Printf("   Found %d collections:\n", len(collections))

	for name, collection := range collections {
		fmt.Printf("   - %s (%d documents)\n", name, collection.Count())

		// Get the collection (we need to provide a dummy embedding function)
		coll := db.GetCollection(name, nil)
		if coll == nil {
			continue
		}

		fmt.Println("\n🔍 Performing instant queries (no loading time!)...")

		// Query 1: Container-related
		fmt.Println("\n   Query 1: 'container technology'")
		queryEmbedding1 := []float32{0.15, 0.85, 0.35, 0.65, 0.45, 0.75, 0.25, 0.55, 0.4, 0.8, 0.15, 0.7, 0.35, 0.6, 0.45, 0.65}
		results1, err := coll.QueryEmbedding(ctx, queryEmbedding1, 3, nil, nil)
		if err != nil {
			panic(err)
		}

		for i, result := range results1 {
			fmt.Printf("      %d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
			fmt.Printf("         %s\n", result.Content)
		}

		// Query 2: Architecture-related with metadata filter
		fmt.Println("\n   Query 2: 'service architecture' (beginner level only)")
		queryEmbedding2 := []float32{0.25, 0.75, 0.45, 0.55, 0.35, 0.85, 0.15, 0.65, 0.3, 0.9, 0.25, 0.6, 0.45, 0.7, 0.35, 0.8}
		results2, err := coll.QueryEmbedding(ctx, queryEmbedding2, 5,
			map[string]string{"difficulty": "beginner"}, nil)
		if err != nil {
			panic(err)
		}

		if len(results2) == 0 {
			fmt.Println("      No results found for beginner difficulty")
		} else {
			for i, result := range results2 {
				fmt.Printf("      %d. [%s] Score: %.4f (Category: %s)\n",
					i+1, result.ID, result.Similarity, result.Metadata["category"])
				fmt.Printf("         %s\n", result.Content)
			}
		}

		// Query 3: Content filter
		fmt.Println("\n   Query 3: Documents containing 'API'")
		queryEmbedding3 := []float32{0.35, 0.65, 0.25, 0.75, 0.45, 0.55, 0.35, 0.85, 0.15, 0.7, 0.35, 0.8, 0.25, 0.9, 0.45, 0.6}
		results3, err := coll.QueryEmbedding(ctx, queryEmbedding3, 5, nil,
			map[string]string{"$contains": "API"})
		if err != nil {
			panic(err)
		}

		if len(results3) == 0 {
			fmt.Println("      No documents containing 'API' found")
		} else {
			for i, result := range results3 {
				fmt.Printf("      %d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
				fmt.Printf("         %s\n", result.Content)
			}
		}
	}

	fmt.Println("\n🎯 Key Benefits:")
	fmt.Println("   ✅ Instant startup - no re-indexing required")
	fmt.Println("   ✅ Data persists between program runs")
	fmt.Println("   ✅ No external database server needed")
	fmt.Println("   ✅ SQLite-like simplicity with vector search power")
	fmt.Println("\n💡 Try deleting the 'chromem-data' folder and run again!")
}
