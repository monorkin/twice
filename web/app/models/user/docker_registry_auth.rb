# frozen_string_literal: true

module User::DockerRegistryAuth
  class Error < StandardError; end
  class UnpermittedScope < Error; end

  extend ActiveSupport::Concern

  DEFAULT_TOKEN_DURATION = 5.minutes

  def authenticate_registry_access(license_key)
    ActiveSupport::SecurityUtils.fixed_length_secure_compare(license_key, self.license_key)
  end

  def generate_registry_access_token(duration: DEFAULT_TOKEN_DURATION, service: nil)
    access_requests = if is_a?(Developer)
      DockerRegistry::AccessRequest.full_access
    else
      DockerRegistry::AccessRequest.pull_access_to_repositories(products.pluck(:repository))
    end

    DockerRegistry::Token.new(
      id: id,
      email: email_address,
      service: service,
      duration: duration,
      access_requests: access_requests
    )
  end
end

