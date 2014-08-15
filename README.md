# Gosoon
## JSON serializer/deserializer for Golang

This project is not intended to be a supremely useful JSON library.
It is just a breakable toy by two aspiring Gophers.

### How I installed go:
On Ubuntu 14.04:

    sudo apt-get install golang-go

```
# Get gvm
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# Install go1.2
gvm install go1.2

# Default to use 1.2
gvm use 1.2

# Set GOPATH to this folder
export GOPATH=....
```

### Version reported by `go version`:
```
go1.2
```

On Ubuntu 14.04:

```
$ go version
go version go1.2.1 linux/amd64
```

### Extras
```
sudo apt-get install vim-syntax-go
```

Additionally, I had to do this on Ubuntu 14.04 for some reason:

```
cp /usr/share/vim/addons/syntax/go.vim ~/.vim/syntax/
```

I found it necessary to add this to my `~/.vimrc` on Linux

    set runtimepath+=$GOROOT/misc/vim
    autocmd BufNewFile,BufRead *.go         setfiletype go

Also, maybe I didn't need to do this, but I couldn't find a better way, so I:

    gvm use 1.2 # Except on Ubuntu 14.04, which is nice to omit!
    export GOPATH=$HOME/code/go
    PATH+=$PATH:$GOPATH/bin

The first two commands ensure that I have my `$GOPATH` set up for any Bash
session. The last one is so that I can run things like `ginkgo` from the
command-line.
