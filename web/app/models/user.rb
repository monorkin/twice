# frozen_string_literal: true

class User < ApplicationRecord
  include LicenseKeyGenerator
  include DockerRegistryAuth

  has_secure_password
  has_many :sessions, dependent: :destroy
  has_and_belongs_to_many :products

  normalizes :email_address, with: ->(e) { e.strip.downcase }

  validates :email, presence: true, uniqueness: true
  validates :license_key, presence: true, uniqueness: true
  validates :password_digest, presence: true
end
