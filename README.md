# Weather Service

> A Go-based HTTP server that provides real-time weather forecasts using the National Weather Service API.

## 📋 Project Overview

This is a coding exercise submission for the Weather Service Assignment. The application is a simple HTTP server that accepts geographic coordinates and returns current weather information for that location.

## ✨ Features

- **Coordinate Input**: Accept latitude and longitude coordinates via a user-friendly web form
- **Weather Forecasting**: Fetch real-time weather data from the National Weather Service API
- **Temperature Characterization**: Classify temperature as "hot", "cold", or "moderate"
- **Time-Based Forecasts**: Return weather for the current time of day (day or night)
- **Responsive UI**: Built with Bootstrap 5 for mobile-friendly design

## 🎯 Requirements

The server exposes an endpoint that:

1. ✅ Accepts latitude and longitude coordinates
2. ✅ Returns the short forecast for that area (e.g., "Partly Cloudy")
3. ✅ Returns temperature characterization ("hot", "cold", or "moderate")
4. ✅ Uses the [National Weather Service API](https://www.weather.gov/documentation/services-web-api) as the data source

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- macOS/Linux/Windows with terminal access

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/Weather_service.git
cd Weather_service
```

2. Run the server:
```bash
go run main.go
```

3. Open your browser and navigate to:
```
http://localhost:8080
```

## 📝 Usage

1. Enter latitude and longitude coordinates in the form
   - Latitude: -90 to 90
   - Longitude: -180 to 180

2. Click "Submit" to fetch weather data

3. View the current weather forecast and temperature characterization

## 🏗️ Project Structure

```
Weather_service/
├── main.go                 # Main server and handlers
├── models.go              # Data structures for API responses
├── page/
│   ├── landingPage.html   # Input form UI
│   └── responsePage.html  # Weather forecast display
└── README.md
```

## 🔌 API Endpoints

### GET `/`
Serves the landing page with the coordinate input form.

### POST `/submit`
Processes form submission with latitude and longitude.
- **Parameters**: latitude, longitude
- **Returns**: Weather forecast and temperature characterization

## 🌡️ Temperature Classification

- **Hot**: > 80°F
- **Moderate**: 50°F - 80°F
- **Cold**: < 50°F

*(Adjust thresholds in code as needed)*

## 🛠️ Technologies Used

- **Language**: Go
- **Framework**: Standard library (net/http)
- **Frontend**: HTML5, Bootstrap 5
- **API**: National Weather Service (NWS)

## 📚 API Data Source

This application uses the **National Weather Service API** which is free and publicly available:
- Points API: `https://api.weather.gov/points/{latitude},{longitude}`
- Forecast API: Uses the forecast URL returned from Points API

## ⚙️ Configuration

The server is configured to listen on **port 8080**. To change this, modify the port in `main.go`:

```go
http.ListenAndServe(":8080", nil)  // Change 8080 to your desired port
```

## 🚨 Limitations

- **US Only**: Currently limited to US coordinates (uses NWS which covers only USA)
- **Fahrenheit**: Temperature returned in Fahrenheit
- **Current Time**: Returns forecast for the current time of day only

## 📄 License

This project is provided as-is for educational purposes.

## 🤝 Contributing

Feel free to fork and submit pull requests for any improvements!


**Last Updated**: February 5, 2026