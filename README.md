# Razorpay MCP Server (Official)

The Razorpay MCP Server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction) server that provides seamless integration with Razorpay APIs, enabling advanced payment processing capabilities for developers and AI tools.

## 🚀 Quick Deploy

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/razorpay/razorpay-mcp-server)

## 🌟 Features

- **🚀 HTTP Transport** - Default transport for web applications and API integrations
- **📡 stdio Transport** - Legacy support for command-line tools  
- **🔧 Multiple Toolsets** - Orders, Payments, Payment Links, Payouts, QR Codes, Refunds, Settlements
- **🔒 Read-only Mode** - Safe mode for query-only operations
- **🌐 Production Ready** - Optimized for cloud deployment with Docker support

## 📖 Quick Start

### Local Development

```bash
# Clone and build
git clone https://github.com/razorpay/razorpay-mcp-server
cd razorpay-mcp-server
go build -o razorpay-mcp-server ./cmd/razorpay-mcp-server

# Run HTTP server (default)
./razorpay-mcp-server --key YOUR_KEY --secret YOUR_SECRET --address :8080

# Run stdio server (legacy)
./razorpay-mcp-server stdio --key YOUR_KEY --secret YOUR_SECRET
```

### Production Deployment

See [DEPLOYMENT.md](./DEPLOYMENT.md) for detailed deployment instructions on Render, AWS, and other platforms.

### MCP Client Integration

**HTTP Transport (Recommended):**
```json
{
  "razorpay-mcp": {
    "url": "https://your-server.onrender.com/mcp"
  }
}
```

**stdio Transport (Local):**
```json
{
  "razorpay-mcp": {
    "command": "./razorpay-mcp-server",
    "args": ["stdio"],
    "env": {
      "RAZORPAY_API_KEY": "your_key",
      "RAZORPAY_API_SECRET": "your_secret"
    }
  }
}
```

## Available Tools

Currently, the Razorpay MCP Server provides the following tools:

