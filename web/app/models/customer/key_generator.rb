# frozen_string_literal: true

module Customer::LicenseKeyGenerator
  KEY_LENGTH = 16

  extend ActiveSupport::Concern

  included do
    before_validation :generate_license_key
  end

  def generate_license_key(length: KEY_LENGTH)
    loop do
      self.key = SecureRandom.base36(length).scan(/.{1,4}/).join("-")
      break unless License.exists?(key: key)
    end
  end
end
