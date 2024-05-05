import random

def generate_keypair(p, q):
    n = p * q
    phi = (p - 1) * (q - 1)

    # Chọn một số nguyên e sao cho e và phi(n) nguyên tố cùng nhau
    e = random.randrange(1, phi)
    g = gcd(e, phi)
    while g != 1:
        e = random.randrange(1, phi)
        g = gcd(e, phi)

    # Tìm số nguyên d thỏa mãn (d * e) % phi = 1
    d = mod_inverse(e, phi)

    # Trả về khóa công khai và khóa riêng tư
    # Khóa công khai là cặp (e, n) và khóa riêng tư là cặp (d, n)
    return ((e, n), (d, n))

def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a

def mod_inverse(a, m):
    m0, x0, x1 = m, 0, 1
    while a > 1:
        q = a // m
        m, a = a % m, m
        x0, x1 = x1 - q * x0, x0
    return x1 + m0 if x1 < 0 else x1

def encrypt(public_key, plaintext):
    e, n = public_key
    # Mã hóa từng ký tự trong thông điệp sử dụng khóa công khai
    encrypted_msg = [pow(ord(char), e, n) for char in plaintext]
    return encrypted_msg

def decrypt(private_key, encrypted_msg):
    d, n = private_key
    # Giải mã từng ký tự trong văn bản mã hóa sử dụng khóa riêng tư
    decrypted_msg = [chr(pow(char, d, n)) for char in encrypted_msg]
    return ''.join(decrypted_msg)

# Số nguyên tố p và q
p = 61
q = 53

# Tạo khóa công khai và khóa riêng tư
public_key, private_key = generate_keypair(p, q)

# Thông điệp cần mã hóa
plaintext = "Hello, RSA!"

print("Thông điệp đã mã hóa:", public_key)
print("Thông điệp đã mã hóa:", private_key)


# Mã hóa thông điệp
encrypted_msg = encrypt(public_key, plaintext)
print("Thông điệp đã mã hóa:", encrypted_msg)

# Giải mã thông điệp
decrypted_msg = decrypt(private_key, encrypted_msg)
print("Thông điệp đã giải mã:", decrypted_msg)
