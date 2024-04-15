# Distributary

Distributary is designed to simplify the restreaming process from OBS (Open Broadcaster Software) to multiple streaming platforms simultaneously. Whether you're a professional streamer, content creator, or just someone looking to expand your audience across various platforms, Distributary makes it easy to broadcast your content to multiple destinations with minimal hassle.

## Features

- **Multi-platform Restreaming**: Distribute your stream to multiple platforms simultaneously.
- **Easy Setup**: Simple configuration process to connect Distributary with your OBS and streaming accounts.
- **Customizable Settings**: Customize streaming settings, ensuring optimal performance and quality.
- **Automatic Failover**: In the event of a connection failure or platform outage, Distributary automatically switches to a backup platform to ensure uninterrupted streaming.
- **Simple Interface**: One-click interface designed for both beginners and experienced streamers alike.

## [Supported Platforms](SUPPORT.md)

## Getting Started

To get started with Distributary, follow these steps:

1. **Installation**: Download the [latest release](https://github.com/ahokinson/distributary/releases/tag/latest).
2. **Configuration**: Copy `distributary.yaml.example` to `distributary.yaml` and edit as necessary to add providers and edit stream options. Don't forget to edit your OBS stream settings to point at the stream host.
3. **Start Streaming**: Once configured, start Distributary and then start your stream in OBS. Distributary will handle the rest, broadcasting your content to all connected platforms simultaneously.

## Requirements

- [ffmpeg](https://ffmpeg.org/download.html) available in your OS path
- OBS (Open Broadcaster Software) or similar streaming software

## Support

For any questions, issues, or feedback, please [open an issue](https://github.com/ahokinson/distributary/issues) on GitHub.

## Contributing

Contributions are welcome! If you'd like to contribute to Distributary, please fork the repository and submit a pull request with your proposed changes.

## License

This project is licensed under the [GNU GPLv3 License](https://github.com/ahokinson/distributary/blob/main/LICENSE.md).
