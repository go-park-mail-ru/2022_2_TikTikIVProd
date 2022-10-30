package middleware

import (
	"net/http"

	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	authUC authUsecase.UseCaseI
}

func NewMiddleware(au authUsecase.UseCaseI) *Middleware {
	return &Middleware{authUC: au}
}

func (m *Middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_token")
		if err == http.ErrNoCookie {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		_, err = m.authUC.Auth(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return next(c)
	}
}

// func (m *Middleware) AccessLog(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		start := time.Now()
// 		logger.Logrus.Logger.SetFormatter(&logrus.JSONFormatter{
// 			TimestampFormat:   "",
// 			DisableTimestamp:  false,
// 			DisableHTMLEscape: false,
// 			DataKey:           "now",
// 			FieldMap:          nil,
// 			CallerPrettyfier:  nil,
// 			PrettyPrint:       true,
// 		})
// 		logger.Logrus.WithFields(logrus.Fields{
// 			"method":      ctx.Request().Method,
// 			"remote_addr": ctx.Request().RemoteAddr,
// 			"work_time":   time.Since(start),
// 		}).Debug(ctx.Request().URL.Path)
// 		return next(ctx)
// 	}
// }






// type (
// 	Stats struct {
// 		Uptime       time.Time      `json:"uptime"`
// 		RequestCount uint64         `json:"requestCount"`
// 		Statuses     map[string]int `json:"statuses"`
// 		mutex        sync.RWMutex
// 	}
// )

// func NewStats() *Stats {
// 	return &Stats{
// 		Uptime:   time.Now(),
// 		Statuses: map[string]int{},
// 	}
// }

// // Process is the middleware function.
// func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if err := next(c); err != nil {
// 			c.Error(err)
// 		}
// 		s.mutex.Lock()
// 		defer s.mutex.Unlock()
// 		s.RequestCount++
// 		status := strconv.Itoa(c.Response().Status)
// 		s.Statuses[status]++
// 		return nil
// 	}
// }

// // Handle is the endpoint to get stats.
// func (s *Stats) Handle(c echo.Context) error {
// 	s.mutex.RLock()
// 	defer s.mutex.RUnlock()
// 	return c.JSON(http.StatusOK, s)
// }

// // ServerHeader middleware adds a `Server` header to the response.
// func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
// 		return next(c)
// 	}
// }

// func main() {
// 	e := echo.New()

// 	// Debug mode
// 	e.Debug = true

// 	//-------------------
// 	// Custom middleware
// 	//-------------------
// 	// Stats
// 	s := NewStats()
// 	e.Use(s.Process)
// 	e.GET("/stats", s.Handle) // Endpoint to get stats

// 	// Server header
// 	e.Use(ServerHeader)

// 	// Handler
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	// Start server
// 	e.Logger.Fatal(e.Start(":1323"))
// }