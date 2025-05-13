# frozen_string_literal: true

class RegistryController < ApplicationController
  allow_unauthenticated_access

  # GET|POST /registry/auth
  def auth
    authenticate_or_request_with_http_basic do |email, license_key|
      user = User.find_by(email_address: email)
      deny_access and return if user.blank?

      license = user.licenses.find_by_key(license_key)

      token = if license.present?
        user.registry_access_token_to_product(license.product, service: params[:service])
      elsif user.is_a?(Developer) && user.authenticate(license_key)
        user.registry_token_for_scope(scope: params[:scope], service: params[:service])
      end

      Rails.logger.info "Registry token: #{token.payload.inspect}"

      deny_access and return if token.blank?

      payload = {
        token: token.to_s,
        expires_in: token.duration.to_i,
        issued_at: Time.now.utc.iso8601
      }

      render(json: payload, status: :ok)
    end
  end

  private

    def deny_access
      render(
        json: {
          errors: [ { code: "UNAUTHORIZED", message: "Invalid credentials" } ]
        },
        status: :unauthorized
      )
    end
end

