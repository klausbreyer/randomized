# Randomized

This Go program generates a random list of names or selects a single name based on the specified route and displays them in HTML. It’s ideal for workshops, demonstrations, random selections in group activities, and rotating daily tasks where everyone sees the same result for the day.

## Prerequisites

Ensure that Go is installed on your system.

## Installation and Execution

1. Clone or download this repository.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to start the server:

   ```sh
   go run main.go
   ```

4. Open your web browser and go to any of the example URLs to see the application in action.

## Usage

The application provides four main routes for randomizing or picking names:

### 1. `/shuffle/`

- Shuffles the list of names each time the URL is accessed.
- Example:
  ```plaintext
  http://localhost:8080/shuffle/steve,tim,johnny
  ```
- The names will be randomly shuffled and displayed in an HTML list. Each refresh results in a different order.

### 2. `/pick/`

- Picks a single name randomly from the list each time the URL is accessed.
- Example:
  ```plaintext
  http://localhost:8080/pick/steve,tim,johnny
  ```
- A random name is selected from the list and displayed. Each refresh may result in a different name.

### 3. `/pick-today/`

- Picks a single name but uses today's date as the seed, ensuring consistent results throughout the day.
- Example:
  ```plaintext
  http://localhost:8080/pick-today/steve,tim,johnny
  ```
- The same name will be selected for the entire day. The result changes only once per day.

### 4. `/shuffle-today/`

- Shuffles the list of names using today's date as the seed, producing the same order throughout the day.
- Example:
  ```plaintext
  http://localhost:8080/shuffle-today/steve,tim,johnny
  ```
- The order of names will remain the same for the entire day, refreshing only when a new day starts.

## Example URLs

- Shuffle names: `http://localhost:8080/shuffle/klaus,linus,jonas,julia`
- Shuffle names with today's seed: `http://localhost:8080/shuffle-today/klaus,linus,jonas,julia`
- Pick a name: `http://localhost:8080/pick/klaus,linus,jonas,julia`
- Pick a name with today's seed: `http://localhost:8080/pick-today/klaus,linus,jonas,julia`

## Live Demo

Try the application live at: [https://randomized.fly.dev/](https://randomized.fly.dev/)

## Customization

### Font Size and Family

The font size is set to 20pt, and the font family includes a series of modern system fonts:

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
