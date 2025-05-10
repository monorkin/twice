# frozen_string_literal: true

class License < ApplicationRecord
  belongs_to :product
  belongs_to :customer
end
