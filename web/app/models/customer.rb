# frozen_string_literal: true

class Customer < ApplicationRecord
  include LicenseKeyGenerator

  validates :email, presence: true, uniqueness: true
  validates :license_key, presence: true, uniqueness: true

  def purchase(product)
    licenses.create!(product: product)
  end
end
