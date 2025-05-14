# frozen_string_literal: true

class Product < ApplicationRecord
  include RepositoryValidations

  has_many :licenses, dependent: :destroy

  normalizes :repository, with: -> { it.strip.downcase }

  validates :name, presence: true
  validates :repository, uniqueness: true
end
