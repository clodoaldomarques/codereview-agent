# Code Review Agent

> Experimental Go-based API for AI-assisted code review using a locally hosted LLM with Ollama.

## Overview

Code Review Agent is a small Go application created to explore the integration between backend services and local Large Language Models (LLMs).

The application exposes an HTTP API that receives source code and sends it to a locally running Ollama model, generating suggestions for code improvements or simple unit tests.

The project was created as a practical experiment in applying Generative AI to software engineering workflows.

## Architecture

The application follows a simple request/response architecture:

```text
                    HTTP Client
                         │
                         │ POST /review
                         ▼
                ┌──────────────────┐
                │  Code Review API │
                │                  │
                │       Go         │
                └────────┬─────────┘
                         │
                         │ Prompt
                         ▼
                ┌──────────────────┐
                │      Ollama      │
                │                  │
                │ qwen2.5-coder    │
                │      1.5b        │
                └────────┬─────────┘
                         │
                         │ Generated response
                         ▼
                ┌──────────────────┐
                │   Review Result  │
                └──────────────────┘
```

The application uses Ollama as a local LLM runtime, avoiding the need to send source code to an external AI API.

## How It Works

The review flow consists of the following steps:

```text
1. Client sends source code
          │
          ▼
2. HTTP handler validates JSON
          │
          ▼
3. Application builds a review prompt
          │
          ▼
4. Prompt is sent to Ollama
          │
          ▼
5. Local LLM generates the response
          │
          ▼
6. API returns the review as JSON
```

The application currently specializes the prompt for Go code and asks the model to either suggest improvements or generate a simple unit test.

## API

### `POST /review`

Receives source code for analysis.

#### Request

```json
{
  "code": "func sum(a int, b int) int { return a + b }",
  "language": "go"
}
```

#### Response

```json
{
  "review": "..."
}
```

The generated review is returned as a JSON response.

## Prompt Engineering

The application builds a contextual prompt before sending the request to the model.

The current prompt instructs the model to behave as a Go code-review specialist and to:

- Analyze the provided code
- Suggest improvements
- Generate a simple unit test when appropriate

This provides a simple foundation for experimenting with prompt engineering applied to software development.

## Local AI

The project uses [Ollama](https://ollama.com/) as the local LLM runtime.

The configured model is:

```text
qwen2.5-coder:1.5b
```

Using a locally hosted model provides several advantages for experimentation:

- No external AI API is required
- Source code remains in the local environment
- No API key is required for the model
- Development can be performed offline after the model is available
- Different models can be evaluated locally

## Technology Stack

| Technology | Purpose |
|---|---|
| Go | Backend/API implementation |
| Go `net/http` | HTTP server |
| Ollama | Local LLM runtime |
| Qwen2.5-Coder 1.5B | Code-oriented language model |
| JSON | API request/response format |

## Project Structure

The project intentionally has a minimal structure:

```text
codereview-agent/
│
├── main.go
├── go.mod
├── go.sum
└── README.md
```

### `main.go`

Contains:

- HTTP API
- Request model
- Ollama integration
- Prompt construction
- Review handler
- Application startup

The minimal structure keeps the experiment focused on the interaction between the HTTP API and the local LLM.

## Getting Started

### Requirements

- Go 1.26+
- Ollama

### Clone

```bash
git clone https://github.com/clodoaldomarques/codereview-agent.git
cd codereview-agent
```

### Install dependencies

```bash
go mod download
```

### Install and start Ollama

Install Ollama according to your operating system and make sure the Ollama service is running.

Pull the configured model:

```bash
ollama pull qwen2.5-coder:1.5b
```

### Run the application

```bash
go run main.go
```

The API starts on:

```text
http://localhost:8080
```

## Testing the API

Using `curl`:

```bash
curl -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{
    "code": "func sum(a int, b int) int { return a + b }",
    "language": "go"
  }'
```

The service will send the code to the local Ollama model and return the generated review.

## Design Decisions

### Why Ollama?

Ollama provides a simple way to run LLMs locally.

For a code-review experiment, local inference also allows source code to remain inside the development environment.

### Why Go?

The project was implemented in Go to explore how a Go backend can integrate directly with a local LLM runtime.

The application uses the Ollama Go client rather than communicating with the runtime through manually constructed HTTP requests.

### Why a simple HTTP API?

Using an HTTP interface makes the experiment easy to integrate with other developer tools in the future.

Potential clients could include:

- IDE extensions
- CLI tools
- Git hooks
- CI/CD pipelines
- GitHub integrations
- Pull Request automation

## Current Limitations

This project is intentionally small and experimental.

The current implementation:

- Focuses on Go code review
- Uses a single configured model
- Does not persist reviews
- Does not integrate directly with GitHub Pull Requests
- Does not perform static analysis independently
- Does not maintain repository-level context
- Does not provide authentication
- Does not provide streaming responses
- Does not orchestrate multiple AI agents

These limitations provide opportunities for future experimentation.

## Future Improvements

Possible next steps include:

### Repository-aware reviews

Instead of reviewing isolated code snippets, provide repository context to the model.

```text
Repository
    │
    ├── Source files
    ├── Tests
    ├── Configuration
    └── Documentation
             │
             ▼
        Context Builder
             │
             ▼
            LLM
```

### GitHub Integration

Automatically review Pull Requests and publish the generated feedback as PR comments.

### Static Analysis

Combine deterministic static analysis with LLM-based analysis:

```text
             Source Code
                  │
          ┌───────┴────────┐
          │                │
          ▼                ▼
   Static Analysis        LLM
          │                │
          └───────┬────────┘
                  ▼
           Review Aggregator
                  │
                  ▼
            Final Report
```

### Structured Findings

Instead of returning free-form text, return structured findings containing:

- Severity
- File
- Line
- Category
- Description
- Suggested fix

### Automated Tests

Add automated tests for:

- HTTP handlers
- Request validation
- Prompt construction
- Ollama integration
- Error handling

### Model Configuration

Allow the model to be selected through configuration rather than hard-coding the model name.

## Engineering Concepts

This project demonstrates practical experimentation with:

- Go backend development
- HTTP APIs
- Large Language Models
- Generative AI
- Local AI inference
- Prompt engineering
- AI-assisted software development
- API integration
- Developer tooling

## Project Status

This project is an experimental proof of concept created to explore the integration of Go backend services with locally hosted Large Language Models.

It is intentionally kept small so that new AI-assisted software engineering concepts can be explored incrementally.

## Author

**Clodoaldo Marques**

Backend Software Engineer focused on Go, Microservices, Distributed Systems, Cloud-Native architectures and emerging AI-assisted development technologies.

- GitHub: https://github.com/clodoaldomarques
- LinkedIn: https://www.linkedin.com/in/clodoaldomarques/