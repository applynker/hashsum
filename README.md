# hashsum

Install:

```bash
go install github.com/applynker/hashsum@latest
```

Usage:

```bash
hashsum [md5|sha1|sha256|sha512] "content to hash"
```
If no content provided, it will read input from stdin.

Example:

```bash
hashsum sha256 "hello world"
echo -n "hello world" | hashsum sha256
cat hello.txt | hashsum sha256
```