| Tool                                 | Description                                            | API
|:-------------------------------------|:-------------------------------------------------------|:-----------------------------------
| `capture_payment`                    | Change the payment status from authorized to captured. | [Payment](https://razorpay.com/docs/api/payments/capture)
| `fetch_payment`                      | Fetch payment details with ID                          | [Payment](https://razorpay.com/docs/api/payments/fetch-with-id)
| `fetch_payment_card_details`         | Fetch card details used for a payment                  | [Payment](https://razorpay.com/docs/api/payments/fetch-payment-expanded-card)
| `fetch_all_payments`                 | Fetch all payments with filtering and pagination       | [Payment](https://razorpay.com/docs/api/payments/fetch-all-payments)
| `update_payment`                     | Update the notes field of a payment                    | [Payment](https://razorpay.com/docs/api/payments/update)|
| `create_payment_link`                | Creates a new payment link (standard)                  | [Payment Link](https://razorpay.com/docs/api/payments/payment-links/create-standard)
| `create_payment_link_upi`            | Creates a new UPI payment link                         | [Payment Link](https://razorpay.com/docs/api/payments/payment-links/create-upi)
| `fetch_all_payment_links`            | Fetch all the payment links                            | [Payment Link](https://razorpay.com/docs/api/payments/payment-links/fetch-all-standard)
| `fetch_payment_link`                 | Fetch details of a payment link                        | [Payment Link](https://razorpay.com/docs/api/payments/payment-links/fetch-id-standard/)
| `send_payment_link`                  | Send a payment link via SMS or email.                  | [Payment Link](https://razorpay.com/docs/api/payments/payment-links/resend)
| `update_payment_link`                | Updates a new standard payment link                    | [Payment Link](https://razorpay.com/docs/api/payments/payment-links/update-standard)
| `create_order`                       | Creates an order                                       | [Order](https://razorpay.com/docs/api/orders/create/)
| `fetch_order`                        | Fetch order with ID                                    | [Order](https://razorpay.com/docs/api/orders/fetch-with-id)
| `fetch_all_orders`                   | Fetch all orders                                       | [Order](https://razorpay.com/docs/api/orders/fetch-all)
| `update_order`                       | Update an order                                        | [Order](https://razorpay.com/docs/api/orders/update) 
| `fetch_order_payments`               | Fetch all payments for an order                        | [Order](https://razorpay.com/docs/api/orders/fetch-payments/)
| `create_refund`                      | Creates a refund                                       | [Refund](https://razorpay.com/docs/api/refunds/create-instant/)
| `fetch_refund`                       | Fetch refund details with ID                           | [Refund](https://razorpay.com/docs/api/refunds/fetch-with-id/)
| `fetch_all_refunds`                  | Fetch all refunds                                      | [Refund](https://razorpay.com/docs/api/refunds/fetch-all)
| `update_refund`                      | Update refund notes with ID                            | [Refund](https://razorpay.com/docs/api/refunds/update/)
| `fetch_multiple_refunds_for_payment` | Fetch multiple refunds for a payment                   | [Refund](https://razorpay.com/docs/api/refunds/fetch-multiple-refund-payment/)
| `fetch_specific_refund_for_payment`  | Fetch a specific refund for a payment                  | [Refund](https://razorpay.com/docs/api/refunds/fetch-specific-refund-payment/)
| `create_qr_code`                     | Creates a QR Code                                      | [QR Code](https://razorpay.com/docs/api/qr-codes/create/)
| `fetch_qr_code`                      | Fetch QR Code with ID                                  | [QR Code](https://razorpay.com/docs/api/qr-codes/fetch-with-id/)
| `fetch_all_qr_codes`                 | Fetch all QR Codes                                     | [QR Code](https://razorpay.com/docs/api/qr-codes/fetch-all/)
| `fetch_qr_codes_by_customer_id`      | Fetch QR Codes with Customer ID                        | [QR Code](https://razorpay.com/docs/api/qr-codes/fetch-customer-id/)
| `fetch_qr_codes_by_payment_id`       | Fetch QR Codes with Payment ID                         | [QR Code](https://razorpay.com/docs/api/qr-codes/fetch-payment-id/)
| `fetch_payments_for_qr_code`         | Fetch Payments for a QR Code                           | [QR Code](https://razorpay.com/docs/api/qr-codes/fetch-payments/)
| `close_qr_code`                      | Closes a QR Code                                       | [QR Code](https://razorpay.com/docs/api/qr-codes/close/)
| `fetch_all_settlements`              | Fetch all settlements                                  | [Settlement](https://razorpay.com/docs/api/settlements/fetch-all)
| `fetch_settlement_with_id`           | Fetch settlement details                               | [Settlement](https://razorpay.com/docs/api/settlements/fetch-with-id)
| `fetch_settlement_recon_details`     | Fetch settlement reconciliation report                 | [Settlement](https://razorpay.com/docs/api/settlements/fetch-recon)
| `create_instant_settlement`          | Create an instant settlement                           | [Settlement](https://razorpay.com/docs/api/settlements/instant/create)
| `fetch_all_instant_settlements`      | Fetch all instant settlements                          | [Settlement](https://razorpay.com/docs/api/settlements/instant/fetch-all)
| `fetch_instant_settlement_with_id`   | Fetch instant settlement with ID                       | [Settlement](https://razorpay.com/docs/api/settlements/instant/fetch-with-id)
| `fetch_all_payouts`                  | Fetch all payout details with A/c number               | [Payout](https://razorpay.com/docs/api/x/payouts/fetch-all/)
| `fetch_payout_by_id`                 | Fetch the payout details with payout ID                | [Payout](https://razorpay.com/docs/api/x/payouts/fetch-with-id)


## Use Cases 
- Workflow Automation: Automate your day to day workflow using Razorpay MCP Server.
- Agentic Applications: Building AI powered tools that interact with Razorpay's payment ecosystem using this Razorpay MCP server.

## Setup

### Prerequisites
- Docker
- Golang (Go)
- Git

To run the Razorpay MCP server, use one of the following methods:

### Using Public Docker Image (Recommended)

You can use the public Razorpay image directly. No need to build anything yourself - just copy-paste the configurations below and make sure Docker is already installed.

> **Note:** To use a specific version instead of the latest, replace `razorpay/mcp` with `razorpay/mcp:v1.0.0` (or your desired version tag) in the configurations below. Available tags can be found on [Docker Hub](https://hub.docker.com/r/razorpay/mcp/tags).

#### Usage with Claude Desktop

This will use the public razorpay image

Add the following to your `claude_desktop_config.json`:

```json
{
    "mcpServers": {
        "razorpay-mcp-server": {
            "command": "docker",
            "args": [
                "run",
                "--rm",
                "-i",
                "-e",
                "RAZORPAY_KEY_ID",
                "-e",
                "RAZORPAY_KEY_SECRET",
                "razorpay/mcp"
            ],
            "env": {
                "RAZORPAY_KEY_ID": "your_razorpay_key_id",
                "RAZORPAY_KEY_SECRET": "your_razorpay_key_secret"
            }
        }
    }
}
```
Please replace the `your_razorpay_key_id` and `your_razorpay_key_secret` with your keys.

- Learn about how to configure MCP servers in Claude desktop: [Link](https://modelcontextprotocol.io/quickstart/user)
- How to install Claude Desktop: [Link](https://claude.ai/download)

#### Usage with VS Code

Add the following to your VS Code settings (JSON):

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "razorpay_key_id",
        "description": "Razorpay Key ID",
        "password": false
      },
      {
        "type": "promptString",
        "id": "razorpay_key_secret",
        "description": "Razorpay Key Secret",
        "password": true
      }
    ],
    "servers": {
      "razorpay": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "-e",
          "RAZORPAY_KEY_ID",
          "-e",
          "RAZORPAY_KEY_SECRET",
          "razorpay/mcp"
        ],
        "env": {
          "RAZORPAY_KEY_ID": "${input:razorpay_key_id}",
          "RAZORPAY_KEY_SECRET": "${input:razorpay_key_secret}"
        }
      }
    }
  }
}
```

Learn more about MCP servers in VS Code's [agent mode documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

### Build from Docker (Alternative)

You need to clone the Github repo and build the image for Razorpay MCP Server using `docker`. Do make sure `docker` is installed and running in your system. 

```bash
# Run the server
git clone https://github.com/razorpay/razorpay-mcp-server.git
cd razorpay-mcp-server
docker build -t razorpay-mcp-server:latest .
```

Once the razorpay-mcp-server:latest docker image is built, you can replace the public image(`razorpay/mcp`) with it in the above configurations.

### Build from source

You can directly build from the source instead of using docker by following these steps:

```bash
# Clone the repository
git clone https://github.com/razorpay/razorpay-mcp-server.git
cd razorpay-mcp-server

# Build the binary
go build -o razorpay-mcp-server ./cmd/razorpay-mcp-server
```
Once the build is ready, you need to specify the path to the binary executable in the `command` option. Here's an example for VS Code settings:

```json
{
  "razorpay": {
    "command": "/path/to/razorpay-mcp-server",
    "args": ["stdio","--log-file=/path/to/rzp-mcp.log"],
    "env": {
      "RAZORPAY_KEY_ID": "<YOUR_ID>",
      "RAZORPAY_KEY_SECRET" : "<YOUR_SECRET>"
    }
  }
}
```

## Configuration

The server requires the following configuration:

- `RAZORPAY_KEY_ID`: Your Razorpay API key ID
- `RAZORPAY_KEY_SECRET`: Your Razorpay API key secret
- `LOG_FILE` (optional): Path to log file for server logs
- `TOOLSETS` (optional): Comma-separated list of toolsets to enable (default: "all")
- `READ_ONLY` (optional): Run server in read-only mode (default: false)

### Command Line Flags

The server supports the following command line flags:

- `--key` or `-k`: Your Razorpay API key ID
- `--secret` or `-s`: Your Razorpay API key secret
- `--log-file` or `-l`: Path to log file
- `--toolsets` or `-t`: Comma-separated list of toolsets to enable
- `--read-only`: Run server in read-only mode

## Debugging the Server

You can use the standard Go debugging tools to troubleshoot issues with the server. Log files can be specified using the `--log-file` flag (defaults to ./logs)

## License

This project is licensed under the terms of the MIT open source license. Please refer to [LICENSE](./LICENSE) for the full terms.
