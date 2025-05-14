# frozen_string_literal: true

module License::KeyGenerator
  KEY_LENGTH = 16

  extend ActiveSupport::Concern

  included do
    before_validation :generate_key
  end

  def generate_key(length: KEY_LENGTH)
    loop do
      self.key = SecureRandom.base36(length).scan(/.{1,4}/).join("-")
      break unless self.class.exists?(key: key)
    end
  end
end
