# frozen_string_literal: true

module License::Search
  extend ActiveSupport::Concern

  class_methods do
    def search(query)
      return all if query.blank?

      query = query.strip.downcase
      left_joins(:product).where("key LIKE :query OR products.name LIKE :query", query: "%#{query}%")
    end
  end
end
