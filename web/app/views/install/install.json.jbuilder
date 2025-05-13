json.owner do
  json.(@license.owner, :id, :email_address)
end

json.product do
  json.(@license.product, :id, :name, :repository)
  json.registry "localhost:5000"
end
