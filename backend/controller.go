package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	bcrypt "golang.org/x/crypto/bcrypt"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// AuthenticateUser vérifie les identifiants via le nom d'utilisateur et renvoie l'utilisateur authentifié.
func AuthenticateUser(db *sql.DB, username, password string) (User, error) {
	user, err := GetUserByUsername(db, username)
	if err != nil {
		return User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return User{}, fmt.Errorf("identifiants invalides")
	}
	return user, nil
}


// Génère un token JWT pour un utilisateur donné.
func generateToken(userID int) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET non défini")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	return token.SignedString([]byte(jwtSecret))
}

// Route d'inscription.
func signupHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur lors du hachage du mot de passe"})
		}
		user.Password = string(hashedPassword)
		user.Role = "user" // Rôle par défaut
		
		if err := CreateUser(db, &user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Inscription réussie"})
	}
}



// Route de connexion.
func loginHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var creds struct {
			UserName    string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&creds); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		user, err := AuthenticateUser(db, creds.UserName, creds.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Identifiants invalides"})
		}
		token, err := generateToken(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur lors de la création du token"})
		}
		session := UserSession{
			UserID: user.ID,
			Token:  token,
			Expiry: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		}
		if err := CreateUserSession(db, session); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur lors de la sauvegarde du token"})
		}
		return c.JSON(fiber.Map{"token": token})
	}
}

// Route protégée qui renvoie les informations de l'utilisateur courant.
func meHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))
		user, err := GetUserByID(db, userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Utilisateur non trouvé"})
		}
		// Ne pas renvoyer le mot de passe
		user.Password = ""
		return c.JSON(user)
	}
}

// Configuration des routes.
func BuildRoutes(app *fiber.App, db *sql.DB) error {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	// Routes publiques
	app.Post("/signup", signupHandler(db))
	app.Post("/login", loginHandler(db))

	// Groupe de routes protégées par JWT.
	jwtSecretEnv := os.Getenv("JWT_SECRET")
	if jwtSecretEnv == "" {
		return fmt.Errorf("JWT_SECRET non défini")
	}
	protected := app.Group("/", jwtMiddleware.New(jwtMiddleware.Config{
		SigningKey: []byte(jwtSecretEnv),
	}))
	protected.Get("/me", meHandler(db))

	app.Listen(":3000")
	return nil
}
