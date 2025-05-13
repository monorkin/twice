json.registry_url "localhost:5000"
json.products @page.records, :id, :name, :repository
json.next_page_url @page.last? ? nil : url_for(page: @page.next_param)
