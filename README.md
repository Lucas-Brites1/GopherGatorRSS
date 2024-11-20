
# RSSGopher

RSSGopher is a simple RSS aggregator built in Golang, focused on learning. It is a Command Line Interface (CLI) program that allows you to manage RSS feeds, create users, aggregate content, and interact with data through commands.

## Features

- **`register`**: Registers a new user in the system.  
- **`login`**: Logs in as a registered user.  
- **`reset`**: Resets the database to its initial state.  
- **`users`**: Lists all users registered in the system.  
- **`agg <URL>`**: Aggregates feed content from a provided URL.  
- **`addfeed <URL>`**: Adds a feed to the database.  
- **`feeds`**: Lists all feeds available in the system.  
- **`follow <FEED_URL>`**: Follows an existing feed.  
- **`following`**: Shows the feeds followed by the current user.  
- **`unfollow <FEED_URL>`**: Unfollows a specific feed.  
- **`browse`**: Displays aggregated posts, with filtering support.  
- **`clear`**: Clears the terminal screen.  
- **`exit`**: Exits the program.

## Requirements

To test and run the program, you will need:

1. **PostgreSQL** - For the database.  
2. **Golang** - To build and run the program.

## How to Use

1. Clone the repository:  
   ```bash
   git clone https://github.com/Lucas-Brites1/RSSGopher.git
   cd RSSGopher
   ```

2. Configure the PostgreSQL database:  
   - Create a database and set up the connection file (e.g., `config.yaml`).

3. Build the program:  
   ```bash
   go build -o rssgopher main.go
   ```

4. Run the program:  
   ```bash
   ./rssgopher
   ```

5. Use the commands listed above directly in the CLI to interact with the system.

## Examples

### Register a new user
```bash
register <name>
```

### Log in
```bash
login <name>
```

### Add a feed
```bash
addfeed https://example.com/feed
```

### Follow a feed
```bash
follow <feed_url>
```

### Browse posts
```bash
browse <limit>
```

### Exit the program
```bash
exit
```

### Clear terminal
```bash
clear
```
