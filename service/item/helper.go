package vote

import (
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUser(id string) (string, bool) {

	if val, ok := userMap[id]; ok {
		return val.Name, true
	}
	return "", false
}

func VoteUserMap(itemid, userid string) {

	if _, ok := userVote[itemid]; !ok {
		userVote[itemid] = append(userVote[itemid], userid)
	}
}

func Converthex(id string) (primitive.ObjectID, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	return ID, err
}

func genTokenuser() {
	for i := 0; i < 10; i++ {
		user := User{}
		userID := fmt.Sprintf("user%d", i+1)
		name := fmt.Sprintf("%s", randomName(i))
		user.Name = name
		//user.Token, _ = GenerateJWT(userID, name)
		userMap[userID] = user
	}
	fmt.Println(userMap)
}

//StringToInt : convert string to i
func StringToInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

// StringToUInt :
func StringToUInt(s string) (i uint) {
	// i, _ = strconv.Atoi(s)
	i64, _ := strconv.ParseUint(s, 10, 32)
	i = uint(i64)
	return
}

//StringToFloat64 : convert string to i
func StringToFloat64(s string) (f float64) {
	f, _ = strconv.ParseFloat(s, 7)
	return
}

//StringToUint8 : convert string to uint8
func StringToUint8(s string) uint8 {
	i, _ := strconv.ParseUint(s, 10, 8)
	return uint8(i)
}

//StringToUint16 : convert string to uint16
func StringToUint16(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

//IsEmpty :
func IsEmpty(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	// fmt.Printf("s: |%s| (%d)", s, len(s))
	// fmt.Printf("s: |%s| (%d)\n", s, strings.Count(s, ""))
	if strings.Count(s, "") > 1 { //>0
		return false
	}
	return true
}

func randomName(index int) string {
	// Generate a random name with a simple pattern
	names := []string{"Alice", "Bob", "Charlie", "David", "Emily", "Frank", "Grace", "Henry", "Isabella", "Jack"}
	return names[index]
}

func UserUniq(id string) {
	userVote[id] = UniqArray(userVote[id])
}

func UniqArray[T comparable](data []T) []T {
	seen := map[T]bool{}
	result := []T{}

	for _, v := range data {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func Contains[T comparable](data []T, target T) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}
	return false
}
