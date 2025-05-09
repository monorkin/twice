# frozen_string_literal: true

class Product < ApplicationRecord
  has_many :licenses, dependent: :destroy

  validates :name, presence: true
  validates :slug, presence: true, uniqueness: true
end
