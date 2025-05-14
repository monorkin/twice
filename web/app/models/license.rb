# frozen_string_literal: true

class License < ApplicationRecord
  include KeyGenerator
  include Search

  belongs_to :owner, class_name: "User"
  belongs_to :product

  validates :key, presence: true, uniqueness: true
end
