package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestCryptor(t *testing.T) {
	t.Parallel()

	t.Run("the encrypted message is correctly decrypted", func(t *testing.T) {
		t.Parallel()

		keys := Keys{
			Key{173, 26, 105, 188, 143, 153, 10, 18,
				122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19,
				217, 84, 254, 254, 144, 159, 61, 99},
		}

		cryptor := NewCryptor(keys)

		expectedMessage := "Hello, world"

		cipher, err := cryptor.Encrypt([]byte(expectedMessage))
		require.NoError(t, err)
		assert.NotEqualValues(t, expectedMessage, string(cipher))
		message, err := cryptor.Decrypt(cipher)
		require.NoError(t, err)
		assert.EqualValues(t, expectedMessage, message)
	})

	t.Run("decryption takes place with the key that was used for encrypted", func(t *testing.T) {
		t.Parallel()

		newestKey := Keys{
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
		}

		keys := Keys{
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
		}

		cryptor := NewCryptor(keys)

		cipher, err := cryptor.Encrypt([]byte("Hello, world"))
		require.NoError(t, err)
		assert.EqualValues(t, 2, cipher[len(cipher)-1:][0])

		cryptor = NewCryptor(append(keys, newestKey...))
		message, err := cryptor.Decrypt(cipher)
		assert.EqualValues(t, "Hello, world", message)

	})

	t.Run("encryption always happens with the freshest key", func(t *testing.T) {
		t.Parallel()

		newestKey := Keys{
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
		}

		keys := Keys{
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
			Key{
				173, 26, 105, 188, 143, 153, 10, 18, 122, 14, 191, 193, 190, 157, 28, 52,
				243, 37, 101, 184, 183, 230, 70, 19, 217, 84, 254, 254, 144, 159, 61, 99,
			},
		}

		cryptor := NewCryptor(keys)

		cipher, err := cryptor.Encrypt([]byte("Hello, world"))
		require.NoError(t, err)
		assert.EqualValues(t, 2, cipher[len(cipher)-1:][0])

		cryptor = NewCryptor(append(keys, newestKey...))

		cipher, err = cryptor.Encrypt([]byte("Hello, world"))
		require.NoError(t, err)
		assert.EqualValues(t, 5, cipher[len(cipher)-1:][0])
	})
}
