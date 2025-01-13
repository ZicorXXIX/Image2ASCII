# **Image2ASCII**

A powerful, lightweight CLI tool that converts images into stunning ASCII art. This tool supports PNG and JPEG formats, offering flexible options to customize the output. Transform your images into text-based art easily and efficiently!

---
## **Features**

- Convert images into ASCII art with customizable dimensions.
- Supports **JPEG** and **PNG** image formats.
- Adjustable width and height for ASCII output.
- Option to reverse the ASCII character set for inverted styles.
- Easy-to-use command-line interface.

---

## **Installation**

You can install the tool directly using Go:

```bash
go install github.com/ZicorXXIX/Image2ASCII@latest
```
---
## **Usage**
Run the tool from the terminal with the following syntax:
```bash
Image2ASCII --path <image-path> [options]
```
### **Options**
| Flag            | Description                                  | Default                     |
|------------------|----------------------------------------------|-----------------------------|
| `--width`        | Set the width of the ASCII art.              | 60                 |
| `--height`       | Set the height of the ASCII art.             | 60                  |
| `--reversed`     | Reverse the ASCII character set for output.  | `false`                    |
| `--colored`      | Output ASCII art in color (if terminal supports). | `true`                     |
| `--pixels`       | Define the set of characters to use for ASCII art. | `[ " .,:;i1tfLCG08@" ]`    |


TODOS: Get terminal width and height and adjust the image accordingly wihout using default width and height props
