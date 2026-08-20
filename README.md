# venancio

CLI for [Drogaria Venancio](https://www.drogariavenancio.com.br), a pharmacy chain in Rio
de Janeiro. Built for humans and AI agents: data to stdout, hints to stderr, stable exit
codes, structured output.

Built on [vtexkit](https://github.com/voska/vtexkit), a shared library for Brazilian VTEX
storefronts.

## Install

```sh
brew install voska/tap/venancio
```

or `make build` for `bin/venancio`.

## Use

Start with `doctor`. It answers "can I order right now?" and prints the exact fix for
anything in the way.

```sh
venancio doctor
venancio auth login --email you@example.com
venancio search dipirona --limit 5
```

```
$ venancio search dipirona --limit 3
80237      Dipirona 1g Prati Donaduzzi 10 Comprimidos           R$9,32  un
70222      Dipirona 500mg Prati Donaduzzi Genérico 10 Comp…     R$3,99  un
6190       Dipirona 500mg Ems Genérico 10 Comprimidos           R$5,19  un
```

### Ordering

```sh
venancio cart add 80237 --qty 1
venancio delivery windows
venancio checkout --window 0            # preview — places nothing
venancio checkout --window 0 --confirm  # places the order
```

### Subscriptions

Venancio has VTEX Subscriptions enabled, which is how recurring medication is delivered.

```sh
venancio subs                  # status, frequency, next delivery
venancio subs <id>             # one subscription's schedule and items
venancio subs pause <id>
venancio subs resume <id>
venancio subs skip <id>        # skip the next delivery only
venancio subs unskip <id>
```

There is deliberately no `cancel`. VTEX has no transition out of `CANCELED`, so a mistyped
ID would destroy a subscription with no way back — cancel on the website instead.

Add `--json` to any command for agent-readable output, and see `venancio exit-codes` for
the exit code contract.

## A note on what this is

A pharmacy CLI orders pharmacy products. It does not know what is safe to take, what
interacts with what, or what needs a prescription — Brazilian law requires a prescription
for controlled and many non-OTC medicines, and this tool does not check or upload one.
Use it to reorder things you already have a decision about.

## What lives here

Only the store descriptor. Everything else — VTEX client, auth, search, cart, checkout,
subscriptions, output modes, exit codes — is in vtexkit.

```
store.go          # The Venancio descriptor — 3 fields
cmd/venancio/     # Entry point
```

## License

MIT
