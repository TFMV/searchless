package main

import (
	"context"
	"fmt"
	"math"

	"github.com/philippgille/chromem-go"
)

// Custom similarity functions to demonstrate different approaches
func cosineSimilarity(a, b []float32) float32 {
	var dotProduct, normA, normB float32
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	return dotProduct / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))
}

func euclideanDistance(a, b []float32) float32 {
	var sum float32
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return float32(math.Sqrt(float64(sum)))
}

func manhattanDistance(a, b []float32) float32 {
	var sum float32
	for i := range a {
		sum += float32(math.Abs(float64(a[i] - b[i])))
	}
	return sum
}

// Convert distance to similarity (lower distance = higher similarity)
func distanceToSimilarity(distance float32) float32 {
	return 1.0 / (1.0 + distance)
}

func main() {
	fmt.Println("🔍 Similarity Modes Demo - Same Query, Different Perspectives")
	fmt.Println("==========================================================")

	ctx := context.Background()

	// Create database and collection
	db := chromem.NewDB()
	collection, err := db.CreateCollection("documents", nil, nil)
	if err != nil {
		panic(err)
	}

	// Sample documents with their embeddings
	documents := []chromem.Document{
		{
			ID:        "doc1",
			Content:   "Database systems for storing and retrieving structured data efficiently",
			Embedding: []float32{0.8, 0.9, 0.1, 0.2, 0.7, 0.8, 0.3, 0.4},
		},
		{
			ID:        "doc2",
			Content:   "Machine learning algorithms for pattern recognition and prediction",
			Embedding: []float32{0.2, 0.3, 0.9, 0.8, 0.1, 0.2, 0.7, 0.6},
		},
		{
			ID:        "doc3",
			Content:   "Web development frameworks for building interactive applications",
			Embedding: []float32{0.5, 0.4, 0.6, 0.7, 0.5, 0.6, 0.4, 0.3},
		},
		{
			ID:        "doc4",
			Content:   "Data storage solutions with high availability and scalability",
			Embedding: []float32{0.7, 0.8, 0.2, 0.3, 0.6, 0.7, 0.4, 0.5},
		},
		{
			ID:        "doc5",
			Content:   "Cloud computing platforms for scalable application deployment",
			Embedding: []float32{0.3, 0.4, 0.5, 0.6, 0.8, 0.7, 0.6, 0.5},
		},
	}

	// Add documents to collection
	err = collection.AddDocuments(ctx, documents, 1)
	if err != nil {
		panic(err)
	}

	// Query embedding
	query := "data storage and database systems"
	queryEmbedding := []float32{0.75, 0.85, 0.15, 0.25, 0.65, 0.75, 0.35, 0.45}

	fmt.Printf("🔍 Query: \"%s\"\n", query)
	fmt.Println("📐 Embedding:", queryEmbedding)
	fmt.Println()

	// 1. Default chromem-go similarity (cosine/dot product)
	fmt.Println("📊 1. COSINE SIMILARITY (chromem-go default)")
	fmt.Println("   Higher values = more similar")
	fmt.Println("   ─────────────────────────────")

	results, err := collection.QueryEmbedding(ctx, queryEmbedding, 5, nil, nil)
	if err != nil {
		panic(err)
	}

	for i, result := range results {
		fmt.Printf("   %d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
		fmt.Printf("      %s\n", result.Content)
	}

	// 2. Manual Euclidean Distance calculation
	fmt.Println("\n📊 2. EUCLIDEAN DISTANCE")
	fmt.Println("   Lower values = more similar")
	fmt.Println("   ─────────────────────────────")

	type docDistance struct {
		id         string
		content    string
		distance   float32
		similarity float32
	}

	var euclideanResults []docDistance
	for _, doc := range documents {
		distance := euclideanDistance(queryEmbedding, doc.Embedding)
		similarity := distanceToSimilarity(distance)
		euclideanResults = append(euclideanResults, docDistance{
			id:         doc.ID,
			content:    doc.Content,
			distance:   distance,
			similarity: similarity,
		})
	}

	// Sort by distance (ascending)
	for i := 0; i < len(euclideanResults); i++ {
		for j := i + 1; j < len(euclideanResults); j++ {
			if euclideanResults[i].distance > euclideanResults[j].distance {
				euclideanResults[i], euclideanResults[j] = euclideanResults[j], euclideanResults[i]
			}
		}
	}

	for i, result := range euclideanResults {
		fmt.Printf("   %d. [%s] Distance: %.4f, Similarity: %.4f\n", i+1, result.id, result.distance, result.similarity)
		fmt.Printf("      %s\n", result.content)
	}

	// 3. Manual Manhattan Distance calculation
	fmt.Println("\n📊 3. MANHATTAN DISTANCE")
	fmt.Println("   Lower values = more similar")
	fmt.Println("   ─────────────────────────────")

	var manhattanResults []docDistance
	for _, doc := range documents {
		distance := manhattanDistance(queryEmbedding, doc.Embedding)
		similarity := distanceToSimilarity(distance)
		manhattanResults = append(manhattanResults, docDistance{
			id:         doc.ID,
			content:    doc.Content,
			distance:   distance,
			similarity: similarity,
		})
	}

	// Sort by distance (ascending)
	for i := 0; i < len(manhattanResults); i++ {
		for j := i + 1; j < len(manhattanResults); j++ {
			if manhattanResults[i].distance > manhattanResults[j].distance {
				manhattanResults[i], manhattanResults[j] = manhattanResults[j], manhattanResults[i]
			}
		}
	}

	for i, result := range manhattanResults {
		fmt.Printf("   %d. [%s] Distance: %.4f, Similarity: %.4f\n", i+1, result.id, result.distance, result.similarity)
		fmt.Printf("      %s\n", result.content)
	}

	fmt.Println("\n🎯 Key Insights:")
	fmt.Println("   • Cosine similarity focuses on vector direction (angle)")
	fmt.Println("   • Euclidean distance measures straight-line distance in space")
	fmt.Println("   • Manhattan distance measures grid-like distance")
	fmt.Println("   • Different metrics can rank results differently!")
	fmt.Println("   • chromem-go uses cosine similarity for best semantic matching")
}
