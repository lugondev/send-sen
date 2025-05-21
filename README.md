# Golang Send-Sen

A robust, modular notification service library for Go applications that provides unified interfaces for Email, SMS, and various notification channels. Built with extensibility in mind, it offers a clean architecture with adapters for popular service providers while maintaining flexibility for custom implementations.

## Installation

```bash
go get github.com/lugondev/send-sen
```

## Features

### Email Module ✉️
- Integrated providers:
  - ✅ SendGrid - Enterprise-grade email delivery
  - ✅ Brevo (formerly Sendinblue) - Comprehensive email marketing solution
  - ✅ Mock adapter (for testing) - Simplifies unit testing

### SMS Module 📱
- Integrated providers:
  - ✅ Twilio - Industry-standard SMS service
  - ✅ Brevo SMS - Cost-effective SMS solution
- Extensible adapter interface for custom providers

### Notification Module 🔔
- Integrated providers:
  - ✅ Telegram - Instant messaging platform integration
  - ✅ Mock logging (for testing) - Facilitates testing scenarios
- Flexible port interface for custom providers

## Architecture

The project follows a clean architecture pattern:
1. **Modules**: Contain the core business logic and define interfaces (ports)
2. **Adapters**: Implement the interfaces defined by modules to connect with external services
3. **Configuration**: Centralized configuration management using Viper
4. **Testing**: Comprehensive test suite with mock adapters for testing

## Configuration

### YAML Configuration
Create a `config.yaml` file:

## Testing
The project includes comprehensive testing support:

### Mock Adapters
- Mock adapters in `adapters/email/mock.go` for email testing
- Mock adapters in `adapters/sms/mock.go` for SMS testing
- Mock adapters in `adapters/notify/mock.go` for notification testing

## Code Style Guidelines

1. Follow Go's standard code style and conventions
2. Use meaningful variable and function names
3. Add appropriate comments for public functions and complex logic
4. Implement proper error handling and logging
5. Keep functions small and focused on a single responsibility
6. Write unit tests for new functionality

## Implementation Guidelines

### For new adapters
- Implement the appropriate interface from the modules package
- Add comprehensive logging
- Include proper error handling
- Add tests in the tests directory

### For bug fixes
- Identify the root cause
- Add tests that reproduce the issue
- Fix the issue
- Verify that the tests pass

### For refactoring
- Ensure all tests pass before and after changes
- Maintain backward compatibility when possible
- Update documentation if interfaces change

## Building and Running

The project is a library, so there's no need to build it separately. However, when testing changes:

```bash
# Verify that the code compiles
go build ./...

# Run tests
go test ./tests/...
```

## Provider Status

### Integrated Providers ✅
- Email Services:
  - SendGrid - Full support for transactional emails
  - Brevo - Complete email sending capabilities
- SMS Services:
  - Twilio - Full SMS functionality
  - Brevo SMS - Complete SMS support
- Notification Services:
  - Telegram - Complete bot integration

### Planned Integrations 🚀
#### Notification Services
- [ ] Slack integration
- [ ] Discord integration
- [ ] Microsoft Teams integration

#### Email Providers
- [ ] Mailgun support
- [ ] Mailchimp support
- [ ] Amazon SES integration

#### SMS & Push Notifications
- [ ] Firebase Cloud Messaging (FCM)
- [ ] Vonage (formerly Nexmo)
- [ ] MessageBird

## Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Areas for Contribution
- Implement planned integrations
- Improve documentation
- Add more test coverage
- Optimize existing code
- Report bugs
- Suggest new features

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.