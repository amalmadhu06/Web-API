package initializers

import "jwt_auth/models"

func SyncDatabase() {

	DB.AutoMigrate(&models.User{}) //this will create a new table in our databse (migrate....)
	//abouve line of code created a table name 'user' in databse 'test3'

	// schema of the table users
	// 	Table "public.users"
	// 	Column   |           Type           | Collation | Nullable |              Default
	//  ------------+--------------------------+-----------+----------+-----------------------------------
	//   id         | bigint                   |           | not null | nextval('users_id_seq'::regclass)
	//   created_at | timestamp with time zone |           |          |
	//   updated_at | timestamp with time zone |           |          |
	//   deleted_at | timestamp with time zone |           |          |
	//   email      | text                     |           |          |
	//   password   | text                     |           |          |
	// -------------+--------------------------+-----------+----------+------------------------------------
	//  Indexes:
	// 	 "users_pkey" PRIMARY KEY, btree (id)
	// 	 "idx_users_deleted_at" btree (deleted_at)
	// 	 "users_email_key" UNIQUE CONSTRAINT, btree (email)

}
