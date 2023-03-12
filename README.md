# snack
A simple static site generator to build modular pages.

## Usage
First, make sure that your theme can be found inside your root directory under
`/themes/<name>`. Then inside your root directory run the command:

``` bash
./snack -a json site.json
./snack -a yaml site.yml
./snack -a file src/
```

This will create a new `build` folder in your working directory.
