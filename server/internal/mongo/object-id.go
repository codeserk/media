package mongo

import (
	"encoding/binary"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToObjectId(objectIdHex string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(objectIdHex)
}

func ObjectIdHexToInt(objectIdHex string) (uint64, error) {
	objectId, err := ToObjectId(objectIdHex)
	if err != nil {
		return 0, fmt.Errorf("invalid hex: %v", objectId)
	}

	return ObjectIdToInt(objectId), nil
}

func ObjectIdToInt(objId primitive.ObjectID) uint64 {
	timestampBytes := objId[:8]

	return uint64(binary.LittleEndian.Uint64(timestampBytes))
}
