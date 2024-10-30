# RSS Feed Aggregator

A Go-based RSS feed aggregator that collects and stores posts from multiple RSS feeds in a PostgreSQL database.

## Features

- Fetch and parse RSS feeds
- Store feed data in PostgreSQL
- Handle multiple RSS feeds concurrently
- Support for various date formats
- Automatic post deduplication
- RESTful API endpoints for feed management

## Prerequisites

- Go 1.x
- PostgreSQL
- [Goose](https://github.com/pressly/goose) for database migrations
- [SQLC](https://sqlc.dev/) for type-safe SQL

## API Endpoints

- `POST /v1/feeds` - Create a new feed
- `GET /v1/feeds` - Get all feeds
- `GET /v1/posts` - Get all posts

## Database Schema

The application uses the following main tables:

- `users` - Stores user information
- `feeds` - Stores RSS feed information
- `feed follows` - Stores feed followed by users
- `posts` - Stores individual posts from feeds

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/rss-feed-aggregator.git
   cd rss-feed-aggregator
   ```

2. Set up the database:

   ```bash
   # Create PostgreSQL database
   createdb rss_aggregator

   # Run migrations
   goose postgres "postgresql://localhost:5432/rss_aggregator?sslmode=disable" up
   ```

3. Configure environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials and other settings
   ```

4. Install dependencies:

   ```bash
   go mod download
   ```

5. Build and run the application:

   ```bash
   go build
   ./rss-feed-aggregator
   ```

6. Test the API:

   ```bash
   # Create a new feed
   curl -X POST -H "Content-Type: application/json" \
        -d '{"name": "Example Feed", "url": "https://example.com/feed.xml"}' \
        http://localhost:8080/v1/feeds

   # Get all feeds
   curl http://localhost:8080/v1/feeds
   ```

## Development

- Generate SQL code: `sqlc generate`
- Update database schema: `goose postgres "postgresql://localhost:5432/rss_aggregator?sslmode=disable" up`
