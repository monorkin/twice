# frozen_string_literal: true

module Product::Search
  extend ActiveSupport::Concern

  class_methods do
    def search(query)
      return all if query.blank?

      query = query.strip.downcase
      where("name LIKE :query OR repository LIKE :query", query: "%#{query}%")
    end
  end
end
