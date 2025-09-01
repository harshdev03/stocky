# stocky

minimal cli tool for fetching Indian stock prices.

## preview

```
$ stocky "Tata Steel"

📊 Stock: Tata Steel
💵 Price: 150.25
💰 Open: 148.50
📈 High: 152.00
📉 Low: 147.75
🔄 Volume: 2345678
```

## installation

```
go install github.com/harshdev03/stocky@latest
```

## build from source

```
git clone https://github.com/harshdev03/stocky.git
cd stocky
go build -o stocky .
```

## usage

create a `.env` file with your api key:

```
RAPIDAPI_KEY=your_api_key_here
```

run directly:

```
stocky "Tata Steel"
stocky "Reliance"
stocky "TCS"
stocky "INFY"
```

## requirements

* requires an api key from [RapidAPI Indian Stock Exchange](https://rapidapi.com/).
* create `.env` file in the same directory as the executable.
* works during market hours (9:15 AM - 3:30 PM IST).

## notes

* supports major Indian stock names.
* simple and clean terminal output.
* lightweight and fast.