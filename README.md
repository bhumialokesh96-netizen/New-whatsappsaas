# WhatsApp Device Management & Automation SaaS

A production-ready Go backend for managing WhatsApp devices with automation capabilities, built with Whatsmeow, PostgreSQL (Supabase), and a responsive web UI.

## 🎯 Features

- **Device Pairing** - Generate 8-character pairing codes without QR scanning
- **Session Persistence** - Store WhatsApp sessions securely in PostgreSQL
- **Auto-Sender** - Send welcome messages to predefined numbers 5 seconds after device linking
- **Auto-Reply Bot** - Automatically reply to incoming private messages
- **Admin Dashboard** - Professional web interface for configuration and monitoring
- **REST API** - Complete API for all operations
- **Multi-Device Support** - Built on Whatsmeow's multi-device architecture

## 🛠️ Tech Stack

- **Backend**: Go 1.21+ with standard `net/http`
- **WhatsApp Library**: `go.mau.fi/whatsmeow` (Multi-device WhatsApp Web)
- **Database**: PostgreSQL (Supabase)
- **DB Driver**: `github.com/lib/pq`
- **Frontend**: Vanilla HTML, CSS, JavaScript

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL database (Supabase recommended)
- A WhatsApp account
- Modern web browser

## 🚀 Installation

### 1. Clone Repository

```bash
git clone https://github.com/bhumialokesh96-netizen/New-whatsappsaas.git
cd New-whatsappsaas
```

### 2. Download Dependencies

```bash
go mod download
go mod tidy
```

### 3. Configure Environment

Create a `.env` file in the root directory:

```bash
cp .env.example .env
```

Edit `.env` with your Supabase PostgreSQL connection string:

```
DATABASE_URL=postgresql://user:password@db.supabase.co:5432/postgres
PORT=8080
```

### 4. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📱 Usage

### Device Linking

1. Open `http://localhost:8080` in your browser
2. Enter your 10-digit phone number
3. Click "Get Pairing Code"
4. On your phone: WhatsApp → Settings → Linked Devices → Link a Device
5. Scan the displayed pairing code
6. Dashboard will automatically redirect once linked

### Configuration

Access the admin dashboard at `http://localhost:8080/admin`

#### Auto Sender Configuration
- Enable/disable automatic welcome messages
- Set target phone numbers
- Customize welcome message
- Messages are sent 5 seconds after device linking with 2-second delays between each

#### Auto Reply Configuration
- Enable/disable automatic replies
- Set auto-reply message
- Replies only to private messages (ignores groups and self-messages)

#### Quick Send
- Send manual messages to any phone number
- Requires active connection

## 📡 API Endpoints

### Pairing
- `GET /pair` - Get pairing code
- `GET /is-linked` - Check if device is linked

### Configuration
- `GET /api/config` - Get current configuration
- `POST /api/config` - Update configuration

### Status
- `GET /api/info` - Get connection status and device JID
- `GET /health` - Health check

### Messaging
- `POST /send` - Send a message

### Authentication
- `POST /logout` - Disconnect device

## 📁 File Structure

```
.
├── main.go              # Core backend application
├── index.html           # Device pairing UI
├── admin.html           # Admin dashboard
├── go.mod              # Go module definition
├── .env.example        # Environment template
├── .env                # Environment configuration (create from .env.example)
└── config.json         # Persistent configuration (auto-generated)
```

## 🏗️ Architecture

### Concurrency & Safety
- Thread-safe configuration management with `sync.RWMutex`
- Buffered channels for async event processing
- Graceful shutdown handling

### Event Handling
- `PairSuccess` - Triggers auto-sender after 5 seconds
- `Message` - Processes incoming messages for auto-reply
- `Connected/Disconnected` - Connection state management
- `KeepAliveTimeout` - Automatic reconnection

### Database Schema
- Whatsmeow sqlstore tables (auto-managed)
- Custom tables for config and status tracking

## 🔐 Security Considerations

- Environment variables for sensitive data
- Database connection pooling
- Timeout protections on HTTP handlers
- Session persistence with security
- No hardcoded credentials

## 🐛 Troubleshooting

### Issue: "DATABASE_URL environment variable not set"
**Solution**: Ensure `.env` file exists and contains `DATABASE_URL`

### Issue: "Failed to connect to WhatsApp"
**Solution**: 
- Check internet connection
- Verify Whatsmeow can reach WhatsApp servers
- Check WhatsApp account isn't banned

### Issue: Pairing code doesn't appear
**Solution**:
- Clear browser cache
- Try again with fresh pairing code
- Ensure WhatsApp app is using Multi-Device

### Issue: Auto-sender doesn't trigger
**Solution**:
- Verify configuration is saved and enabled
- Check phone numbers are in correct format (10 digits)
- Ensure device stays connected for 5 seconds after linking

### Issue: Auto-reply not working
**Solution**:
- Verify auto-reply is enabled in dashboard
- Ensure messages are private (not in groups)
- Check that received messages are text type

## 📝 Logging

All events are logged to stdout with timestamps:
- `[WA]` prefix for WhatsApp client logs
- Connection status changes
- Message send/receive events
- Configuration updates

## 📚 Next Steps

- Implement WebSocket support for real-time updates
- Add user authentication and multi-tenant support
- Implement message templates
- Add scheduled message sending
- Add contact management
- Implement message history

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## 📄 License

MIT License

## ⚠️ Disclaimer

- This project is for educational purposes
- WhatsApp ToS compliance is the user's responsibility
- Use responsibly and respect WhatsApp's policies
- Automated messaging may violate WhatsApp ToS

## 📞 Support

For issues, questions, or suggestions, please open an GitHub issue.

---

**Built with ❤️ for WhatsApp automation enthusiasts**