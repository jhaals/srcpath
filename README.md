# srcpath

srcpath provides go import style folder structure to your git clones

`gclone https://github.com/jhaals/yopass`
places the checkout in `~/src/github.com/jhaals/yopass`

## Install

Run `go build -o srcpath`. Put it in your \$PATH.

Add function .bashrc/.zshrc

```bash
function gclone {
    git clone $1 $(srcpath $1)
}
```

You can set your own base checkout folder with `-src`
