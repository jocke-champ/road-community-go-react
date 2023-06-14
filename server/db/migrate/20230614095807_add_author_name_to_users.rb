class AddAuthorNameToUsers < ActiveRecord::Migration[7.0]
  def change
    add_column :users, :author_name, :string
  end
end
