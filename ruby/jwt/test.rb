require 'openssl'
require 'base64'
require 'jwt'

rsa_private = OpenSSL::PKey::RSA.generate(2048)
rsa_public = rsa_private.public_key

puts rsa_private.class
puts rsa_private
puts rsa_public.class
puts rsa_public

payload = {data: 'test'}
token = JWT.encode payload, rsa_private, 'RS256'
puts token

decoded_token = JWT.decode token, rsa_public, true, { algorithm: 'RS256' }
puts decoded_token

rsa_private.tap { |p| puts '', 'PRIVATE RSA KEY (URL-safe Base64 encoded, PEM):', '', Base64.urlsafe_encode64(p.to_pem), '', 'PUBLIC RSA KEY (URL-safe Base64 encoded, PEM):', '', Base64.urlsafe_encode64(p.public_key.to_pem) }
