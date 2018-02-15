package dbutil_test

import (
  "testing"
  "os"
  db "github.com/ChrHan/golang-sqlite-wrapper/dbutil"
  _ "github.com/mattn/go-sqlite3"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
  "github.com/icrowley/fake"
)

const DB_FILENAME = "test.db"

type DbutilSuite struct {
  suite.Suite
  dbutil *db.Dbutil
}

func TestDbutilSuite(t *testing.T) {
  suite.Run(t, &DbutilSuite{})
}

func (dc *DbutilSuite) SetupSuite() {
  os.Remove(DB_FILENAME)
  dc.dbutil = db.New(DB_FILENAME)
  dc.dbutil.Prepare()
}

func (dc *DbutilSuite) Test1Select() {
  result := dc.dbutil.Select()
  resultCount := dc.dbutil.SelectCount()
  assert.NotNil(dc.T(), result, "Result should not be Nil")
  assert.Equal(dc.T(), 0, resultCount, "Result should be 0")
}

func (dc *DbutilSuite) Test2Insert() {
  id := fake.Digits()
  product_name := fake.Product()
  dc.dbutil.Insert(id, product_name)
  result := dc.dbutil.Select()
  resultCount := dc.dbutil.SelectCount()
  assert.NotNil(dc.T(), result, "Result should not be Nil")
  assert.NotEqual(dc.T(), 0, resultCount, "Result should NOT be 0")
}
