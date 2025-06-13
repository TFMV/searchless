package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/philippgille/chromem-go"
)

func main() {
	fmt.Println("📚 Semantic Snippets Demo - Real Documentation Search")
	fmt.Println("====================================================")

	ctx := context.Background()
	start := time.Now()

	// Create database and collection
	db := chromem.NewDB()
	collection, err := db.CreateCollection("docs", nil, nil)
	if err != nil {
		panic(err)
	}

	// Create realistic documentation snippets
	documents := createDocumentationSnippets()

	fmt.Printf("📝 Loading %d documentation snippets...\n", len(documents))

	// Add documents to collection
	err = collection.AddDocuments(ctx, documents, 1)
	if err != nil {
		panic(err)
	}

	loadTime := time.Since(start)
	fmt.Printf("⚡ Loaded in: %v\n\n", loadTime)

	// Interactive search examples
	searchQueries := []struct {
		query       string
		description string
		resultCount int
	}{
		{"how to deploy applications", "Deployment & Operations", 3},
		{"debugging and troubleshooting errors", "Debugging & Error Handling", 3},
		{"database performance optimization", "Performance & Optimization", 3},
		{"security best practices", "Security Guidelines", 3},
		{"container orchestration", "Container Management", 3},
		{"API authentication methods", "Authentication & APIs", 3},
	}

	for _, sq := range searchQueries {
		performSemanticSearch(ctx, collection, sq.query, sq.description, sq.resultCount)
		fmt.Println()
	}

	// Demonstrate metadata filtering
	fmt.Println("🔍 ADVANCED SEARCH - Filtering by Category")
	fmt.Println("==========================================")

	// Search within specific categories
	fmt.Println("\n📋 Backend-only search for 'data processing':")
	backendResults, err := collection.QueryEmbedding(ctx,
		[]float32{0.4, 0.8, 0.2, 0.9, 0.3, 0.7, 0.5, 0.6, 0.1, 0.8, 0.4, 0.9, 0.2, 0.7, 0.3, 0.6},
		3,
		map[string]string{"category": "backend"},
		nil)
	if err != nil {
		panic(err)
	}

	for i, result := range backendResults {
		fmt.Printf("   %d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
		fmt.Printf("      %s\n", result.Content)
		fmt.Printf("      Category: %s | Difficulty: %s\n",
			result.Metadata["category"], result.Metadata["difficulty"])
	}

	fmt.Println("\n📋 DevOps-only search for 'monitoring':")
	devopsResults, err := collection.QueryEmbedding(ctx,
		[]float32{0.6, 0.7, 0.3, 0.8, 0.4, 0.9, 0.2, 0.5, 0.7, 0.6, 0.8, 0.3, 0.9, 0.4, 0.5, 0.2},
		3,
		map[string]string{"category": "devops"},
		nil)
	if err != nil {
		panic(err)
	}

	for i, result := range devopsResults {
		fmt.Printf("   %d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
		fmt.Printf("      %s\n", result.Content)
		fmt.Printf("      Category: %s | Difficulty: %s\n",
			result.Metadata["category"], result.Metadata["difficulty"])
	}

	// Content-based filtering
	fmt.Println("\n🔍 CONTENT FILTERING - Documents mentioning specific terms")
	fmt.Println("=========================================================")

	fmt.Println("\n📋 All documents containing 'Docker':")
	dockerResults, err := collection.QueryEmbedding(ctx,
		[]float32{0.5, 0.6, 0.4, 0.7, 0.3, 0.8, 0.2, 0.9, 0.1, 0.6, 0.5, 0.7, 0.4, 0.8, 0.3, 0.9},
		10,
		nil,
		map[string]string{"$contains": "Docker"})
	if err != nil {
		panic(err)
	}

	if len(dockerResults) == 0 {
		fmt.Println("   No documents containing 'Docker' found")
	} else {
		for i, result := range dockerResults {
			fmt.Printf("   %d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
			// Highlight the Docker mention
			content := result.Content
			content = strings.ReplaceAll(content, "Docker", "**Docker**")
			fmt.Printf("      %s\n", content)
		}
	}

	// Summary
	totalTime := time.Since(start)
	fmt.Printf("\n🎯 SUMMARY\n")
	fmt.Printf("=========\n")
	fmt.Printf("📊 Documents: %d\n", len(documents))
	fmt.Printf("⚡ Total time: %v\n", totalTime)
	fmt.Printf("🔍 Queries performed: %d\n", len(searchQueries)+3)
	fmt.Printf("💡 Average query time: ~%.2fms\n", float64(totalTime.Nanoseconds())/float64(len(searchQueries)+3)/1000000)
	fmt.Printf("🚀 Pure in-memory semantic search - no external services!\n")
}

func performSemanticSearch(ctx context.Context, collection *chromem.Collection, query, description string, count int) {
	fmt.Printf("🔍 %s\n", description)
	fmt.Printf("Query: \"%s\"\n", query)
	fmt.Println(strings.Repeat("─", 50))

	start := time.Now()

	// Create a simple embedding for the query (in a real app, you'd use the same model as for documents)
	queryEmbedding := generateQueryEmbedding(query)

	results, err := collection.QueryEmbedding(ctx, queryEmbedding, count, nil, nil)
	if err != nil {
		panic(err)
	}

	queryTime := time.Since(start)

	for i, result := range results {
		fmt.Printf("%d. [%s] Score: %.4f\n", i+1, result.ID, result.Similarity)
		fmt.Printf("   %s\n", result.Content)
		fmt.Printf("   📂 %s | 🎯 %s\n",
			result.Metadata["category"], result.Metadata["difficulty"])
	}

	fmt.Printf("⚡ Query time: %v\n", queryTime)
}

// Simple embedding generator based on query content (for demo purposes)
func generateQueryEmbedding(query string) []float32 {
	// This is a simplified approach - in reality you'd use the same embedding model
	// that was used to create the document embeddings
	hash := 0
	for _, char := range query {
		hash = hash*31 + int(char)
	}

	// Generate a 16-dimensional embedding based on query characteristics
	embedding := make([]float32, 16)
	base := float32(hash%1000) / 1000.0

	for i := range embedding {
		embedding[i] = base + float32(i)*0.1 + float32(len(query)%10)*0.01
		// Normalize to 0-1 range
		for embedding[i] > 1.0 {
			embedding[i] -= 1.0
		}
		for embedding[i] < 0.0 {
			embedding[i] += 1.0
		}
	}

	return embedding
}

func createDocumentationSnippets() []chromem.Document {
	return []chromem.Document{
		// Backend Development
		{
			ID:        "be-001",
			Content:   "Set up a REST API using Node.js and Express framework. Configure middleware for logging, CORS, and authentication. Define routes for CRUD operations.",
			Embedding: []float32{0.1, 0.9, 0.2, 0.8, 0.3, 0.7, 0.4, 0.6, 0.5, 0.8, 0.1, 0.9, 0.2, 0.7, 0.3, 0.6},
			Metadata:  map[string]string{"category": "backend", "difficulty": "intermediate", "topic": "api"},
		},
		{
			ID:        "be-002",
			Content:   "Database connection pooling in PostgreSQL. Configure maximum connections, timeout settings, and connection retry logic for production environments.",
			Embedding: []float32{0.2, 0.8, 0.3, 0.7, 0.4, 0.6, 0.5, 0.9, 0.1, 0.7, 0.2, 0.8, 0.3, 0.6, 0.4, 0.9},
			Metadata:  map[string]string{"category": "backend", "difficulty": "advanced", "topic": "database"},
		},
		{
			ID:        "be-003",
			Content:   "Implement caching strategies using Redis. Set up cache invalidation policies, handle distributed caching, and optimize cache hit ratios.",
			Embedding: []float32{0.3, 0.7, 0.4, 0.6, 0.5, 0.9, 0.1, 0.8, 0.2, 0.6, 0.3, 0.7, 0.4, 0.9, 0.5, 0.8},
			Metadata:  map[string]string{"category": "backend", "difficulty": "advanced", "topic": "performance"},
		},
		{
			ID:        "be-004",
			Content:   "Handle file uploads securely. Validate file types, limit file sizes, scan for malware, and store files in cloud storage with proper access controls.",
			Embedding: []float32{0.4, 0.6, 0.5, 0.9, 0.1, 0.8, 0.2, 0.7, 0.3, 0.5, 0.4, 0.6, 0.7, 0.8, 0.9, 0.2},
			Metadata:  map[string]string{"category": "backend", "difficulty": "intermediate", "topic": "security"},
		},

		// Frontend Development
		{
			ID:        "fe-001",
			Content:   "Create responsive layouts using CSS Grid and Flexbox. Implement mobile-first design patterns and ensure cross-browser compatibility.",
			Embedding: []float32{0.5, 0.5, 0.6, 0.4, 0.7, 0.3, 0.8, 0.2, 0.9, 0.4, 0.5, 0.6, 0.7, 0.8, 0.3, 0.2},
			Metadata:  map[string]string{"category": "frontend", "difficulty": "intermediate", "topic": "css"},
		},
		{
			ID:        "fe-002",
			Content:   "State management in React applications. Use Redux Toolkit for complex state, Context API for simple state, and implement proper state normalization.",
			Embedding: []float32{0.6, 0.4, 0.7, 0.3, 0.8, 0.2, 0.9, 0.1, 0.5, 0.3, 0.6, 0.7, 0.8, 0.2, 0.9, 0.1},
			Metadata:  map[string]string{"category": "frontend", "difficulty": "advanced", "topic": "react"},
		},
		{
			ID:        "fe-003",
			Content:   "Optimize web performance using lazy loading, code splitting, and image optimization. Implement service workers for offline functionality.",
			Embedding: []float32{0.7, 0.3, 0.8, 0.2, 0.9, 0.1, 0.6, 0.4, 0.5, 0.2, 0.7, 0.8, 0.9, 0.1, 0.6, 0.3},
			Metadata:  map[string]string{"category": "frontend", "difficulty": "advanced", "topic": "performance"},
		},
		{
			ID:        "fe-004",
			Content:   "Form validation and user input handling. Implement client-side validation, sanitize inputs, provide meaningful error messages, and handle accessibility.",
			Embedding: []float32{0.8, 0.2, 0.9, 0.1, 0.6, 0.4, 0.5, 0.3, 0.7, 0.1, 0.8, 0.9, 0.6, 0.4, 0.5, 0.7},
			Metadata:  map[string]string{"category": "frontend", "difficulty": "intermediate", "topic": "forms"},
		},

		// DevOps & Infrastructure
		{
			ID:        "do-001",
			Content:   "Deploy applications using Docker containers. Create optimized Dockerfiles, manage multi-stage builds, and implement container orchestration with Kubernetes.",
			Embedding: []float32{0.9, 0.1, 0.8, 0.2, 0.7, 0.3, 0.6, 0.4, 0.5, 0.9, 0.1, 0.8, 0.7, 0.6, 0.4, 0.5},
			Metadata:  map[string]string{"category": "devops", "difficulty": "advanced", "topic": "containers"},
		},
		{
			ID:        "do-002",
			Content:   "Set up CI/CD pipelines using GitHub Actions. Automate testing, building, and deployment processes. Configure environment-specific deployments.",
			Embedding: []float32{0.1, 0.8, 0.2, 0.9, 0.3, 0.6, 0.4, 0.7, 0.5, 0.8, 0.1, 0.9, 0.2, 0.7, 0.3, 0.6},
			Metadata:  map[string]string{"category": "devops", "difficulty": "intermediate", "topic": "cicd"},
		},
		{
			ID:        "do-003",
			Content:   "Monitor applications using Prometheus and Grafana. Set up metrics collection, create alerting rules, and build comprehensive dashboards.",
			Embedding: []float32{0.2, 0.7, 0.3, 0.8, 0.4, 0.5, 0.6, 0.9, 0.1, 0.7, 0.2, 0.8, 0.3, 0.6, 0.4, 0.9},
			Metadata:  map[string]string{"category": "devops", "difficulty": "advanced", "topic": "monitoring"},
		},
		{
			ID:        "do-004",
			Content:   "Infrastructure as Code using Terraform. Define cloud resources, manage state files, and implement proper resource lifecycle management.",
			Embedding: []float32{0.3, 0.6, 0.4, 0.7, 0.5, 0.4, 0.7, 0.8, 0.2, 0.6, 0.3, 0.7, 0.4, 0.9, 0.5, 0.8},
			Metadata:  map[string]string{"category": "devops", "difficulty": "advanced", "topic": "infrastructure"},
		},

		// Security
		{
			ID:        "sec-001",
			Content:   "Implement OAuth 2.0 authentication flow. Configure authorization servers, handle token refresh, and secure API endpoints with proper scopes.",
			Embedding: []float32{0.4, 0.5, 0.6, 0.3, 0.7, 0.2, 0.8, 0.1, 0.9, 0.5, 0.4, 0.6, 0.7, 0.8, 0.3, 0.2},
			Metadata:  map[string]string{"category": "security", "difficulty": "advanced", "topic": "authentication"},
		},
		{
			ID:        "sec-002",
			Content:   "Secure API endpoints against common attacks. Implement rate limiting, input validation, SQL injection prevention, and CSRF protection.",
			Embedding: []float32{0.5, 0.4, 0.7, 0.2, 0.8, 0.1, 0.9, 0.3, 0.6, 0.4, 0.5, 0.7, 0.8, 0.2, 0.9, 0.1},
			Metadata:  map[string]string{"category": "security", "difficulty": "intermediate", "topic": "api-security"},
		},
		{
			ID:        "sec-003",
			Content:   "Data encryption at rest and in transit. Use AES encryption for stored data, implement TLS properly, and manage encryption keys securely.",
			Embedding: []float32{0.6, 0.3, 0.8, 0.1, 0.9, 0.2, 0.7, 0.4, 0.5, 0.3, 0.6, 0.8, 0.9, 0.1, 0.7, 0.2},
			Metadata:  map[string]string{"category": "security", "difficulty": "advanced", "topic": "encryption"},
		},

		// Database
		{
			ID:        "db-001",
			Content:   "Optimize database queries for better performance. Use proper indexing strategies, analyze query execution plans, and implement query caching.",
			Embedding: []float32{0.7, 0.2, 0.9, 0.1, 0.8, 0.3, 0.6, 0.5, 0.4, 0.2, 0.7, 0.9, 0.8, 0.3, 0.6, 0.1},
			Metadata:  map[string]string{"category": "database", "difficulty": "advanced", "topic": "performance"},
		},
		{
			ID:        "db-002",
			Content:   "Database backup and recovery strategies. Implement automated backups, test restore procedures, and set up point-in-time recovery.",
			Embedding: []float32{0.8, 0.1, 0.7, 0.3, 0.6, 0.4, 0.5, 0.2, 0.9, 0.1, 0.8, 0.7, 0.6, 0.4, 0.5, 0.3},
			Metadata:  map[string]string{"category": "database", "difficulty": "intermediate", "topic": "backup"},
		},
		{
			ID:        "db-003",
			Content:   "Database migration best practices. Plan schema changes, handle data transformations, and ensure zero-downtime deployments.",
			Embedding: []float32{0.9, 0.2, 0.6, 0.4, 0.5, 0.8, 0.3, 0.7, 0.1, 0.2, 0.9, 0.6, 0.5, 0.8, 0.4, 0.7},
			Metadata:  map[string]string{"category": "database", "difficulty": "intermediate", "topic": "migration"},
		},

		// Testing
		{
			ID:        "test-001",
			Content:   "Write comprehensive unit tests using Jest and React Testing Library. Test components, hooks, and async operations with proper mocking.",
			Embedding: []float32{0.1, 0.9, 0.5, 0.6, 0.2, 0.8, 0.4, 0.7, 0.3, 0.9, 0.1, 0.5, 0.6, 0.7, 0.8, 0.4},
			Metadata:  map[string]string{"category": "testing", "difficulty": "intermediate", "topic": "unit-testing"},
		},
		{
			ID:        "test-002",
			Content:   "Integration testing for API endpoints. Test database interactions, external service calls, and end-to-end workflows with realistic data.",
			Embedding: []float32{0.2, 0.8, 0.6, 0.5, 0.3, 0.7, 0.5, 0.8, 0.4, 0.8, 0.2, 0.6, 0.7, 0.8, 0.9, 0.5},
			Metadata:  map[string]string{"category": "testing", "difficulty": "advanced", "topic": "integration-testing"},
		},
		{
			ID:        "test-003",
			Content:   "Automated browser testing with Playwright. Create reliable end-to-end tests, handle dynamic content, and implement visual regression testing.",
			Embedding: []float32{0.3, 0.7, 0.7, 0.4, 0.4, 0.6, 0.6, 0.9, 0.5, 0.7, 0.3, 0.7, 0.8, 0.9, 0.1, 0.6},
			Metadata:  map[string]string{"category": "testing", "difficulty": "advanced", "topic": "e2e-testing"},
		},

		// Performance
		{
			ID:        "perf-001",
			Content:   "Application performance monitoring and optimization. Use profiling tools, identify bottlenecks, and implement performance improvements.",
			Embedding: []float32{0.4, 0.6, 0.8, 0.3, 0.5, 0.5, 0.7, 0.1, 0.6, 0.6, 0.4, 0.8, 0.9, 0.1, 0.2, 0.7},
			Metadata:  map[string]string{"category": "performance", "difficulty": "advanced", "topic": "monitoring"},
		},
		{
			ID:        "perf-002",
			Content:   "Load testing and capacity planning. Use tools like JMeter or k6 to simulate traffic, identify system limits, and plan for scaling.",
			Embedding: []float32{0.5, 0.5, 0.9, 0.2, 0.6, 0.4, 0.8, 0.2, 0.7, 0.5, 0.5, 0.9, 0.1, 0.2, 0.3, 0.8},
			Metadata:  map[string]string{"category": "performance", "difficulty": "advanced", "topic": "load-testing"},
		},

		// Debugging
		{
			ID:        "debug-001",
			Content:   "Debug production issues using logging and monitoring tools. Set up structured logging, analyze error patterns, and implement alerting.",
			Embedding: []float32{0.6, 0.4, 0.1, 0.1, 0.7, 0.3, 0.9, 0.3, 0.8, 0.4, 0.6, 0.1, 0.2, 0.3, 0.4, 0.9},
			Metadata:  map[string]string{"category": "debugging", "difficulty": "intermediate", "topic": "production"},
		},
		{
			ID:        "debug-002",
			Content:   "Memory leak detection and resolution. Use memory profiling tools, identify leak sources, and implement proper memory management.",
			Embedding: []float32{0.7, 0.3, 0.2, 0.9, 0.8, 0.2, 0.1, 0.4, 0.9, 0.3, 0.7, 0.2, 0.3, 0.4, 0.5, 0.1},
			Metadata:  map[string]string{"category": "debugging", "difficulty": "advanced", "topic": "memory"},
		},
		{
			ID:        "debug-003",
			Content:   "Distributed tracing for microservices. Implement OpenTelemetry, trace requests across services, and analyze performance bottlenecks.",
			Embedding: []float32{0.8, 0.2, 0.3, 0.8, 0.9, 0.1, 0.2, 0.5, 0.1, 0.2, 0.8, 0.3, 0.4, 0.5, 0.6, 0.2},
			Metadata:  map[string]string{"category": "debugging", "difficulty": "advanced", "topic": "tracing"},
		},
	}
}
