package snowflake

import (
	"fmt"
	"testing"
)

func TestGenerateSnowflakeID(t *testing.T) {
	fmt.Println(GenerateSnowflakeID())
}

func TestGenerateSha1SnowflakeID(t *testing.T) {
	fmt.Println(GenerateSha1SnowflakeID())
}

func TestExtractSnowflakeTime(t *testing.T) {
	snowflakeID := GenerateSnowflakeID()
	fmt.Println(ExtractTimeFromSnowflakeID(snowflakeID))
}

func TestGenerateSha1SnowflakeIDWithTime(t *testing.T) {
	fmt.Println(GenerateSha1SnowflakeIDWithTime())
}

func TestExtractGenerateSha1SnowflakeID(t *testing.T) {
	snowflake := GenerateSha1SnowflakeIDWithTime()
	fmt.Println(snowflake)
	fmt.Println(ExtractSnowflakeTime(GenerateSha1SnowflakeIDWithTime()))
}
