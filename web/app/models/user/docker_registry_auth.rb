# frozen_string_literal: true

module User::DockerRegistryAuth
  class Error < StandardError; end
  class UnpermittedScope < Error; end

  extend ActiveSupport::Concern

  DEFAULT_TOKEN_DURATION = 5.minutes

  def full_access_registry_token(duration: DEFAULT_TOKEN_DURATION, service: nil)
    access_requests = DockerRegistry::AccessRequest.full_access_to_registry

    DockerRegistry::Token.new(
      id: id,
      email: email_address,
      service: service,
      duration: duration,
      access_requests: access_requests
    )
  end

  def registry_token_for_scope(scope:, duration: DEFAULT_TOKEN_DURATION, service: nil)
    access_requests = DockerRegistry::AccessRequest.parse(scope)

    DockerRegistry::Token.new(
      id: id,
      email: email_address,
      service: service,
      duration: duration,
      access_requests: access_requests
    )
  end

  def registry_access_token_to_product(product = nil, duration: DEFAULT_TOKEN_DURATION, service: nil)
    access_request = DockerRegistry::AccessRequest.pull_access_to_repository(product.repository)

    DockerRegistry::Token.new(
      id: id,
      email: email_address,
      service: service,
      duration: duration,
      access_requests: [access_request]
    )
  end
end

