package vote

import (
	"9mookapook/vote/core"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

type (
	Admin struct {
		Admins map[string]int
	}
)

var (
	admin *Admin
)

func LoadAdmin() {
	if admin == nil {
		admin = &Admin{
			Admins: map[string]int{
				"Emily":    1,
				"Isabella": 1,
				"Alice":    1,
			},
		}
	}
}

func (a *Admin) IsAdmin(u string) bool {
	if _, ok := a.Admins[u]; ok {
		return true
	}
	return false
}

func Route() {
	app := core.APP()
	c := NewController()

	app.GET("/v1/item", c.GetAllItem, JWTAuthMiddleware)
	// Create Item
	app.POST("/v1/itemcreate", c.CreateitemVote, JWTAuthMiddleware)

	// Update Item
	app.PUT("/v1/itemcreate/:id", c.UpdateitemVote, JWTAuthMiddleware)
	// Remove Item

	app.DELETE("/v1/itemcreate/:id", c.Removeitem, JWTAuthMiddleware)

	// Vote By Item
	app.POST("/v1/itemvote/:id", c.itemVoteByID, JWTAuthMiddleware)

	// unVote By Item
	app.PUT("/v1/itemvote/:id", c.UnitemVoteByID, JWTAuthMiddleware)

	// Remove Data All and Vote member All
	app.PUT("/v1/itemclear", c.ClearALL, JWTAuthMiddleware)
	// Clear data by id

	app.PUT("/v1/itemclearbyid/:id", c.ClearbyItem, JWTAuthMiddleware)

	// Set Staus Open Close
	app.PUT("/v1/itemopenclose/:id", c.OpenCloseItem, JWTAuthMiddleware)

	// Login
	app.POST("/v1/login", c.LoginUser)

	// Export Data ALL
	app.GET("/v1/export", c.ExportItem, JWTAuthMiddleware)

	// Export Data ALL
	app.GET("/v1/exporvoteitem/:id", c.ExportVoteByItem, JWTAuthMiddleware)
}

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		token = strings.TrimSpace(token)
		log.Println("sss " + token)
		if token == "" || len(token) < 8 || strings.ToLower(token[:7]) != "bearer " {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		token = strings.TrimSpace(token[7:])
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		data, err := ClaimJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		log.Println(data)
		userId := data["userId"].(string)
		name := data["name"].(string)
		if userId == "" || name == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		c.Set("userId", userId)
		// if admin.IsAdmin(name) {
		// 	c.Set("admin", true)
		// }
		return next(c)
	}
}

func Adminmidleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Get("admin").(bool)
		if !name {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}

}

var jwtSecret = []byte("itemvote")

func GenerateJWT(userId string, name string) (string, error) {
	// Set expiration time to 3 months from now
	expirationTime := time.Now().Add(time.Hour * 24 * 30 * 3)

	// Create claims with user information
	claims := jwt.MapClaims{
		"userId": userId,
		"name":   name,
		"exp":    expirationTime.Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ClaimJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Use the same secret key used for signing tokens
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid claims")
}
