# redir

**redir** is a minimal URL redirect tool [inspired by Netlify redirects](https://docs.netlify.com/routing/redirects/)

## Usage

- Put all your URL redirects in a file named `_redirects` like so:

```
/ https://example.com
/foo https://bar.com
/baz https://test.com
```

- Run `go run .` to start the server locally at `http://localhost:3000`.
- Or by using docker:

```sh
docker run -p 3000:3000 -v _redirects:/app/_redirects -d hossainalhaidari/redir
```
