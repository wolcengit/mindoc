package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/lifei6671/mindoc/models"
	"gopkg.in/russross/blackfriday.v2"
)

type DocumentList struct {
	DocumentId   int         `json:"id"`
	DocumentName string      `json:"title"`
	ParentId     interface{} `json:"parent"`
	Identify     string      `json:"identify"`
}

// MinDocRestController struct.
type MinDocRestController struct {
	BaseController
}

// PostContent API 设置文档内容.
func (c *MinDocRestController) PostContent() {
	c.Prepare()

	var req models.MinDocRest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	folderkey := req.Folder
	doctitle := req.Title
	dockey := req.Identify
	tokenkey := req.Token
	textmd := req.TextMD
	texthtml := req.TextHTML

	if doctitle == "" {
		c.JsonResult(6004, "文档名称不能为空")
	}
	if dockey == "" {
		c.JsonResult(6004, "文档标识不能为空")
	}
	if ok, err := regexp.MatchString(`^[a-z]+[a-zA-Z0-9_\-]*$`, dockey); !ok || err != nil {

		c.JsonResult(6003, "文档标识只能包含小写字母、数字，以及“-”和“_”符号,并且只能小写字母开头")
	}

	book, err := models.NewBook().FindByFieldFirst("private_token", tokenkey)
	if err != nil {
		beego.Error("token => ", err)
		c.JsonResult(6002, "系统权限不足["+tokenkey+"]")
	}
	beego.Info("req tokenkey =>" + tokenkey + "  -->" + fmt.Sprintf("%d", book.BookId))

	folder, err := models.NewDocument().FindByIdentityFirst(folderkey, book.BookId)
	if err != nil {
		beego.Error("folder => ", err)
		c.JsonResult(6002, "项目或类目不存在或权限不足")
	}
	beego.Info("req folderkey =>" + folderkey + "  -->" + fmt.Sprintf("%d", folder.BookId))
	if folder.BookId != book.BookId {
		c.JsonResult(6002, "folder和token不匹配")
	}

	doc, err := models.NewDocument().FindByIdentityFirst(dockey, folder.BookId)
	if err != nil {
		if err != orm.ErrNoRows {
			c.JsonResult(6006, "文档获取失败")
		}
	}
	if doc.BookId > 0 && folder.BookId != doc.BookId {
		c.JsonResult(6002, "文档标识已经被使用")
	}
	doc.BookId = book.BookId
	doc.MemberId = book.MemberId
	doc.Identify = dockey
	doc.Version = time.Now().Unix()
	doc.DocumentName = doctitle
	doc.ParentId = folder.DocumentId
	doc.Markdown = textmd
	if texthtml == "nil" {
		doc.Content = string(blackfriday.Run([]byte(doc.Markdown)))
	} else {
		doc.Content = texthtml
	}
	doc.Release = ""
	if err := doc.InsertOrUpdate(); err != nil {
		beego.Error("InsertOrUpdate => ", err)
		c.JsonResult(6005, "保存失败")
	}
	book.Version = time.Now().Unix()
	book.Update()

	//减少返回信息
	doc.Markdown = ""
	doc.Content = ""
	doc.Release = ""
	c.JsonResult(0, "ok", doc)
}

// BookCatalog
func (c *MinDocRestController) BookCatalog() {
	c.Prepare()

	dockey := c.Ctx.Input.Param(":key")
	tokenkey := c.GetString("token")

	if dockey == "" {
		c.JsonResult(6001, "项目没有指定")
	}
	book, err := models.NewBook().FindByFieldFirst("private_token", tokenkey)
	if err != nil {
		beego.Error("token => ", err)
		c.JsonResult(6002, "系统权限不足["+tokenkey+"]")
	}
	var docs []*models.Document

	count, err := orm.NewOrm().QueryTable(new(models.Document)).Filter("book_id", book.BookId).OrderBy("order_sort", "document_id").Limit(math.MaxInt32).All(&docs, "document_id", "document_name", "parent_id", "identify")

	if err != nil {
		beego.Error("read => ", err)
		c.JsonResult(6003, "读取失败")
	}
	catalogs := make([]*DocumentList, count)
	for index, item := range docs {
		cat := &DocumentList{}
		cat.DocumentId = item.DocumentId
		cat.DocumentName = item.DocumentName
		cat.ParentId = item.ParentId
		cat.Identify = item.Identify
		catalogs[index] = cat
	}

	//c.JsonResult(0, "ok", catalogs)
	returnJSON, err := json.Marshal(catalogs)

	if err != nil {
		beego.Error(err)
	}

	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))

	c.StopRun()
}

// BookSectionMarkdown
func (c *MinDocRestController) BookSectionMarkdown() {
	c.Prepare()

	dockey := c.Ctx.Input.Param(":key")
	docid := c.Ctx.Input.Param(":id")
	tokenkey := c.GetString("token")

	if dockey == "" || docid == "" {
		c.JsonResult(6001, "项目没有指定")
	}
	book, err := models.NewBook().FindByFieldFirst("private_token", tokenkey)
	if err != nil {
		beego.Error("token => ", err)
		c.JsonResult(6002, "系统权限不足["+tokenkey+"]")
	}

	doc, err := models.NewDocument().FindByIdentityFirst(docid, book.BookId)
	if err != nil {
		beego.Error("read => ", err)
		c.JsonResult(6003, "读取失败")
	}

	//c.JsonResult(0, doc.Markdown)
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/text; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(doc.Markdown))

	c.StopRun()
}

// BookSectionHtml
func (c *MinDocRestController) BookSectionHtml() {
	c.Prepare()

	dockey := c.Ctx.Input.Param(":key")
	docid := c.Ctx.Input.Param(":id")
	tokenkey := c.GetString("token")

	if dockey == "" || docid == "" {
		c.JsonResult(6001, "项目没有指定")
	}
	book, err := models.NewBook().FindByFieldFirst("private_token", tokenkey)
	if err != nil {
		beego.Error("token => ", err)
		c.JsonResult(6002, "系统权限不足["+tokenkey+"]")
	}

	doc, err := models.NewDocument().FindByIdentityFirst(docid, book.BookId)
	if err != nil {
		beego.Error("read => ", err)
		c.JsonResult(6003, "读取失败")
	}

	//c.JsonResult(0, doc.Content)
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/text; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(doc.Content))

	c.StopRun()
}
