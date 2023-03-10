package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	g "github.com/matoous/go-nanoid/v2"
)

func ValidateToken(tokenKey string) (interface{}, error) {
	token, err := jwt.Parse(tokenKey, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(configs.APP_SECRET_KEY), nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	return claims["name"], nil
}

func AuthorizationRequired(c *fiber.Ctx) error {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusUnauthorized
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if token == "" {
		r.Message = "Token is Required!"
		return c.Status(r.StatusCode).JSON(&r)
	}

	// Check Token On DB
	db := configs.Store
	var jwtToken models.JwtToken
	if err := db.Where("id=?", token).First(&jwtToken).Error; err != nil {
		r.Message = "Token is Invalid!"
		return c.Status(r.StatusCode).JSON(&r)
	}

	if jwtToken.ID == "" {
		r.Message = "Token is Invalid!"
		return c.Status(r.StatusCode).JSON(&r)
	}

	_, err := ValidateToken(jwtToken.Token)
	if err != nil {
		r.Message = err.Error()
		db.Delete(&jwtToken)
		return c.Status(r.StatusCode).JSON(&r)
	}
	return c.Next()
}

func CreateToken(user *models.User) models.AuthSession {
	db := configs.Store
	var obj models.AuthSession
	var admin models.Administrator
	db.Select("id").First(&admin, "user_id=?", &user.ID)
	if admin.ID != "" {
		obj.IsAdmin = true
	}
	obj.Header = "Authorization"
	obj.User = &user
	obj.JwtType = "Bearer"
	obj.JwtToken, _ = g.New(60)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = obj.JwtToken
	claims["name"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenKey, err := token.SignedString([]byte(configs.APP_SECRET_KEY))
	if err != nil {
		panic(err)
	}
	/// Insert Token Key to DB
	t := new(models.JwtToken)
	t.ID = obj.JwtToken
	t.UserID = &user.ID
	t.Token = tokenKey
	// Delete UserID before creating TokenID
	if err := db.Where("user_id=?", t.UserID).Delete(&models.JwtToken{}).Error; err != nil {
		panic(err)
	}

	if err := db.Create(&t).Error; err != nil {
		panic(err)
	}
	return obj
}

func GetJWTToken(c *fiber.Ctx) string {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusUnauthorized
	s := c.Get("Authorization")
	return strings.TrimPrefix(s, "Bearer ")
}

func GetUserID(c *fiber.Ctx) string {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusUnauthorized
	s := c.Get("Authorization")
	var jwt models.JwtToken
	if err := configs.Store.First(&jwt, "id", strings.TrimPrefix(s, "Bearer ")).Error; err != nil {
		panic(err)
	}
	return *jwt.UserID
}

func GetEmpID(c *fiber.Ctx) string {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusUnauthorized
	s := c.Get("Authorization")
	var jwt models.JwtToken
	if err := configs.Store.First(&jwt, "id", strings.TrimPrefix(s, "Bearer ")).Error; err != nil {
		panic(err)
	}

	var user models.User
	if err := configs.Store.First(&user, "id", jwt.UserID).Error; err != nil {
		panic(err)
	}
	return user.UserName
}

func SignOut(c *fiber.Ctx) string {
	var r models.Response
	r.At = time.Now()
	r.StatusCode = fiber.StatusUnauthorized
	s := c.Get("Authorization")
	if err := configs.Store.Delete(&models.JwtToken{}, "id", strings.TrimPrefix(s, "Bearer ")).Error; err != nil {
		panic(err)
	}
	return "?????????????????????????????????????????????????????????????????????"
}
