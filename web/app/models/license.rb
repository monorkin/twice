# frozen_string_literal: true

class License < ApplicationRecord
  include KeyGenerator

  belongs_to :product
  belongs_to :customer

  validates :key, presence: true, uniqueness: true

  def self.find_by_key!(key)
    key = "" if key.blank?
    find_by!(key: key.to_s.strip.downcase.tr("-", ""))
  end

  def formatted_key
    return if key.blank?
    key.scan(/.{1,4}/).join("-")
  end
end
