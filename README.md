# zeropad

Takes a string and an integer `X`, and returns the string with every whole number left-padded with zeros to `X` characters wide. Decimal fractions are left untouched.

## Examples

| Input | X | Output |
|-------|---|--------|
| `"James Bond 7"` | 3 | `"James Bond 007"` |
| `"PI=3.14"` | 2 | `"PI=03.14"` |
| `"It's 3:13pm"` | 2 | `"It's 03:13pm"` |
| `"It's 12:13pm"` | 2 | `"It's 12:13pm"` |
| `"99UR1337"` | 6 | `"000099UR001337"` |

## How to run

**Prerequisites:** Go 1.22 or later.

**Run the tests:**
```bash
go test ./...
```

**Run from the command line:**
```bash
go run ./main "James Bond 7" 3
# James Bond 007
```

## Approach

Single pass over the input runes using a `strings.Builder`:

1. If the current character is not a digit, write it and continue.
2. If it is a digit, collect the full consecutive digit run.
3. If that run is immediately preceded by a `.`, it's a decimal fraction — write it as-is.
4. Otherwise, left-pad the run to width `X` using `fmt.Sprintf` and write it.

## Complexity

| | Complexity | Notes |
|-|------------|-------|
| Time | O(n) | Single pass through the input; no backtracking |
| Space | O(n) | Output string grows proportional to input length plus any padding added |

No external libraries are used beyond the Go standard library (`fmt`, `strings`, `unicode`), so there is no third-party complexity overhead to account for.
