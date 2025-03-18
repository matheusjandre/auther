# auther

This is a simple authenticator system to help you put one more step on your authentication.

It works by:

  - Generating a hashed password with a configured .env key
  - Placing it in the clipboard
  - Erasing it from the clipboard after the configured lifespan

It was designed for Linux using Go, so it's necessary to have the X11 clipboard system installed.

## Configuring

