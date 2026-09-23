---
name: venancio
description: >-
  Order from Drogaria Venancio (www.drogariavenancio.com.br), a Rio de Janeiro
  pharmacy, using the `venancio` CLI. Use for pharmacy items, personal care,
  reordering medication the user already takes, and managing recurring
  subscriptions.
allowed-tools: Bash, Read
---

# venancio

Order from Drogaria Venancio using the `venancio` CLI.

## Read this before anything else

This CLI orders **medicine**. The rules below are not style preferences.

- **Never choose a medication from symptoms.** "I have a headache" is not a
  request for dipirona. Ask what the user already takes, or have them name the
  product. Reordering a known item is fine; selecting a drug is not.
- **Never present product copy as medical advice**, and never reason about
  dosage, interactions, or substituting one active ingredient for another.
- **Generic ≠ interchangeable, for this purpose.** A generic may well be the
  right call, but that is the user's or a pharmacist's decision to state, not
  one to make on their behalf while building a cart.
- **Prescriptions are not handled.** Brazilian law requires one for controlled
  and many non-OTC medicines. This CLI cannot upload or verify a prescription;
  if an item needs one, say so and stop.
- When unsure whether something is safe to reorder, stop and ask. There is no
  cost to asking and a real cost to guessing.

Personal care, hygiene, supplements, and household items carry none of this —
treat them like any other shopping.

## Always start here

```bash
venancio doctor
```

Exit 0 means ordering will work. Any other exit code means it will not. Each
failed line prints the exact fix. Do that fix, or report it to the user. Do not
retry the command that failed.

## The flow

```bash
venancio search dipirona --limit 5      # 1. find a SKU — first column is the SKU
venancio cart add 80237 --qty 1         # 2. add it
venancio delivery windows               # 3. pick a window number
venancio checkout --window 0            # 4. preview — places nothing
venancio checkout --window 0 --confirm  # 5. order
```

Between steps 4 and 5: **show the preview to the user and get explicit
approval.** `--confirm` spends real money.

Every result carries both `sku` and `productId`. Commands take the `sku` —
the two are separate sequences and the same number routinely appears in
both, naming two unrelated products.

```bash
venancio cart show
venancio cart update 0 --qty 2    # index from 'cart show'
venancio cart remove 0
venancio cart clear
```

**Payment** defaults to pix. `venancio checkout payments` lists what the store
accepts and any card saved on the account. For a card, add `--cvv 123`.

## Subscriptions

Recurring medication is delivered through VTEX Subscriptions.

```bash
venancio subs                  # status, frequency, next delivery
venancio subs <id>             # one subscription's schedule and items
venancio subs pause <id>       # pause indefinitely
venancio subs resume <id>
venancio subs skip <id>        # skip the next delivery only
venancio subs unskip <id>
```

**Be conservative here — this is someone's medication schedule.** Prefer `skip`
to `pause`: skipping drops one delivery and the schedule continues, while
pausing stops everything until someone remembers to resume. Never pause or skip
a subscription the user did not name, and confirm the ID with `venancio subs`
first.

**There is no `cancel`.** VTEX has no transition out of `CANCELED`, so the CLI
does not offer it. Tell the user to cancel on the website.

## Output for scripts and agents

```bash
venancio search dipirona --json
venancio search dipirona --json --select sku,name,price
venancio search dipirona --plain          # tab-separated
venancio subs --json
```

Prices in `--json` are integer centavos: `2499` is R$24,99. On a subscription
item the field is `priceAtSubscriptionDate` — the price locked in when the
subscription started, not today's price.

Data goes to stdout; progress and errors go to stderr.

## Exit codes

| Code | Meaning | What to do |
|---|---|---|
| 0 | success | continue |
| 2 | bad arguments | fix the command; do not retry it unchanged |
| 3 | empty result | tell the user nothing matched |
| 4 | login required | `venancio auth login --email <email>` |
| 5 | not found | the SKU or subscription ID is wrong; search again |
| 7, 8 | temporary | wait, then retry once |
| 9 | store rule refused | read the message; it names the rule |
| 10 | not set up | `venancio doctor` and follow the fixes |

Full table: `venancio exit-codes --json`

## Rules

- Never run `--confirm` without the user approving that exact cart and total.
- If a command fails twice the same way, stop and report it. Do not loop.
- `venancio doctor` diagnoses anything unexpected; its output names the fix.
