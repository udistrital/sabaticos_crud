package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EstadoSabatico struct {
	Id                int    `orm:"column(id);pk;auto"`
	CodigoAbreviacion string `orm:"column(codigo_abreviacion)"`
	NombreEstado      string `orm:"column(nombre_estado);size(50)"`
	Activo            bool   `orm:"column(activo)"`
	FechaCreacion     string `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *EstadoSabatico) TableName() string {
	return "estado_sabatico"
}

func init() {
	orm.RegisterModel(new(EstadoSabatico))
}

// AddEstadoSabatico insert a new EstadoSabatico into database and returns
// last inserted Id on success.
func AddEstadoSabatico(m *EstadoSabatico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEstadoSabaticoById retrieves EstadoSabatico by Id. Returns error if
// Id doesn't exist
func GetEstadoSabaticoById(id int) (v *EstadoSabatico, err error) {
	o := orm.NewOrm()
	v = &EstadoSabatico{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEstadoSabatico retrieves all EstadoSabatico matches certain condition. Returns empty list if
// no records exist
func GetAllEstadoSabatico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EstadoSabatico)).RelatedSel()

	qs = applyEstadoSabaticoFilters(qs, query)

	sortFields, err := buildEstadoSabaticoSortFields(sortby, order)
	if err != nil {
		return nil, err
	}

	qs = qs.OrderBy(sortFields...)

	var l []EstadoSabatico
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}

	for _, v := range l {
		m := make(map[string]interface{})
		val := reflect.ValueOf(v)
		for _, fname := range fields {
			m[fname] = val.FieldByName(fname).Interface()
		}
		ml = append(ml, m)
	}

	return ml, nil
}

func applyEstadoSabaticoFilters(qs orm.QuerySeter, query map[string]string) orm.QuerySeter {
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}
	return qs
}

func buildEstadoSabaticoSortFields(sortby []string, order []string) ([]string, error) {
	var sortFields []string

	if len(sortby) == 0 {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
		return sortFields, nil
	}

	if len(sortby) == len(order) {
		for i, v := range sortby {
			orderby, err := buildEstadoSabaticoOrder(v, order[i])
			if err != nil {
				return nil, err
			}
			sortFields = append(sortFields, orderby)
		}
		return sortFields, nil
	}

	if len(order) == 1 {
		for _, v := range sortby {
			orderby, err := buildEstadoSabaticoOrder(v, order[0])
			if err != nil {
				return nil, err
			}
			sortFields = append(sortFields, orderby)
		}
		return sortFields, nil
	}

	return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
}

func buildEstadoSabaticoOrder(field string, direction string) (string, error) {
	if direction == "desc" {
		return "-" + field, nil
	}
	if direction == "asc" {
		return field, nil
	}
	return "", errors.New("Error: Invalid order. Must be either [asc|desc]")
}

// UpdateEstadoSabatico updates EstadoSabatico by Id and returns error if
// the record to be updated doesn't exist
func UpdateEstadoSabaticoById(m *EstadoSabatico) (err error) {
	o := orm.NewOrm()
	v := EstadoSabatico{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEstadoSabatico deletes EstadoSabatico by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEstadoSabatico(id int) (err error) {
	o := orm.NewOrm()
	v := EstadoSabatico{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EstadoSabatico{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
