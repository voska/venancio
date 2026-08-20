# Venancio CLI

Go CLI for ordering from Drogaria Venancio (`www.drogariavenancio.com.br`), a Rio de
Janeiro pharmacy chain.

All logic lives in the shared library **`github.com/voska/vtexkit`**, which also powers the
Frescatto, Zona Sul, and Mantiqueira CLIs. This repo holds only the store descriptor.

## Project Structure

```
store.go          # The Venancio descriptor — 3 fields
cmd/venancio/     # Entry point, ~15 lines
skills/venancio/  # Claude Code agent skill
```

Everything else — VTEX client, auth strategies, search, cart, checkout, subscriptions,
output modes, exit codes — is in vtexkit. Change behavior there, not here.

## Why the descriptor is only three fields

This is a stock VTEX store. The account name derives correctly from the base URL (unlike
Mantiqueira, whose account is `grupomantiqueira`), classic and access-key auth are both
enabled, its OAuth providers are ones VTEX ID drives generically, and both search backends
answer. Nothing needs declaring.

Subscriptions are enabled, so `venancio subs` works with no extra configuration.

## Safety

This orders medicine. Never assemble a cart of medication from a symptom description, and
never present product copy as medical advice. Prescription handling is not implemented —
the CLI cannot upload one.

## Build & test

`make build` `make test` `make lint` `make vet` `make ci`

## Commits

Conventional Commits (`feat:`, `fix:`, `chore:`, `docs:`, `test:`, `refactor:`).
