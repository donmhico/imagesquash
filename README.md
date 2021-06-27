# What it does?

Listens to new images added in a directory and automatically create a compressed version of the image in
an output directory.

## Inspiration

As a developer, my everyday work involves taking screenshots and uploading them to the internet.
Typically my steps to do this would be:

1. Take a screenshot.

2. Upload the screenshot to an image compression website.

3. Download the compressed image.

4. Upload the compressed image to where I need it to be uploaded.

The steps are extremely simple and fast, and I don't see any need for change. But while reading "The Pragmatic Programmer" Chapter 3 - The Basic Tools, I was inspired to put more efforts in improving my tools and making it "my own".

ImageSquash does steps 2 and 3, which basically is the most time-consuming and requires more actions, for me.

## Why Go?

I built ImageSquash with [Go](https://golang.org/) simply because I want to learn Go.

## Word of caution

I'm building ImageSquash while learning Go so it's unavoidable to see code which lacks best practices.
If you see some, i'll be more than thankful if you can create an issue regarding it and point out what
I should have done.

ImageSquash will improve overtime as I learn more about Go.

## Usage

1. Clone the repo.
2. Edit some variables accordingly. See Variables below.
3. In your command line, navigate to the cloned repo.
4. You can either:

   **Build then run the program**

   ```bash
   go build
   ```

   Or just run it directly.

   ```bash
   go run .
   ```

## Variables

To make the program work, you need to edit some variables directly in [imagesquash.go](https://github.com/donmhico/imagesquash/blob/main/imagesquash.go#L14-L26).

**quality** - Between 1-100. The higher the value, the larger the compressed filesize but the better the quality.

**watchDir** - Full path of the directory to watch for new images.

**outputDir** - Full path where the compressed files are generated.

## Screenshot

![screenshot](https://user-images.githubusercontent.com/5747475/123534592-d83a7f80-d750-11eb-838d-83a508a2b435.jpg)

## Project Status

**ALPHA** - Basically it is tailored for my specific needs and workflow, at least for now.

## Limitations

1. Only tested on specific environments.

   - Windows 10 (Since this is my current OS. Although I don't see any reason why it won't work on other OS)
   - Go 1.16

2. Only works for JPG images.

## Special Thanks

[WebDevStudios](https://webdevstudios.com/) - My awesome company! I started ImageSquash during one of our [#5FTF](https://webdevstudios.com/about/how-wds-gives-back/).

The Pragmatic Programmer - I can't recommend this book enough to any programmer / devs.
