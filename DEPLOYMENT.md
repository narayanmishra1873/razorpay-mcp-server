# Deployment Guide for Render

This guide walks you through deploying the Razorpay MCP Server on Render with HTTP transport.

## Prerequisites

1. **Render Account** - Sign up at [render.com](https://render.com)
2. **GitHub Repository** - Fork or clone this repository
3. **Razorpay API Keys** - Get them from [Razorpay Dashboard](https://dashboard.razorpay.com/app/keys)

## Quick Deploy

### Option 1: One-Click Deploy (Recommended)

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/razorpay/razorpay-mcp-server)

### Option 2: Manual Setup

1. **Fork this repository** to your GitHub account

2. **Create a new Web Service** on Render:
   - Go to [Render Dashboard](https://dashboard.render.com)
   - Click "New +" → "Web Service"
   - Connect your GitHub repository

3. **Configure the service:**
   ```
   Name: razorpay-mcp-server
   Runtime: Docker
   Region: Oregon (or your preferred region)
   Branch: main
   Build Command: (leave empty - uses Dockerfile)
   Start Command: ./server --key $RAZORPAY_API_KEY --secret $RAZORPAY_API_SECRET --address :$PORT --endpoint-path /mcp
   ```

4. **Set Environment Variables:**
   ```
   RAZORPAY_API_KEY=your_razorpay_key_here
   RAZORPAY_API_SECRET=your_razorpay_secret_here
   TOOLSETS=orders,payments,payment_links,payouts,qr_codes,refunds,settlements
   READ_ONLY=false
   ENDPOINT_PATH=/mcp
   ```

5. **Deploy** - Click "Create Web Service"

## Environment Variables

### Required
- `RAZORPAY_API_KEY` - Your Razorpay API key (test or live)
- `RAZORPAY_API_SECRET` - Your Razorpay API secret (test or live)

### Optional
- `TOOLSETS` - Comma-separated list of toolsets to enable (default: all)
- `READ_ONLY` - Set to "true" for read-only mode (default: false)
- `ENDPOINT_PATH` - Custom endpoint path (default: "/mcp")
- `PORT` - Server port (automatically set by Render)

### Available Toolsets
- `orders` - Order management
- `payments` - Payment operations
- `payment_links` - Payment link creation and management
- `payouts` - Payout operations
- `qr_codes` - QR code generation
- `refunds` - Refund processing
- `settlements` - Settlement data

## Using the Deployed Server

Once deployed, your server will be available at:
```
https://your-service-name.onrender.com/mcp
```

### MCP Client Configuration

```json
{
  "razorpay-mcp": {
    "url": "https://your-service-name.onrender.com/mcp"
  }
}
```

### Testing the Deployment

1. **Health Check:**
   ```bash
   curl https://your-service-name.onrender.com/mcp
   ```

2. **Initialize MCP:**
   ```bash
   curl -X POST https://your-service-name.onrender.com/mcp \
     -H "Content-Type: application/json" \
     -d '{
       "jsonrpc": "2.0",
       "id": 1,
       "method": "initialize",
       "params": {
         "protocolVersion": "2024-11-05",
         "capabilities": {},
         "clientInfo": {
           "name": "test-client",
           "version": "1.0.0"
         }
       }
     }'
   ```

## Monitoring

Render provides built-in monitoring:

1. **Logs** - View real-time logs in the Render dashboard
2. **Metrics** - CPU, memory, and request metrics
3. **Health Checks** - Automatic health monitoring
4. **Alerts** - Email notifications for issues

## Scaling

### Free Tier
- 750 hours/month
- Automatic sleep after 15 minutes of inactivity
- Cold start on first request

### Paid Plans
- Always-on service
- Auto-scaling
- Custom domains
- SSL certificates

## Security Best Practices

1. **Use Environment Variables** - Never commit API keys to code
2. **Enable HTTPS** - Render provides SSL certificates automatically
3. **Use Read-Only Mode** - For query-only operations
4. **Monitor Access** - Check logs regularly
5. **Rotate Keys** - Update API keys periodically

## Troubleshooting

### Common Issues

**Service won't start:**
- Check environment variables are set correctly
- Verify API keys are valid
- Check build logs for errors

**Authentication errors:**
- Ensure API keys match your Razorpay account type (test/live)
- Verify keys have required permissions

**Connection timeouts:**
- Check if service is sleeping (free tier)
- Verify URL is correct
- Test with health check endpoint

### Debug Commands

```bash
# Check service status
curl -I https://your-service-name.onrender.com/mcp

# View logs (in Render dashboard)
# Go to your service → Logs tab

# Test with verbose curl
curl -v https://your-service-name.onrender.com/mcp
```

## Custom Domain (Paid Plans)

1. Go to your service settings
2. Add custom domain
3. Update DNS records as instructed
4. SSL certificate will be automatically provisioned

## Support

- [Render Documentation](https://render.com/docs)
- [Razorpay API Documentation](https://razorpay.com/docs/api/)
- [MCP Specification](https://modelcontextprotocol.io/)

## Cost Estimation

- **Free Tier**: $0/month (750 hours, sleeps after inactivity)
- **Starter**: $7/month (always-on, 0.5 CPU, 512MB RAM)
- **Standard**: $25/month (1 CPU, 2GB RAM)
- **Pro**: $85/month (2 CPU, 4GB RAM)
