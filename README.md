# genthumb

CLI to generate YouTube thumbnail. Being a blind person, generating image 
is simply a pain for me. I wanted a quick way to create ~decent YT thumbnail 
for my videos.

Not thinking that anyone would use this, but if you do, well please do not use 
my photos... left.png, right.png and center.png.

## Installing chrome on Ubuntu (Windows WSL2)

1. `wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb`
2. `sudo apt install ./google-chrome-stable_current_amd64.deb`

## Usage

Enter code in `code.sample` file.

```sh
$ go run . -dir left -code "$(<code.sample)" -title "title here"
```

The `-dir` accepts `left, right, center` as value and fits with the 
corresponding photos.