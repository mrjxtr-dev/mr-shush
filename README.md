> ⚠️ This Readme is AI generated for now and might not contain most up to date and relevant information

# Mr Shush 🤫

**A Secure, Offline-First Password & Secrets Manager written in Go**

_Your secrets are safe with me... shhh._

---

## ✨ What is mr-shush?
**mr-shush** is a lightweight, offline-first password generator and secrets manager built with speed and simplicity in mind — all thanks to Go. Designed for developers, sysadmins, privacy-conscious users, and anyone who wants full control over their sensitive data without relying on the cloud.

It's fast. It's secure. It's quiet.

---

## 🚀 Features

- 🔐 **Secure local storage** with strong encryption
- 🧪 **Generate passwords** with customizable options
- 🧰 **Manage secrets** like environment variables, API keys, database creds, etc.
- 💻 **Offline CLI**: Simple commands, fast execution
- 🌐 **Self-hosted Web App** with modern UI built using templ and TailwindCSS (coming soon)
- 📦 **Export** secrets to `.env` files or shell scripts
- 🛡️ **No phoning home** — no internet required

---

## 🎯 Who is it for?

- Devs managing local `.env` files
- Tinfoil hat wearers who don't trust the cloud
- Teams wanting a private, internal secrets manager
- You — yes, you — the curious Go enthusiast

---

## 🛠️ Installation

Coming soon as a downloadable binary...  
But for now:

```bash
git clone https://github.com/mrjxtr/mr-shush.git
cd mr-shush
go build -o mr-shush .
./mr-shush
```

---

## 💡 Example CLI Commands

```bash
# General format of commands
mr-shush [command] [options] [arguments]

# Generate a strong password
mr-shush generate --length 10 --strong

# Generate a weak password
mr-shush generate --length 10 --weak

# Generate and store a password
mr-shush generate --length 10 --store --name github-password

# Store a password directly
mr-shush store --name github-password --password supersecret123!

# Store a new secret
mr-shush store --vault vault-name --key DB_PASSWORD --value supersecret123!

# Store multiple secrets at once
mr-shush store --vault vault-name --key DB_PASSWORD --value supersecret123!
mr-shush store --vault vault-name --key API_KEY --value abc123xyz
mr-shush store --vault vault-name --key AWS_SECRET --value aws-secret-key

# Store all secrets from a .env file
mr-shush store --vault vault-name --env-file path/to/.env

# List all vaults
mr-shush list

# List secrets in a specific vault (names only, not values 😉)
mr-shush list --vault vault-name

# Export secrets to .env format
mr-shush export --format env > path/to/.env

# Launch the local web UI (soon)
mr-shush start
```

---

## 🧠 How It Works

- All secrets are stored **locally**, encrypted using AES-256.
- A **master password** protects the entire vault.
- Data is saved in a local file-based database (SQLite).
- Nothing leaves your machine. Ever unless you want it to.

---

## 🏗️ Project Structure

```plaintext
mr-shush/
├── cmd/                # Application entry points
│   ├── app/            # Web application
│   └── cli/            # Command-line interface
├── internal/           # Private application code
│   ├── config/         # App configuration and flag parsing
│   ├── db/             # SQLite database operations for secrets
│   ├── generator/      # Password/API key generation logic
│   ├── middleware/     # HTTP middleware
│   ├── models/         # Data structures for secrets, settings
│   ├── server/         # Route handlers, static file server
│   ├── templates/      # UI templates (templ)
│   └── vault/          # Encryption, master password, locking
├── static/             # Web UI static assets
│   ├── css/            # TailwindCSS styles
│   ├── img/            # Images and icons
│   └── scripts/        # Frontend JavaScript
└── Makefile            # Development helpers
```

### Inside `internal/`

This is where the core functionality lives, private to the application:

- `db/`: Database interaction for storing/retrieving secrets using SQLite
- `vault/`: Encryption/decryption with master password protection
- `generator/`: Secure generation of passwords, API keys, and tokens
- `models/`: Data structures for secrets, user settings, vault entries
- `config/`: Application configuration, default paths, CLI flags
- `server/`: Web server with route handlers and middleware
- `middleware/`: HTTP middleware for authentication, logging, etc.
- `templates/`: UI templates built with the templ engine

### Inside `cmd/`

Clean separation between different application entry points:

- `app/`: Web interface entry point - initializes and runs the web server
- `cli/`: Command-line interface - parses arguments and calls internal functions

This separation ensures that the core functionality can be accessed through multiple interfaces while maintaining a single source of truth for business logic.

## 🛠️ Tech Stack (GoTTH)

- **Backend**: Go (1.24+)
- **Web UI**: templ templating engine with TailwindCSS
- **Frontend**: HTMX for interactive UI components without JavaScript
- **Development**: Air for hot-reloading
- **Database**: SQLite for local storage
- **Security**: AES-256 encryption

---

## 🔒 Security Notes

- Don't lose your master password. Seriously. We can't help you get it back.
- Always keep backups of your vault file in a safe place.
- mr-shush is offline-first by design — keep it that way for maximum privacy.

---

## 📦 Roadmap

- ✅ CLI with generate/store/export features
- 🔒 Master password & encryption
- 🖥️ Local-only web UI (with login)
- 🔍 Search & filter secrets
- 📤 Auto `.env` injection for dev environments
- ☁️ Optional local network sync, Cloud sync, or USB backup support

---

## 💻 Developer Quickstart

```bash
# Install development dependencies
make go-install-air        # Hot-reloading
make go-install-templ      # UI templating
make get-install-tailwindcss # CSS framework

# Build and run the project
make build                 # One-time build
make watch                 # Development mode with hot-reload

# UI development
make tailwind-watch        # Live CSS compilation
make templ-watch           # Live template compilation
```

---

## 🤝 Contributing

Got ideas? Found a bug? Want to add a feature?  
PRs are welcome! Just open an issue or fork and go wild.

Please review our [Code of Conduct](CODE_OF_CONDUCT.md) before contributing.
Security issues should be reported according to our [Security Policy](SECURITY.md).

---

## ☕ Support

This tool will always be free and open-source.  
If it made your life easier, consider [buying me a coffee](https://buymeacoffee.com/mrjxtr) or just giving the repo a ⭐️.

---

## 🧙‍♂️ License

MIT — do whatever you want, just don't blame me if you forget your master password and lock yourself out of your secrets 🫠

---

**Built with Go, caffeine, and a deep mistrust of cloud-based secret managers.**
