# frozen_string_literal: true

class RegistryController < ApplicationController
  TOKEN_DURATION = 5.minutes

  allow_unauthenticated_access

  # /registry/auth
  def auth
    authenticate_or_request_with_http_basic do |email, license_key|
      user = User.find_by(email_address: email)

      if user&.authenticate_registry_access(license_key)
        token = user.generate_registry_access_token(duration: TOKEN_DURATION, service: params[:service])
        payload = {
          token: token.to_s,
          expires_in: TOKEN_DURATION.to_i,
          issued_at: Time.now.utc.iso8601
        }

        Rails.logger.info token.payload.inspect
        Rails.logger.info payload.inspect

        render(json: payload, status: :ok)
      else
        deny_access
      end
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

