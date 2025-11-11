package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	registerWithWeight(90, func() gin.HandlerFunc {
		return Cors()
	})
}

func Cors() gin.HandlerFunc {
	// Configuration for development (ALLOW ALL ORIGINS - INSECURE FOR PRODUCTION)
	// return cors.New(cors.Config{
	//      AllowAllOrigins:  true, //  ONLY FOR DEVELOPMENT - INSECURE FOR PRODUCTION
	//      AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	//      AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	//      ExposeHeaders:    []string{"Content-Length"},
	//      AllowCredentials: true, // Careful with this in development, essential for cookies/auth
	//      MaxAge: 12 * time.Hour,
	// })

	// Configuration for production (SPECIFY ALLOWED ORIGINS)
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true // Allow ALL origins
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, //  Necessary if you use cookies or HTTP auth
		MaxAge:           12 * time.Hour,
	})

	// OR, if you need more dynamic control over origins:
	// return cors.New(cors.Config{
	//      AllowOriginFunc: func(origin string) bool {
	//              allowedOrigins := []string{"https://yourdomain.com", "https://anotheralloweddomain.com"}
	//              for _, allowedOrigin := range allowedOrigins {
	//                      if origin == allowedOrigin {
	//                              return true
	//                      }
	//              }
	//              return false
	//      },
	//      AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	//      AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	//      ExposeHeaders:    []string{"Content-Length"},
	//      AllowCredentials: true,
	//      MaxAge: 12 * time.Hour,
	// })
}
