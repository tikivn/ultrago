package u_password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateDefault(password string) (passwordHash string, err error) {
	conf := &PasswordConfig{
		prefix:  "password",
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  64,
	}
	return Generate(conf, password)
}

func Generate(conf *PasswordConfig, password string) (passwordHash string, err error) {
	// Generate a Salt
	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return "", err
	}

	keyHash := argon2.IDKey([]byte(password), salt, conf.time, conf.memory, conf.threads, conf.keyLen)
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(keyHash)

	passwordHash = fmt.Sprintf("$%s$v=%d$m=%d,t=%d,p=%d$%s$%s",
		conf.prefix,
		argon2.Version,
		conf.memory,
		conf.time,
		conf.threads,
		b64Salt,
		b64Hash,
	)
	return passwordHash, nil
}

func Compare(password string, passwordHash string) (bool, error) {
	if passwordHash == "" {
		return false, fmt.Errorf("invalid password")
	}

	c := &PasswordConfig{}
	parts := strings.Split(passwordHash, "$")
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)
	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}

type PasswordConfig struct {
	prefix  string // service name
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}
