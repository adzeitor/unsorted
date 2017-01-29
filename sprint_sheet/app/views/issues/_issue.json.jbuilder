json.extract! issue, :id, :title, :description, :sprint_id, :board_id, :assigned_to_id, :created_at, :updated_at
json.url issue_url(issue, format: :json)