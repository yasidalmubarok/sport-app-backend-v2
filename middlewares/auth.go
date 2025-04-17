package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sport-app-backend/helper"
	"sport-app-backend/repositories"
)

type AuthService interface {
	AuthMiddleware() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	OperatorAuthorization() gin.HandlerFunc
	OwnerAuthorization() gin.HandlerFunc
}

type authService struct {
	db                  *gorm.DB
	userOwnerRepository repositories.UserOwnerRepository
}

func NewAuthService(db *gorm.DB, userOwnerRepository repositories.UserOwnerRepository) AuthService {
	return &authService{
		db:                  db,
		userOwnerRepository: userOwnerRepository,
	}
}

// Middleware untuk memvalidasi token JWT
func (s *authService) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			ctx.Abort()
			return
		}

		// Format token harus "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			ctx.Abort()
			return
		}

		// Parse token
		claims, err := helper.ParseJWT(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		// Simpan claims ke dalam context
		ctx.Set("user", claims)
		ctx.Next()
	}
}

func (s *authService) AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, helper.NewUnauthenticatedError("missing authorization header"))
			c.Abort()
			return
		}

		// Format token: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, helper.NewUnauthenticatedError("invalid token format"))
			c.Abort()
			return
		}

		// Parse token
		claims, err := helper.ParseJWT(tokenParts[1])
		if err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
			return
		}

		// Cek apakah role adalah "admin"
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, helper.NewUnathorizedError("access denied"))
			c.Abort()
			return
		}

		// Simpan user di context untuk akses selanjutnya
		c.Set("user", claims)
		c.Next()
	}
}

func (s *authService) OwnerAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, helper.NewUnauthenticatedError("missing authorization header"))
			c.Abort()
			return
		}

		// Format token: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, helper.NewUnauthenticatedError("invalid token format"))
			c.Abort()
			return
		}

		// Parse token
		claims, err := helper.ParseJWT(tokenParts[1])
		if err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
			return
		}

		// Cek apakah role adalah "owner"
		if claims.Role != "owner" {
			c.JSON(http.StatusForbidden, helper.NewUnathorizedError("access denied"))
			c.Abort()
			return
		}

		// Simpan user di context untuk akses selanjutnya
		c.Set("user", claims)
		c.Next()
	}
}

func (s *authService) OperatorAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, helper.NewUnauthenticatedError("missing authorization header"))
			c.Abort()
			return
		}

		// Format token: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, helper.NewUnauthenticatedError("invalid token format"))
			c.Abort()
			return
		}

		// Parse token
		claims, err := helper.ParseJWT(tokenParts[1])
		if err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
			return
		}

		// Cek apakah role adalah "owner"
		if claims.Role != "cashier" && claims.Role != "manager" {
			c.JSON(http.StatusForbidden, helper.NewUnathorizedError("access denied"))
			c.Abort()
			return
		}

		// Simpan user di context untuk akses selanjutnya
		c.Set("user", claims)
		c.Next()
	}
}
