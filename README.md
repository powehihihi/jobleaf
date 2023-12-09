# Jobleaf
### What is it
Ever struggled with boilerplate code when writing your resume? 

Jobleaf is an ultimate *low-latency high-level ai-powered low-code blockchainsawman maybe-encrypted almost-magic solution* for writting your CV!

### What is it 2
Jobleaf converts yaml to latex/html/md/groff templates, that's all.

### Examples!!!
This piece of art is generated from [lorem.yaml](examples/lorem.yaml):

![cool resume](https://powehihihi.github.io/lorem_resume.jpg)

### Getting started
First of all, download jobleaf (it's cli btw)
```
go install github.com/powehihihi/jobleaf@latest
```
Optional - you can create blank resume.yaml in your config directory:
```
jobleaf init
```

Now, RUN:
```
jobleaf run --input PATH/TO/YAML --output WHERE/PUT/TEX
```

### Using your own template
Jobleaf uses ```text/template``` for code generation. Adapting your template to my resume format shouldn't be hard. However, it's too boring. Please open an issue, i'll help you.

### Plans...
I really like this tiny project of mine, i'm planning to add:
- [ ] TUI wizard for filling resume.yaml
- [ ] [Mods](https://github.com/charmbracelet/mods) integration/example??
- [ ] Clear example of how to use golang ```text/template```
- [ ] HTML/markdown templates
- [ ] What if... TUI resume template

Also, code is a mess now. That's the first thing i'll fix.

### Special thanks
[main template](template/resume.tmpl.tex) is a fork of [Muhammadjon's resume](https://github.com/mrhakimov/resume). He did a great job!
