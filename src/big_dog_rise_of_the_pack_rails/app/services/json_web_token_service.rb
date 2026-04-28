class JsonWebTokenService
  SECRET = ENV.fetch("JWT_SECRET")

  def self.encode(payload)
    JWT.encode(payload, SECRET, "HS256")
  end

  def self.decode(token)
    JWT.decode(token, SECRET, true, { algorithm: "HS256" })[0]
  end
end
