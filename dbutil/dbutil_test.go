package dbutil_test

import (
  "testing"
  db "github.com/ChrHan/golang-sqlite-wrapper/dbutil"
  _ "github.com/mattn/go-sqlite3"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type DbutilSuite struct {
  suite.Suite
  dbutil *db.Dbutil
}

func TestDbutilSuite(t *testing.T) {
  suite.Run(t, &DbutilSuite{})
}

func (dc *DbutilSuite) SetupSuite() {
  dc.dbutil = db.New("test.db")
  dc.dbutil.Prepare()
}

func (dc *DbutilSuite) TestSelect() {
  result := dc.dbutil.Select()
  resultCount := dc.dbutil.SelectCount()
  assert.NotNil(dc.T(), result, "Result should not be Nil")
  assert.Equal(dc.T(), 0, resultCount, "Result should be 0")
}



