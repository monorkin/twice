# frozen_string_literal: true

class DockerRegistry::Token
  ISSUER = ENV.fetch("TOKEN_ISSUER", "registry-auth")
  PRIVATE_KEY_PATH = Rails.root.join("storage", "docker_registry.key").freeze
  PUBLIC_KEY_PATH = Rails.root.join("storage", "docker_registry.crt").freeze
  PRIVATE_KEY = begin
    OpenSSL::PKey::RSA.new(PRIVATE_KEY_PATH.read)
  rescue => e
    Rails.logger.error("Failed to load private key: #{e.message}")
    raise "Failed to load private key: #{e.message}"
  end
  PUBLIC_KEY = begin
    OpenSSL::X509::Certificate.new(PUBLIC_KEY_PATH.read)
  rescue => e
    Rails.logger.error("Failed to load public key: #{e.message}")
    raise "Failed to load public key: #{e.message}"
  end
  KID = OpenSSL::Digest.new("SHA256").update(PRIVATE_KEY.public_key.to_der).hexdigest.upcase.scan(/.{1,2}/).join(":")
  X5C = Base64.strict_encode64(PUBLIC_KEY.to_der)

  attr_reader :duration, :email, :access_requests, :service, :id

  def initialize(duration:, email:, access_requests:, service:, id:)
    @duration = duration
    @email = email
    @access_requests = access_requests
    @service = service
    @id = id
  end

  def to_s
    JWT.encode(payload, PRIVATE_KEY, "RS256", headers)
  end

  def headers
    { kid: KID, x5c: [ X5C ] }
  end

  def payload
    {
      iss: ISSUER,
      sub: email,
      aud: service,
      exp: duration.from_now.to_i,
      nbf: issued_at.to_i,
      iat: issued_at.to_i,
      jti: jti,
      access: access_requests.map(&:to_h)
    }
  end

  def issued_at
    Time.now.utc
  end

  def jti
    digest = Digest::SHA256.new
    digest << id.to_s
    digest << Time.now.to_f.to_s

    access_requests.sort_by(&:to_s).each do |access_request|
      digest << access_request.to_s
    end

    digest.hexdigest
  end
end
