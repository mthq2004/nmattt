from Crypto.Random import get_random_bytes

# Generate a random 256-bit (32-byte) AES key
aes_key = get_random_bytes(16)

print("AES Key:", aes_key.hex())
