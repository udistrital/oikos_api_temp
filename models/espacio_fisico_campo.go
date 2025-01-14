package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EspacioFisicoCampo struct {
	Id                int            `orm:"column(id);pk;auto"`
	Valor             string         `orm:"column(valor)"`
	EspacioFisico     *EspacioFisico `orm:"column(espacio_fisico);rel(fk)"`
	Campo             *Campo         `orm:"column(campo);rel(fk)"`
	FechaInicio       time.Time      `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin          time.Time      `orm:"column(fecha_fin);type(timestamp without time zone)"`
	Activo            bool           `orm:"column(activo)"`
	FechaCreacion     time.Time      `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time      `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

type EspacioFisicoCampoV2 struct {
	Id                int              `orm:"column(id);pk;auto"`
	Valor             string           `orm:"column(valor)"`
	EspacioFisicoId   *EspacioFisicoV2 `orm:"column(espacio_fisico_id);rel(fk)"`
	CampoId           *CampoV2         `orm:"column(campo_id);rel(fk)"`
	Activo            bool             `orm:"column(activo)"`
	FechaInicio       time.Time        `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin          time.Time        `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	FechaCreacion     time.Time        `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time        `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *EspacioFisicoCampoV2) TableName() string {
	return "espacio_fisico_campo"
}

func init() {
	orm.RegisterModel(new(EspacioFisicoCampoV2))
}

// AddEspacioFisicoCampo insert a new EspacioFisicoCampo into database and returns
// last inserted Id on success.
func AddEspacioFisicoCampo(m *EspacioFisicoCampoV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEspacioFisicoCampoById retrieves EspacioFisicoCampo by Id. Returns error if
// Id doesn't exist
func GetEspacioFisicoCampoById(id int) (v *EspacioFisicoCampoV2, err error) {
	o := orm.NewOrm()
	v = &EspacioFisicoCampoV2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEspacioFisicoCampo retrieves all EspacioFisicoCampo matches certain condition. Returns empty list if
// no records exist
func GetAllEspacioFisicoCampo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EspacioFisicoCampoV2)).RelatedSel(5)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EspacioFisicoCampoV2
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEspacioFisicoCampo updates EspacioFisicoCampo by Id and returns error if
// the record to be updated doesn't exist
func UpdateEspacioFisicoCampoById(m *EspacioFisicoCampoV2) (err error) {
	o := orm.NewOrm()
	v := EspacioFisicoCampoV2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEspacioFisicoCampo deletes EspacioFisicoCampo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEspacioFisicoCampo(id int) (err error) {
	o := orm.NewOrm()
	v := EspacioFisicoCampoV2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EspacioFisicoCampoV2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
