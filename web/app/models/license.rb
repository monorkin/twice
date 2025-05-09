# frozen_string_literal: true

class License < ApplicationRecord
  include KeyGenerator

  belongs_to :product
  belongs_to :customer

  validates :key, presence: true, uniqueness: true
end
