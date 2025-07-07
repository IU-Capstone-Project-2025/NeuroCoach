# AI Models Integration

This project uses OpenRouter API to work with various AI models, providing stable and cost-effective workout plan generation and chat functionality.

## Architecture

### OpenRouter Client
- Custom HTTP client in `internal/services/openrouter.go`
- Support for multiple models with automatic switching
- Built-in retry system with smart model switching

### AI Service
- Main service in `internal/services/ai.go`
- Features: workout plan generation, chat, plan regeneration
- MongoDB integration for data storage

## Available Models

The system automatically switches between the following free models:

1. `deepseek/deepseek-chat-v3-0324:free` - DeepSeek Chat (Best performance)
2. `openrouter/cypher-alpha:free` - Cypher Alpha model
3. `google/gemma-7b-it:free` - Google's Gemma model
4. `openchat/openchat-7b:free` - Optimized for conversations
5. `anthropic/claude-3-haiku:free` - Claude 3 Haiku
6. `meta-llama/llama-3-8b-instruct:free` - LLaMA 3 from Meta
7. `nousresearch/nous-hermes-2-mixtral:free` - Nous Hermes Mixtral
8. `microsoft/phi-3-mini-128k-instruct:free` - Microsoft Phi-3

## Setup

1. Get API key from [OpenRouter](https://openrouter.ai)
2. Create `.env` file:
```bash
OPENROUTER_KEY=sk-or-v1-your-key-here
```
3. Run the application

## Features

- **Automatic Switching**: When one model fails, system switches to the next
- **Retry Mechanism**: Smart retry system with multiple attempts
- **JSON Structuring**: Automatic processing of structured responses
- **Context Memory**: Chat history preservation for better understanding

## API Endpoints

- `POST /api/chat` - Chat with AI assistant
- `POST /api/generate-plan` - Generate workout plan
- `POST /api/regenerate-plan` - Update plan based on feedback
- `GET /api/chat/history` - Chat history