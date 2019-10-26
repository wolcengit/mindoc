package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type LinkDocument struct {
	LinkId int `json:"link_id"`
}

type LinkDocumentTree struct {
	DocumentId   int    `json:"id"`
	ParentId     int    `json:"pId"`
	DocumentName string `json:"name"`
	Checked      bool   `json:"checked"`
	Open         bool   `json:"open"`
}

func NewLinkDocument() *LinkDocument {
	return &LinkDocument{}
}

func (m *LinkDocument) SetLinkBookDocuments(book_id int, link_docs string) (err error) {

	book, err := NewBook().Find(book_id)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	sql := `DELETE FROM md_documents WHERE book_id = ? AND FIND_IN_SET(link_id,?) = 0 `
	o.Raw(sql, book.BookId, link_docs).Exec()

	sql = `INSERT INTO md_documents (document_name,identify,book_id,parent_id,order_sort,create_time,member_id,modify_time,modify_at,version,link_id ) SELECT document_name,CONCAT(identify,document_id),?,parent_id,order_sort,create_time,member_id,modify_time,modify_at,version,document_id  FROM md_documents  WHERE book_id = ? AND FIND_IN_SET(document_id,?) > 0  AND document_id NOT IN(SELECT link_id FROM md_documents WHERE book_id = ? )`
	o.Raw(sql, book.BookId, book.LinkBook, link_docs, book.BookId).Exec()

	sql = `UPDATE md_documents a,md_documents b SET a.parent_id = b.parent_id WHERE a.link_id = b.document_id AND a.book_id = ? `
	o.Raw(sql, book.BookId).Exec()

	sql = `UPDATE md_documents as a, (select b.parent_id as bpid,b.document_id as bdid,c.document_id as cpid from md_documents b ,md_documents c where b.parent_id = c.link_id and c.book_id = ?) as b SET a.parent_id = b.cpid  WHERE a.book_id = ? AND a.parent_id > 0 AND a.parent_id = b.bpid;`
	o.Raw(sql, book.BookId, book.BookId).Exec()

	m.UpdateLinkBookDocuments(book.BookId)

	return nil
}

func (m *LinkDocument) UpdateLinkBookDocument(doc_id int) (err error) {
	o := orm.NewOrm()
	sql := `UPDATE md_documents a,md_documents b SET 
		a.document_name = b.document_name,
		a.identify = b.identify,
		a.markdown = b.markdown,
		a.release = b.release,
		a.content = b.content,
		a.order_sort = b.order_sort,
		a.member_id = b.member_id,
		a.modify_time = b.modify_time,
		a.modify_at = b.modify_at,
		a.version = b.version 
		WHERE a.document_id = ? AND a.link_id = b.document_id `
	o.Raw(sql, doc_id).Exec()

	return nil
}

func (m *LinkDocument) UpdateLinkBookDocuments(book_id int) (err error) {
	o := orm.NewOrm()
	sql := `UPDATE md_documents a,md_documents b SET 
		a.document_name = b.document_name,
		a.identify = b.identify,
		a.markdown = b.markdown,
		a.release = b.release,
		a.content = b.content,
		a.order_sort = b.order_sort,
		a.member_id = b.member_id,
		a.modify_time = b.modify_time,
		a.modify_at = b.modify_at,
		a.version = ? 
		WHERE a.book_id = ? AND a.link_id = b.document_id `
	o.Raw(sql, time.Now().Unix(),book_id).Exec()

	return nil
}

// GetLinkBookDocuments ...[{ id:1, pId:0, name:"数据错误"}];
func (m *LinkDocument) GetLinkBookDocuments(book_id int) (doclinks string, doclist string, err error) {
	o := orm.NewOrm()
	book, err := NewBook().Find(book_id)
	if err != nil {
		return "", "", err
	}
	var docs0 []*Document
	sql1 := `SELECT document_id,link_id FROM md_documents WHERE book_id = ? ORDER BY order_sort  ,document_id  `
	count0, _ := o.Raw(sql1, book_id).QueryRows(&docs0)
	doclinks = ""
	if count0 > 0 {
		for _, item := range docs0 {
			doclinks = doclinks + strconv.Itoa(item.LinkId) + ","
		}
	}

	var docs []*Document
	sql2 := `SELECT document_id,parent_id,document_name,FIND_IN_SET(document_id,?) AS modify_at FROM md_documents WHERE book_id = ? ORDER BY order_sort  ,document_id  `
	count, _ := o.Raw(sql2, doclinks, book.LinkBook).QueryRows(&docs)
	doclist = ""
	if count > 0 {
		trees := make([]*LinkDocumentTree, count)
		for index, item := range docs {
			tree := &LinkDocumentTree{}
			tree.DocumentId = item.DocumentId
			tree.ParentId = item.ParentId
			tree.DocumentName = item.DocumentName
			if item.ModifyAt > 0 {
				tree.Checked = true
			} else {
				tree.Checked = false
			}
			if item.ParentId == 0 {
				tree.Open = true
			} else {
				tree.Open = false
			}
			trees[index] = tree
		}
		data, err := json.Marshal(trees)
		if err != nil {
			return doclinks, "", err
		}
		doclist = string(data)
	}
	return doclinks, doclist, nil
}


func (m *LinkDocument) FixedBookDocuments(book_id int) (err error) {
	o := orm.NewOrm()

	sql := `UPDATE md_documents SET identify = UUID() WHERE book_id = ? AND identify = '';`
	o.Raw(sql, book_id).Exec()

	return nil
}


//分页查询指定项目的链接
func (m *LinkDocument) FindToLinksPager(pageIndex, pageSize, booKId int) (books []*BookResult, totalCount int, err error) {

	o := orm.NewOrm()

	sql1 := `SELECT count(*) AS total_count FROM md_books WHERE link_book = ?`

	err = o.Raw(sql1, booKId).QueryRow(&totalCount)

	if err != nil {
		return
	}

	offset := (pageIndex - 1) * pageSize
	sql2 := `SELECT
          book.*,
		  (select m.account from md_members as m where m.member_id = 
           (select rel1.member_id from md_relationship AS rel1 where book.book_id = rel1.book_id AND rel1.role_id = 0)
          ) as create_name
        FROM md_books AS book
        WHERE book.link_book = ?  
        ORDER BY book.order_index, book.book_id DESC limit ?,?`
	_, err = o.Raw(sql2, booKId, offset, pageSize).QueryRows(&books)
	if err != nil {
		logs.Error("分页查询项目列表 => ", err)
		return
	}
	sql := "SELECT m.account,doc.modify_time FROM md_documents AS doc LEFT JOIN md_members AS m ON doc.modify_at=m.member_id WHERE book_id = ? LIMIT 1 ORDER BY doc.modify_time DESC"

	if err == nil && len(books) > 0 {
		for index, book := range books {
			var text struct {
				Account    string
				ModifyTime time.Time
			}

			err1 := o.Raw(sql, book.BookId).QueryRow(&text)
			if err1 == nil {
				books[index].LastModifyText = text.Account + " 于 " + text.ModifyTime.Format("2006-01-02 15:04:05")
			}
		}
	}
	return
}

