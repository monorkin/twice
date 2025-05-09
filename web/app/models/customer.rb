# frozen_string_literal: true

class Customer < ApplicationRecord
  has_many :licenses, dependent: :destroy
  has_many :products, through: :licenses

  validates :email, presence: true, uniqueness: true

  def purchase(product)
    licenses.create!(product: product)
  end
end
