# GoWeather
A CLI tool to provide the weather on a given location written in Go using the [Open Weather API](https://openweathermap.org/)

## Usage
To run the program follow the steps below:
```bash
  git clone https://github.com/1063551/GoWeather.git
  cd GoWeather/
```

To see the valid format of the arguments:
```bash
  go run main.go -h
    -location string
    	  a valid location (default "Barcelona")
```
And now run the program with your preferred location (in this case, Barcelona):
```bash
  go run main.go Barcelona
```
