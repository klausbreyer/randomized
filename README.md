# Parameter Roulette

This Go program generates a random list of names and displays them in HTML. Good to be used in workshops etc.

## Prerequisites

Ensure that Go is installed on your system.

## Installation and Execution

1. Clone or download this repository.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to start the server:

   ```sh
   go run main.go
   ```

4. Open your web browser and go to `http://localhost:8080/klaus,manpreet,vishnu,marko,enzo,jasmin` (or any other list of names separated by commas or semicolons).

## Usage

- The server accepts a list of names separated by commas or semicolons.
- Example: `http://localhost:8080/klaus,manpreet,vishnu,marko,enzo,jasmin`
- The names will be randomly sorted and displayed in an HTML list.

## Example URL

```plaintext
http://localhost:8080/klaus,linus,jonas,julia
```

## Customization

### Font Size and Family

The font size is set to 20pt and the font family includes a series of modern system fonts:

```css
font-family: -apple-system, BlinkMacSystemFont, avenir next, avenir, segoe ui,
  helvetica neue, helvetica, Cantarell, Ubuntu, roboto, noto, arial, sans-serif;
```

### Colors

- Background color: `#1e1e1e`
- Text color: `#d4d4d4`
- Headline color: `#db2777`

These can be adjusted in the `style` block of the HTML template.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
