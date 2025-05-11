# frozen_string_literal: true

class Product < ApplicationRecord
  include RepositoryValidations

  has_and_belongs_to_many :users

  normalizes :repository, with: -> { it.strip.downcase }

  validates :name, presence: true
  validates :repository, uniqueness: true
end
