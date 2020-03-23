package models

import "github.com/volatiletech/sqlboiler/queries/qm"

// BloggerQuery defines a func to query blogger table
type BloggerQuery func(mods ...qm.QueryMod) bloggerQuery

// UserdefinecategoryQuery defines a func to query blogger table
type UserdefinecategoryQuery func(mods ...qm.QueryMod) userdefinecategoryQuery
