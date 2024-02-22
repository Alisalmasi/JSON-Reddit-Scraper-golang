# My Awesome Go Project

Reddit JSON rss Scraper 🚀

## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository: `git clone https://github.com/Alisalmasi/Reddit-JSON-Scraper-golang`
3. Navigate to the project directory: `JSON-Reddit-Scraper-golang`
4. Run the project: `go run main.go`

## Features

- Feature 1: it can scrape the Link , Title and thumbnail of Reddit entries.
- Feature 2: an then it will format it into json.
- Feature 3: lastly it will save it in a mongodb database.

## Usage

1. Run the project.
2. Make a POST request to `http://localhost:5000/parser` with the following json format:
3.  {
    "url":"https://www.reddit.com/r/recipes/.json"
    }
4. and it will return the first 25 entries and save them in database.

## API Endpoints

- `/parse` (POST): Parses data from a provided URL and returns Reddit posts.

## Environment Variables

- `MONGO_URL`: MongoDB connection URL.
- `MONGO_DATABASE`: Name of the MongoDB database.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
