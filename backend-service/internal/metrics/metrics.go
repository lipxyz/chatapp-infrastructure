// metrics/metrics.go
package metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"time"
)

var (
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Длительность HTTP-запросов",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"}, // добавил статус как label
	)

	RedisRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_requests_total",
			Help: "Общее количество запросов к Redis",
		},
	)
	PostgresRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "postgres_requests_total",
			Help: "Общее количество запросов к Postgres",
		},
	)
)

var (
	// Метрики для логинов и регистрации
	LoginSuccessTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "login_success_total",
			Help: "Количество успешных логинов",
		},
	)
	LoginFailTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "login_fail_total",
			Help: "Количество неудачных логинов",
		},
	)
	LoginDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "login_duration_seconds",
			Help:    "Время обработки логина",
			Buckets: prometheus.DefBuckets,
		},
	)

	RegisterTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "register_total",
			Help: "Количество регистраций",
		},
	)
	RegisterDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "register_duration_seconds",
			Help:    "Время обработки регистрации",
			Buckets: prometheus.DefBuckets,
		},
	)
)
var (
	// Метрики для твитов
	TweetCreateTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tweet_create_total",
			Help: "Общее количество созданных твитов",
		},
	)
	TweetLikeTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tweet_like_total",
			Help: "Общее количество лайков твитов",
		},
	)
	TweetSlurmTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tweet_slurm_total",
			Help: "Общее количество удаленных твитов (slurm)",
		},
	)
	TweetShareTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tweet_share_total",
			Help: "Общее количество репостов твитов",
		},
	)
	TweetDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "tweet_duration_seconds",
			Help:    "Время обработки операции с твитом",
			Buckets: prometheus.DefBuckets,
		},
	)
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Общее количество HTTP-запросов",
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	// Регистрируем метрики в Prometheus
	prometheus.MustRegister(
		HttpRequestsTotal,
		LoginSuccessTotal,
		LoginFailTotal,
		LoginDuration,
		RegisterTotal,
		RegisterDuration,
		TweetCreateTotal,
		TweetLikeTotal,
		TweetSlurmTotal,
		TweetShareTotal,
		TweetDuration,
		HttpRequestDuration,
		RedisRequestsTotal,
		PostgresRequestsTotal)
}
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		if strings.HasPrefix(path, "/api/tweets") {
			RedisRequestsTotal.Inc()
		}
		if strings.HasPrefix(path, "/api/register") ||
			strings.HasPrefix(path, "/api/login") ||
			strings.HasPrefix(path, "/api/users") {
			PostgresRequestsTotal.Inc()
		}
		method := c.Request.Method
		status := fmt.Sprintf("%d", c.Writer.Status())
		duration := time.Since(start).Seconds()

		HttpRequestsTotal.WithLabelValues(path, method, status).Inc()
		HttpRequestDuration.WithLabelValues(path, method, status).Observe(duration)

	}
}
