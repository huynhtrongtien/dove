package services

import (
	"fmt"

	"github.com/speps/go-hashids/v2"
)

const hashIDMinLenght = 8

type HashID struct {
	HashIDs *hashids.HashID
}

func (h *HashID) EncodeInt64(numbers []int64) (string, error) {
	return h.HashIDs.EncodeInt64(numbers)
}

func (h *HashID) DecodeInt64(hash string) ([]int64, error) {
	return h.HashIDs.DecodeInt64WithError(hash)
}

var HashKeyGenerator *HashID

func initHashIDs(saltKey string) {
	hd := hashids.NewData()
	hd.Salt = saltKey
	hd.MinLength = hashIDMinLenght
	hashIDs, _ := hashids.NewWithData(hd)
	HashKeyGenerator = &HashID{
		HashIDs: hashIDs,
	}
}

func GenerateToken(categoryID, objectID int64) (string, error) {
	return HashKeyGenerator.EncodeInt64([]int64{categoryID, objectID})
}

func ExtractToken(token string) (int64, int64, error) {
	ints, err := HashKeyGenerator.DecodeInt64(token)
	if err != nil {
		return 0, 0, err
	}

	if len(ints) < 3 {
		return 0, 0, fmt.Errorf("token invalid")
	}

	return ints[0], ints[1], nil
}
