package routergin

import (
	"FaceIDApp/internal/config"
	_ "FaceIDApp/internal/docs"
	"FaceIDApp/internal/infrastructure/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type Answer struct {
	Code    int    `json:"code"  example:"200"`
	Message string `json:"message" example:"OK"`
}

// @title Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService https://swagger.io/terms/

// @contact.name API Support
// @contact.email stas@kuznec.team

// @BasePath /

type RouterGin struct {
	*gin.Engine
	h   *handler.Handlers
	cfg config.API
}

func NewRouterGin(h *handler.Handlers, cfg config.API) *RouterGin {
	gin.SetMode(gin.ReleaseMode)
	ret := gin.New()
	ret.Use(defaultStructLogger())
	ret.Use(gin.Recovery())

	router := &RouterGin{
		h:   h,
		cfg: cfg,
	}

	swag := ret.RouterGroup
	swag.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	open := ret.RouterGroup

	open.GET("/health", router.Health)

	private := ret.RouterGroup
	private.GET("/photo", router.GetPhotoInfo)
	router.Engine = ret

	return router
}

func defaultStructLogger() gin.HandlerFunc {
	return structureLogger(&log.Logger)
}

func structureLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Структура параметров логов GIN
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		param.Request = c.Request
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		var logEvent *zerolog.Event

		if c.Writer.Status() >= http.StatusBadRequest {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}
		clientRealIP := param.Request.Header.Get("X-Forwarded-For")
		if clientRealIP == "" {
			clientRealIP = param.ClientIP
		}
		scheme := param.Request.Header.Get("X-Forwarded-Proto")
		if scheme == "" {
			scheme = param.Request.URL.Scheme
		}

		userAgent := param.Request.UserAgent()
		logEvent.
			Str("client_ip", param.ClientIP).
			Str("client_real_ip", clientRealIP).
			Str("user_agent", userAgent).
			Str("scheme", scheme).
			Str("err", param.ErrorMessage).
			Str("method", param.Method).
			Str("host", param.Request.Host).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg("GIN")
	}
}

// Health godoc
// @Summary health check
// @Schemes
// @Description health check
// @Tags Public API
// @Accept json
// @Produce json
// @Success 200
// @Router /health [get]
// @Success 200 {object} Answer "OK response"
func (g RouterGin) Health(c *gin.Context) {
	answer := Answer{
		Code:    http.StatusOK,
		Message: "OK",
	}
	c.JSON(answer.Code, answer)
}

func (g RouterGin) GetPhotoInfo(c *gin.Context) {

}
