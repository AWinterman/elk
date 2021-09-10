// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// OwnersColumns holds the columns for the "owners" table.
	OwnersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
	}
	// OwnersTable holds the schema information for the "owners" table.
	OwnersTable = &schema.Table{
		Name:       "owners",
		Columns:    OwnersColumns,
		PrimaryKey: []*schema.Column{OwnersColumns[0]},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "owner_pets", Type: field.TypeInt, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_owners_pets",
				Columns:    []*schema.Column{PetsColumns[3]},
				RefColumns: []*schema.Column{OwnersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoryPetsColumns holds the columns for the "category_pets" table.
	CategoryPetsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeInt},
		{Name: "pet_id", Type: field.TypeString},
	}
	// CategoryPetsTable holds the schema information for the "category_pets" table.
	CategoryPetsTable = &schema.Table{
		Name:       "category_pets",
		Columns:    CategoryPetsColumns,
		PrimaryKey: []*schema.Column{CategoryPetsColumns[0], CategoryPetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_pets_category_id",
				Columns:    []*schema.Column{CategoryPetsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_pets_pet_id",
				Columns:    []*schema.Column{CategoryPetsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PetFriendsColumns holds the columns for the "pet_friends" table.
	PetFriendsColumns = []*schema.Column{
		{Name: "pet_id", Type: field.TypeString},
		{Name: "friend_id", Type: field.TypeString},
	}
	// PetFriendsTable holds the schema information for the "pet_friends" table.
	PetFriendsTable = &schema.Table{
		Name:       "pet_friends",
		Columns:    PetFriendsColumns,
		PrimaryKey: []*schema.Column{PetFriendsColumns[0], PetFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pet_friends_pet_id",
				Columns:    []*schema.Column{PetFriendsColumns[0]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "pet_friends_friend_id",
				Columns:    []*schema.Column{PetFriendsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		OwnersTable,
		PetsTable,
		CategoryPetsTable,
		PetFriendsTable,
	}
)

func init() {
	PetsTable.ForeignKeys[0].RefTable = OwnersTable
	CategoryPetsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategoryPetsTable.ForeignKeys[1].RefTable = PetsTable
	PetFriendsTable.ForeignKeys[0].RefTable = PetsTable
	PetFriendsTable.ForeignKeys[1].RefTable = PetsTable
}
