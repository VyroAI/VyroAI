package snowflake

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var counter int64 = 0
var workerID int64 = 0
var systemID int64 = 0

func GenerateSnowflakeID() int64 {
	if counter == 4095 {
		counter = 0
	}

	currentUnixTime := time.Now().UnixNano() / int64(time.Millisecond)
	binaryUnixTime := fmt.Sprintf("%048b", currentUnixTime)
	workerIDBinaryValue := strconv.FormatInt(workerID, 2)
	systemIDBinaryValue := strconv.FormatInt(systemID, 2)
	counterBinaryValue := fmt.Sprintf("%014b", counter)

	var binarySnowflakeID strings.Builder

	binarySnowflakeID.WriteString(binaryUnixTime)
	binarySnowflakeID.WriteString(workerIDBinaryValue)
	binarySnowflakeID.WriteString(systemIDBinaryValue)
	binarySnowflakeID.WriteString(counterBinaryValue)

	counter++

	snowflake, _ := strconv.ParseInt(binarySnowflakeID.String(), 2, 64)

	return snowflake
}

func ExtractTimeFromSnowflakeID(snowflakeID int64) int64 {
	snowflakeBinary := strconv.FormatInt(snowflakeID, 2)
	paddedBinary := fmt.Sprintf("%064s", snowflakeBinary)

	unixTimeBinary, _ := strconv.ParseInt(paddedBinary[:48], 2, 64)

	return unixTimeBinary

}

func GenerateSha1SnowflakeID() string {
	hash := sha1.Sum([]byte(strconv.Itoa(int(GenerateSnowflakeID()))))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func GenerateSha1SnowflakeIDWithTime() string {
	snowflakeID := GenerateSnowflakeID()
	unixTime := ExtractTimeFromSnowflakeID(snowflakeID)

	hash := sha1.Sum([]byte(strconv.Itoa(int(snowflakeID))))
	hashString := hex.EncodeToString(hash[:])
	insertString := strings.Split(strconv.Itoa(int(unixTime)), "")

	str := strings.Split(hashString, "")

	var newString []string

	for k, v := range str {

		if k%2 == 0 || k >= 26 {
			newString = append(newString, v)
		} else {
			newString = append(newString, insertString[k/2])
			newString = append(newString, v)
		}
	}

	return strings.Join(newString, "")
}

func ExtractSnowflakeTime(snowflakeID string) int64 {
	insertString := strings.Split(snowflakeID, "")
	var unixString strings.Builder

	for i := 1; i <= 37; i += 3 {
		unixString.WriteString(insertString[i])
	}
	unixInt, _ := strconv.Atoi(unixString.String())
	return int64(unixInt)

}
