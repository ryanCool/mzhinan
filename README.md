# Michelin Restaurants Data Scraper

This Go program is designed to scrape data from Michelin-starred restaurants from the Michelin Guide's Chinese website and output the information into a CSV file. It uses Go's standard libraries for HTTP requests, CSV and JSON handling, and leverages the `goquery` library for HTML parsing and DOM traversal.

## Prerequisites

- Go (version 1.12 or higher is recommended)
- `goquery` library 

You can install `goquery` by running the following command:

```bash
go get github.com/PuerkitoBio/goquery
```

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/your_username/your_repository.git
    cd your_repository
    ```

2. Compile and run the program:

    ```bash
    go run main.go
    ```

    Or build an executable:

    ```bash
    go build main.go
    ```

    Then execute the binary:

    ```bash
    ./main
    ```

3. Once the program has finished running, a `result.csv` file will be created in the same directory. This file contains the scraped information about the restaurants.

## Data

The scraped data includes the following information about each restaurant:

- Chinese and English name
- Address
- Telephone number
- Website URL
- Michelin star rating
- Cuisine type
- Description in Chinese and English
- Latitude and longitude coordinates
- Opening hours
- Pricing information

## Configuration

The program uses a hard-coded API token and maps of city, cuisine, and star ratings. If these change on the Michelin website, you'll need to update them in the code.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For any issues, questions, or suggestions, please open an issue in the GitHub repository.
