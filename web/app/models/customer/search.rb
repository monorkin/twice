# frozen_string_literal: true

module Customer::Search
  extend ActiveSupport::Concern

  class_methods do
    def search(query)
      return all if query.blank?

      query = query.strip.downcase
      left_joins(:licenses).distinct.where("email_address LIKE :query OR licenses.key LIKE :query", query: "%#{query}%")
    end
  end
end
