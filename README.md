# Gosoon
## JSON serializer/deserializer for Golang

This project is not intended to be a supremely useful JSON library.
It is just a breakable toy by two aspiring Gophers.

### How I installed go:
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

### Extras
```
sudo apt-get install vim-syntax-go
```

I found it necessary to add this to my `~/.vimrc` on Linux

    set runtimepath+=$GOROOT/misc/vim
    autocmd BufNewFile,BufRead *.go         setfiletype go
