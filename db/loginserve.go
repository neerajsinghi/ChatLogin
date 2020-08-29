package db

import (
	model "LoginServer/models"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Insert allows populating database
func Register(user model.User) (res model.ResponseResult) {

	var result model.User
	err := model.FindOne(bson.M{"username": user.Username}, bson.M{}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				return
			}
			user.Password = string(hash)
			user.HostID = []string{""}
			_, err = model.InsertOne(user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				return
			}
			res.Result = "Registration Successful"
			update, err := model.UpdateOne(bson.M{"username": user.Username}, bson.M{"$pull": bson.M{"hostid": ""}})
			if err != nil {
				log.Println(err)
			} else {
				log.Println(update.MatchedCount)
			}
			return
		}

		res.Error = err.Error()
		return
	}

	res.Result = "Username already Exists!!"
	return
}
func Login(user model.User) (res model.ResponseResult, result model.User) {

	err := model.FindOne(bson.M{"username": user.Username}, bson.M{}).Decode(&result)

	if err != nil {
		res.Error = "Invalid username"
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		res.Error = "Invalid password"
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.Username,
		"firstname": result.FirstName,
		"lastname":  result.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		return
	}
	if user.HostID != nil {
		update, err := model.UpdateOne(bson.M{"username": user.Username}, bson.M{"$addToSet": bson.M{"hostid": user.HostID[0]}})
		if err != nil {

		} else {
			log.Println(update.MatchedCount)
		}
	}
	result.Token = tokenString
	result.Password = ""
	result.HostID, err = GetHosts(tokenString)
	return
}
func Profile(tokenString string) (result model.User, err error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)

		return
	} else {
		return
	}
}
func GetHosts(tokenString string) (hostid []string, err error) {
	var result model.User
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)
		var result2 model.User
		err = model.FindOne(bson.M{"username": result.Username}, bson.M{"hostid": 1}).Decode(&result2)
		if err != nil {
			log.Println(err)
			return
		}
		hostid = append(hostid, result2.HostID...)
		return
	} else {
		return
	}
}
