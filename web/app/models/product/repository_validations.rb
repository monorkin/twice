# frozen_string_literal: true

module Product::RepositoryValidations
  extend ActiveSupport::Concern

  REPOSITORY_NAME_FORMAT = /\A[a-z0-9._-]+\Z/

  included do
    validate :validate_repository_format
  end

  def validate_repository_format
    repository = self.repository

    if repository.blank?
      errors.add(:repository, "can't be blank")
      return
    end

    if repository.include?(":")
      errors.add(:repository, "can't contain a tag (remove the colon and everything after it)")
      repository = repository.split(":").first
    end

    if repository.length > 255
      errors.add(:repository, "is too long (maximum is 255 characters)")
    end

    if repository.match?(/-{2,}|\.{2,}|_{2,}/)
      errors.add(:repository, "can't contain consecutive dashes, dots, or underscores")
    end

    parts = repository.split("/")
    namespace = nil
    name = nil

    if parts.count == 1
      namespace = nil
      name = parts.first
    elsif parts.count == 2
      namespace, name = parts
    else
      errors.add(:repository, "can't contain more than one slash")
      return
    end

    if name.blank?
      errors.add(:repository, "name can't be blank (the part after '/')")
    elsif !name.match?(REPOSITORY_NAME_FORMAT)
      errors.add(:repository, "name must start with a lowercase letter or number and can only contain letters, numbers, dashes, dots and underscores")
    end

    return if namespace.nil?

    if namespace.blank?
      errors.add(:repository, "namespace can't be blank (the part before '/')")
    elsif !namespace.match?(REPOSITORY_NAME_FORMAT)
      errors.add(:repository, "namespace must start with a lowercase letter or number and can only contain letters, numbers, dashes, dots and underscores")
    end
  end
end
