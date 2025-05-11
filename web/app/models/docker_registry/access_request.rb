# frozen_string_literal: true

class DockerRegistry::AccessRequest < Data.define(:type, :name, :actions)
  REPOSITORY_TYPE = "repository"
  REGISTRY_TYPE = "registry"
  WILDCARD_ACTIONS = ["*"].freeze
  PULL_ACTION = "pull"
  PUSH_ACTION = "push"

  def self.full_access
    [
      new("repository", "*", ["*"]),
      new("registry", "*", ["*"])
    ]
  end

  def self.pull_access_to_repositories(repositories)
    Array(repositories).map { |repository| new("repository", repository, ["pull"]) }
  end

  def self.parse(scope)
    type, name, action = scope.downcase.split(":")
    actions = action&.split(",")
    new(type, name, actions)
  end

  def to_s
    [type, name, actions&.join(",")].join(":")
  end
end

