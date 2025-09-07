# MarvelCTL

A command-line tool for exploring the Marvel Universe using the official Marvel Developers API. MarvelCTL lets you search for characters, get detailed information about your favorite heroes and villains, and discover their appearances across comics, series, and events.

## Features

- 🦸 **Character Search**: Find Marvel characters by name or partial matches
- 📖 **Character Information**: Get detailed information including descriptions, comics, series, and stories
- ⚙️ **Configuration Management**: Securely store and manage your Marvel API credentials
- 🔒 **Secure Authentication**: Uses Marvel's required authentication with public/private key pairs
- 🚀 **Fast and Lightweight**: Built with Go for optimal performance

## Prerequisites

Before using MarvelCTL, you'll need to obtain Marvel API credentials:

1. Visit the [Marvel Developer Portal](https://developer.marvel.com/)
2. Sign up for a free account
3. Create an application to get your public and private API keys

## Installation

### From Source

```bash
git clone https://github.com/fabiogoma/marvelctl.git
cd marvelctl
make build
```

The binary will be created as `marvelctl` in the project directory.

### Using Go Install

```bash
go install github.com/fabiogoma/marvelctl@latest
```

## Configuration

Before using MarvelCTL to query the Marvel API, you need to configure your API credentials.

### Setting Up API Keys

```bash
# Set your Marvel API public key
marvelctl config set public_key YOUR_PUBLIC_KEY

# Set your Marvel API private key
marvelctl config set private_key YOUR_PRIVATE_KEY
```

### Configuration Management Commands

```bash
# View a configuration value
marvelctl config get public_key

# Reset all configuration (removes ~/.marvelctl.yaml)
marvelctl config reset
```

Configuration is stored in `~/.marvelctl.yaml` and includes:
- `public_key`: Your Marvel API public key
- `private_key`: Your Marvel API private key (used for authentication hash generation)

## Usage

### Character Commands

#### Get Character Information

Retrieve detailed information about a specific Marvel character:

```bash
marvelctl character info "Spider-Man"
```

Example output:
```
Name: Spider-Man
Description: Bitten by a radioactive spider, high school student Peter Parker gained the speed, strength and powers of a spider. Adopting the name Spider-Man, Peter hoped to start a career using his new abilities. Taught that with great power comes great responsibility, Spidey has vowed to use his powers to help people.
Comics: 4035
Series: 846
Stories: 4982
```

#### Search for Characters

Search for characters by name or partial match:

```bash
# Search for characters starting with "Spider"
marvelctl character search "Spider"
```

Example output:
```
Name: Spider-Girl (Anya Corazon)
ID: 1009608

Name: Spider-Girl (May Parker)
ID: 1009609

Name: Spider-Man
ID: 1009610

Name: Spider-Man (Peter Parker)
ID: 1009610
```

### Configuration Commands

#### Set Configuration Values

```bash
# Set your Marvel API credentials
marvelctl config set public_key "your_public_key_here"
marvelctl config set private_key "your_private_key_here"
```

#### Get Configuration Values

```bash
# View your public key
marvelctl config get public_key

# View your private key
marvelctl config get private_key
```

#### Reset Configuration

```bash
# Remove all configuration (you'll need to set keys again)
marvelctl config reset
```

## Development

### Prerequisites

- Go 1.19 or later
- Make

### Building from Source

```bash
# Clone the repository
git clone https://github.com/fabiogoma/marvelctl.git
cd marvelctl

# Install dependencies
make deps

# Build the application
make build
```

### Available Make Targets

```bash
make build        # Build the application
make test         # Run tests
make test-coverage # Run tests with coverage
make clean        # Clean build artifacts
make fmt          # Format code
make lint         # Run linter (if golangci-lint is installed)
make vet          # Run go vet
make run          # Run the application
make install      # Install the application
make deps         # Tidy dependencies
```

### Project Structure

```
marvelctl/
├── cmd/
│   ├── character/          # Character-related commands
│   │   ├── character.go    # Main character command
│   │   ├── info.go        # Character info subcommand
│   │   └── search.go      # Character search subcommand
│   └── config/            # Configuration commands
│       ├── config.go      # Main config command
│       ├── get.go         # Get configuration values
│       ├── reset.go       # Reset configuration
│       └── set.go         # Set configuration values
├── internal/
│   ├── marvel/           # Marvel API client
│   │   └── client.go     # HTTP client and API functions
│   ├── models/           # Data models
│   │   ├── character.go  # Character data structures
│   │   └── client.go     # Client data structures
│   └── config.go         # Configuration management
├── main.go               # Application entry point
├── Makefile             # Build automation
└── README.md            # This file
```

## API Integration

MarvelCTL integrates with the [Marvel Comics API](https://developer.marvel.com/docs) using:

- **Authentication**: MD5 hash-based authentication using timestamp, private key, and public key
- **Endpoints**: Currently supports character-related endpoints
- **Rate Limiting**: Respects Marvel's API rate limits
- **Error Handling**: Comprehensive error handling for API responses

### Authentication Flow

1. Generate timestamp
2. Create MD5 hash of `timestamp + private_key + public_key`
3. Send requests with `ts`, `apikey` (public key), and `hash` parameters

## Examples

### Basic Workflow

```bash
# 1. Configure your API keys
marvelctl config set public_key "your_public_key"
marvelctl config set private_key "your_private_key"

# 2. Search for a character
marvelctl character search "Iron"

# 3. Get detailed information
marvelctl character info "Iron Man"

# 4. Explore other characters
marvelctl character info "Captain America"
marvelctl character search "X-Men"
```

### Character Search Tips

- Use partial names for broader searches: `marvelctl character search "Spider"`
- Use exact names for specific matches: `marvelctl character info "Spider-Man"`
- Character names are case-sensitive
- Some characters may have multiple variants (e.g., different versions across universes)

## Error Handling

Common errors and solutions:

### Configuration Errors

```bash
# Error: Marvel API keys not configured
marvelctl character info "Spider-Man"
# Error creating Marvel client: error fetching public key: error loading configuration file: Config File ".marvelctl.yaml" Not Found in "[/Users/username]"

# Solution: Set your API keys
marvelctl config set public_key "your_public_key"
marvelctl config set private_key "your_private_key"
```

### API Errors

```bash
# Error: Character not found
marvelctl character info "NonexistentCharacter"
# Character not found.

# Error: Invalid API credentials
# API error: {"message":"Invalid Referer"}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

MarvelCTL is an unofficial tool and is not affiliated with Marvel Entertainment, LLC. All character names, images, and related properties are trademarks and copyrights of Marvel Entertainment, LLC.

This tool is for educational and personal use only. Please respect Marvel's terms of service and API usage guidelines.

## Acknowledgments

- [Marvel Entertainment](https://www.marvel.com/) for providing the Marvel Comics API
- [Cobra CLI](https://github.com/spf13/cobra) for the command-line interface framework
- [Viper](https://github.com/spf13/viper) for configuration management

---

**Author**: Fabio Gonçalves Martins (fabiogoma@gmail.com)

**Repository**: https://github.com/fabiogoma/marvelctl