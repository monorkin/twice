Rails.application.routes.draw do
  resource :session
  resources :passwords, param: :token
  get "up" => "rails/health#show", as: :rails_health_check

  match "/registry/auth", to: "registry#auth", as: :registry_auth, via: %i[ get post ]

  get "/install/:license_key", to: "install#install", as: :install
  get "/install/:license_key/download", to: "install#download", as: :install_download

  post "webhooks/order/created", to: "webhooks#order_created", as: :order_created_webhook
end
