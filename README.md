# snack
snack is a static site generator (SSG) with a focus on modularity and
simplicity.

## Why?
You might ask yourself: Why do we need *yet another* generator for static
websites? If you look at the current landscape of static site generators the
question is justified. There are already quite sophisticated and popular static
site generators like [Hugo](https://github.com/gohugo/hugo),
[jekyll](https://github.com/jekyll/jekyll) and e.g.
[Gatsby](https://github.com/gatsbyjs/gatsby).

But, depending on your needs, there are sometimes a couple of limitations with
those site generators:

1) They are most often designed for (personal) blogs.
2) They are too complicated.
3) They force you to use a specific frontend technology.
4) They don't allow you to build custom pages.

This project aims to resolve the above-mentioned issues by providing a
frontend technology independent way of creating modular, static sites.

## Features

* modular pages: each page can have different sections
* global sections: sections like headers, and footers that should be on every page
* technology independent: use plain HTML, vanilla JS or React - snack doesn't care
* build your website based on JSON, YAML or from a file structure
* simplicity: there are pages and sections - that's it

## Usage
First, make sure that your theme can be found inside your root directory under
`/themes/<name>`. Then inside your root directory run the command:

``` bash
./snack -a json site.json
./snack -a yaml site.yml
./snack -a file src/
```

This will create a new `build` folder in your working directory.
Check out the `examples` folder in this repository for a quickstart example.
