Simple CLI tool for ordering beers written in go

This does not order beers in anyway and only serves as a basic example for understanding go and the cli

Makes use of the urfave/cli package

# To test run:

git clone https://github.com/ozombo/gocli.git

cd gocli

go build main.go

### using aliases
./main p for peanuts and beer
./main s for suya and beer
./main ps for pepper soup and beer

### using name
./main peanuts for peanuts and beer
./main suya for suya and beer
./main peppersoup for pepper soup and beer

### help and info
./main --help
./main --version