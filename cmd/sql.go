/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sqldb, err := sql.Open("sqlite3", "./database.db")
		if err != nil {
			panic(err)
		}
		db := bun.NewDB(sqldb, sqlitedialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
		sqldb.SetMaxIdleConns(1000)
		sqldb.SetConnMaxLifetime(100)

		u := new(Users)
		ctx := context.Background()
		// err = db.NewSelect().Model(u).Where("id = ?", 13).Scan(ctx, u)
		// fmt.Println(u)

		res, err := db.NewUpdate().
			Model(u).
			Column("email").
			Set("email = ?", "news").
			Where("id = ?", 13).
			Exec(ctx)

		fmt.Println(res.LastInsertId())

	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)

}

type Users struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID           string `param:"id" query:"id" form:"id" json:"id" xml:"id" bun:",pk,autoincrement"`
	Email        string `valid:"type(string),required" param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required,email" mod:"trim" bun:"email,notnull"`
	PasswordHash string `valid:"type(string),required" param:"passwordhash" query:"passwordhash" form:"passwordhash" json:"passwordhash" xml:"passwordhash"  bun:"passwordhash,notnull"`
	PasswordRaw  string `valid:"type(string)" param:"passwordraw" query:"passwordraw" form:"passwordraw" json:"passwordraw" xml:"passwordraw" validate:"required" scrub:"password" mod:"trim"  bun:"passwordraw,notnull"`
	Isdisabled   string `valid:"type(string)" param:"isdisabled" query:"isdisabled" form:"isdisabled" json:"isdisabled" xml:"isdisabled"  bun:"isdisabled,notnull"`
	SessionKey   string `valid:"type(string)" param:"sessionkey" query:"sessionkey" form:"sessionkey" json:"sessionkey" xml:"sessionkey"  bun:"sessionkey,notnull"`
	SessionName  string `valid:"type(string)" param:"sessionname" query:"sessionname" form:"sessionname" json:"sessionname" xml:"sessionname"  bun:"sessionname,notnull"`
	SessionToken string `valid:"type(string)" param:"sessiontoken" query:"sessiontoken" form:"sessiontoken" json:"sessiontoken" xml:"sessiontoken"  bun:"sessiontoken,notnull"`
	SiteToken    string `valid:"type(string),required" param:"sitetoken" query:"sitetoken" form:"sitetoken" json:"sitetoken" xml:"sitetoken"  bun:"sitetoken,notnull"`
}
