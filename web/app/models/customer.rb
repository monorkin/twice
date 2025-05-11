# frozen_string_literal: true

class Customer < User
  before_validation :generate_password, on: :create

  def purchase(product)
    unless self.products.exists?(product)
      self.products << product
    end
  end

  def generate_password
    if password_digest.blank?
      self.password = SecureRandom.hex(32)
    end
  end
end
