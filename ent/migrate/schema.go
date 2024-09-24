// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BangumisColumns holds the columns for the "bangumis" table.
	BangumisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "english_title", Type: field.TypeString, Nullable: true},
		{Name: "japanese_title", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "score", Type: field.TypeFloat64, Nullable: true},
		{Name: "tags", Type: field.TypeString, Nullable: true},
		{Name: "synopsis", Type: field.TypeString, Nullable: true},
		{Name: "cover_image", Type: field.TypeString, Nullable: true},
		{Name: "trailer_url", Type: field.TypeString, Nullable: true},
	}
	// BangumisTable holds the schema information for the "bangumis" table.
	BangumisTable = &schema.Table{
		Name:       "bangumis",
		Columns:    BangumisColumns,
		PrimaryKey: []*schema.Column{BangumisColumns[0]},
	}
	// CastMembersColumns holds the columns for the "cast_members" table.
	CastMembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "character_name", Type: field.TypeString, Nullable: true},
		{Name: "profile_path", Type: field.TypeString, Nullable: true},
		{Name: "bangumi_cast_members", Type: field.TypeInt, Nullable: true},
		{Name: "episode_cast_members", Type: field.TypeInt, Nullable: true},
		{Name: "season_id", Type: field.TypeInt, Nullable: true},
	}
	// CastMembersTable holds the schema information for the "cast_members" table.
	CastMembersTable = &schema.Table{
		Name:       "cast_members",
		Columns:    CastMembersColumns,
		PrimaryKey: []*schema.Column{CastMembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cast_members_bangumis_cast_members",
				Columns:    []*schema.Column{CastMembersColumns[4]},
				RefColumns: []*schema.Column{BangumisColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "cast_members_episodes_cast_members",
				Columns:    []*schema.Column{CastMembersColumns[5]},
				RefColumns: []*schema.Column{EpisodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "cast_members_seasons_cast_members",
				Columns:    []*schema.Column{CastMembersColumns[6]},
				RefColumns: []*schema.Column{SeasonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EpisodesColumns holds the columns for the "episodes" table.
	EpisodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "episode_number", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString, Nullable: true},
		{Name: "air_date", Type: field.TypeTime, Nullable: true},
		{Name: "season_id", Type: field.TypeInt, Nullable: true},
	}
	// EpisodesTable holds the schema information for the "episodes" table.
	EpisodesTable = &schema.Table{
		Name:       "episodes",
		Columns:    EpisodesColumns,
		PrimaryKey: []*schema.Column{EpisodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "episodes_seasons_episodes",
				Columns:    []*schema.Column{EpisodesColumns[5]},
				RefColumns: []*schema.Column{SeasonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SeasonsColumns holds the columns for the "seasons" table.
	SeasonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "season_number", Type: field.TypeInt},
		{Name: "air_date", Type: field.TypeTime, Nullable: true},
		{Name: "episode_count", Type: field.TypeInt},
		{Name: "bangumi_id", Type: field.TypeInt, Nullable: true},
	}
	// SeasonsTable holds the schema information for the "seasons" table.
	SeasonsTable = &schema.Table{
		Name:       "seasons",
		Columns:    SeasonsColumns,
		PrimaryKey: []*schema.Column{SeasonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "seasons_bangumis_seasons",
				Columns:    []*schema.Column{SeasonsColumns[4]},
				RefColumns: []*schema.Column{BangumisColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BangumisTable,
		CastMembersTable,
		EpisodesTable,
		SeasonsTable,
		UsersTable,
	}
)

func init() {
	CastMembersTable.ForeignKeys[0].RefTable = BangumisTable
	CastMembersTable.ForeignKeys[1].RefTable = EpisodesTable
	CastMembersTable.ForeignKeys[2].RefTable = SeasonsTable
	EpisodesTable.ForeignKeys[0].RefTable = SeasonsTable
	SeasonsTable.ForeignKeys[0].RefTable = BangumisTable
}