# frozen_string_literal: true

class User < ApplicationRecord
  include DockerRegistryAuth

  has_secure_password
  has_many :sessions, dependent: :destroy
  has_many :licenses, dependent: :destroy, foreign_key: :owner_id
  has_many :products, through: :licenses

  normalizes :email_address, with: ->(e) { e.strip.downcase }

  validates :email_address, presence: true, uniqueness: true
  validates :password_digest, presence: true
end
