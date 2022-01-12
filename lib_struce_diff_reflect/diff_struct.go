package lib_struce_diff_reflect

import (
	"fmt"
	"reflect"
	"strings"
)

type ReconvertFieldMeta struct {
	CN        string
	FieldTag  string
	FieldType string // 换成 reflect.Kind ?
	IsFmt     string
}

type RecursionFieldMeta struct {
	StructTag string
	PkField   string
	PKType    string // pk mulit[field_1.field_2]
}

type DiffStruct struct {
	DiffTag   string // 默认json
	Depth     int
	StructIdx int // 列表元素中的 索引
	StructTag string
	//PkField    string // 结构体主健，根据索引还是根据 字段名？
	ActionTag  string // TODO: 记录最外层的操作类型？
	DiffFields []*DiffField
	Children   []*DiffStruct
}

type DiffField struct {
	FieldNum  int // 字段索引
	FieldTag  string
	After     interface{}
	Before    interface{}
	ActionTag string
}

func NewDiffStruct() DiffStruct {
	return DiffStruct{
		DiffTag:    "json",
		Depth:      1,
		StructTag:  "scene",
		DiffFields: []*DiffField{},
		Children:   []*DiffStruct{},
	}

}

// Diff 根据结构体 字段定义的顺序遍历 入参tag: 为结构体的名称
//  typeof 参数
func (d *DiffStruct) Diff(v1, v2 interface{}, diffFieldMap map[string]map[string]ReconvertFieldMeta, recursionTagMap map[string]RecursionFieldMeta) {
	v1Typeof := reflect.TypeOf(v1)
	//v2Typeof := reflect.TypeOf(v2)
	v1ValueOf := reflect.ValueOf(v1)
	v2ValueOf := reflect.ValueOf(v2)

	if v1ValueOf.Kind() != reflect.Struct || v2ValueOf.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v1Typeof.NumField(); i++ {
		// TODO: 根据tag判断是否递归，暂时不递归处理 scene chain api 之外的
		fmt.Println("tag: ", d.StructTag+"."+v1Typeof.Field(i).Tag.Get(d.DiffTag))
		recursionField, recursionOK := recursionTagMap[d.StructTag+"."+v1Typeof.Field(i).Tag.Get(d.DiffTag)] // TODO: 耦合？
		if v1ValueOf.Field(i).Kind() == reflect.Slice && recursionOK && recursionField.PKType == "pk" {
			if v1ValueOf.Field(i).Len() == 0 {
				continue
			}
			for j := 0; j < v1ValueOf.Field(i).Len(); j++ {
				if v1ValueOf.Field(i).Index(j).Kind() != reflect.Struct {
					continue
				}

				// TODO: 如何判断前后列表中对应数据处于不同 列表位置
				// after: [{id:1,name:11}, {id:2,name:22}] before:  [{id:2,name:222}, {id:1,name:111}]
				children := &DiffStruct{
					DiffTag:    d.DiffTag,
					Depth:      d.Depth + 1,
					StructIdx:  j,
					StructTag:  d.StructTag + "." + v1Typeof.Field(i).Tag.Get(d.DiffTag),
					ActionTag:  "",
					DiffFields: []*DiffField{},
					Children:   []*DiffStruct{},
				}
				d.Children = append(d.Children, children)
				beforeV2, action := d.getBeforeV2(v1ValueOf.Field(i).Index(j).Interface(), v2ValueOf.Field(i).Interface(), recursionField.PkField)
				children.ActionTag = action
				if children.ActionTag != "update" {
					children.DiffFields = append(children.DiffFields, &DiffField{
						FieldNum:  i,
						FieldTag:  "",
						After:     nil,
						Before:    nil,
						ActionTag: children.ActionTag,
					})
					continue
				}
				children.Diff(v1ValueOf.Field(i).Index(j).Interface(), beforeV2, diffFieldMap, recursionTagMap)
			}
			continue
		} else if v1ValueOf.Field(i).Kind() == reflect.Slice && recursionOK && recursionField.PKType == "multi" {
			// TODO: 结构体多个字段对比
			diffFields := d.diffMuilFieldToStruct2(v1ValueOf.Field(i).Interface(), v2ValueOf.Field(i).Interface(), i, recursionField.PkField, v1Typeof.Field(i).Tag.Get(d.DiffTag))
			d.DiffFields = append(d.DiffFields, diffFields...)
			continue
		} else if v1ValueOf.Field(i).Kind() == reflect.Struct {
			d.Diff(v1ValueOf.Field(i).Interface(), v2ValueOf.Field(i).Interface(), diffFieldMap, recursionTagMap)
			continue
		}

		// 过滤字段
		if _, ok := diffFieldMap[d.StructTag][v1Typeof.Field(i).Tag.Get(d.DiffTag)]; !ok {
			continue
		}

		// TODO: map slice 对比； 接口的header
		if d.ActionTag == "create" {
			_diffField := &DiffField{
				FieldNum:  i,
				FieldTag:  d.StructTag + "." + v1Typeof.Field(i).Tag.Get(d.DiffTag),
				After:     v1ValueOf.Field(i).Interface(),
				Before:    nil,
				ActionTag: d.ActionTag,
			}
			d.DiffFields = append(d.DiffFields, _diffField)
			continue
		}

		if !reflect.DeepEqual(v1ValueOf.Field(i).Interface(), v2ValueOf.Field(i).Interface()) {
			_diffField := &DiffField{
				FieldNum:  i,
				FieldTag:  d.StructTag + "." + v1Typeof.Field(i).Tag.Get(d.DiffTag),
				After:     v1ValueOf.Field(i).Interface(),
				Before:    v2ValueOf.Field(i).Interface(),
				ActionTag: d.ActionTag,
			}
			d.DiffFields = append(d.DiffFields, _diffField)
		}

	}

}

func (d *DiffStruct) getBeforeV2(v1, v2 interface{}, PkField string) (interface{}, string) {
	if v2 == nil {
		return nil, "create"
	}
	// TODO: 区分 创建、更新、删除

	v1ValueOf := reflect.ValueOf(v1) // v1ValueOf具体值
	v2ValueOf := reflect.ValueOf(v2) // 列表
	if v2ValueOf.Len() == 0 {
		return nil, "create"
	}
	if v2ValueOf.Kind() != reflect.Slice {
		return nil, ""
	}

	for i := 0; i < v2ValueOf.Len(); i++ {
		if v2ValueOf.Index(i).Kind() == reflect.Slice {
			continue
		}
		if reflect.DeepEqual(v1ValueOf.FieldByName(PkField).Interface(), v2ValueOf.Index(i).FieldByName(PkField).Interface()) {
			return v2ValueOf.Index(i).Interface(), "update"
		}
	}

	return nil, "create"
}

func (d *DiffStruct) diffMuilFieldToStruct(v1, v2 interface{}, PkField string, FieldTag string) []*DiffField {
	var (
		diffFields []*DiffField
		pkFields   []string
	)

	pkFields = strings.Split(PkField, ".")
	v1ValueOf := reflect.ValueOf(v1)
	v2ValueOf := reflect.ValueOf(v2)

	for i := 0; i < v1ValueOf.Len(); i++ {
		var (
			v1PkValues []interface{}
		)
		if v1ValueOf.Index(i).Kind() != reflect.Struct {
			continue
		}

		for fi, field := range pkFields {
			v1PkValue := v1ValueOf.Index(i).FieldByName(field)

			if fi == 0 && v1PkValue.Interface() == nil && v1PkValue.IsNil() {
				break
			}
			v1PkValues = append(v1PkValues, v1PkValue.Interface())
		}
		if len(v1PkValues) == 0 {
			continue
		}

		diffField := &DiffField{
			FieldNum:  i,
			FieldTag:  FieldTag,
			After:     nil,
			Before:    nil,
			ActionTag: "update",
		}
		if v2ValueOf.Len() == 0 {
			diffField.ActionTag = "create"
			diffField.After = v1PkValues

			if len(diffField.After.([]interface{})) != 0 {
				diffFields = append(diffFields, diffField)
			}
			continue
		}

		// TODO: 不如转换成map进行对比？
		// 必须遍历找到对应的切片元素？
		// 循环上次操作记录的数据列表，找出第一个字段值相等的再对比后面的字段
		var v1PkFieldExists bool
		for j := 0; j < v2ValueOf.Len(); j++ {
			var v2PkValues []interface{}
			var unEqual = true
			for fj, field := range pkFields {
				v2PkValue := v2ValueOf.Index(j).FieldByName(field).Interface()
				v2PkValues = append(v2PkValues, v2PkValue)
				// 第一个字段值不相等 直接退出当前循环
				if fj == 0 && !reflect.DeepEqual(v1PkValues[fj], v2PkValue) {
					break
				}
				v1PkFieldExists = true
				// 从第二个字段开始比较值是否相等，如果值是空的怎么比较？
				if !reflect.DeepEqual(v1PkValues[fj], v2PkValue) {
					unEqual = false
				}
			}
			if !unEqual {
				diffField.After = v1PkValues
				diffField.Before = v2PkValues
				break
			}
		}

		// 之前的操作记录中没有这个数据
		if !v1PkFieldExists {
			diffField.After = v1PkValues
			diffField.ActionTag = "create"
		}

		// TODO: 删除处理的记录？
		if diffField.After == nil {
			continue
		}

		diffFields = append(diffFields, diffField)
	}

	return diffFields

}

func (d *DiffStruct) diffMuilFieldToStruct2(v1, v2 interface{}, fieldNum int, PkField string, FieldTag string) []*DiffField {
	var (
		v1PkValueArray [][]interface{}
		v2PkValueArray [][]interface{}

		diffFields []*DiffField
		pkFields   []string
	)

	pkFields = strings.Split(PkField, ".")
	v1ValueOf := reflect.ValueOf(v1)
	v2ValueOf := reflect.ValueOf(v2)

	vFunc := func(vValueof reflect.Value) [][]interface{} {
		var (
			vPkValueArray [][]interface{}
		)
		for i := 0; i < vValueof.Len(); i++ {
			var (
				vPkValues []interface{}
			)
			if vValueof.Index(i).Kind() != reflect.Struct {
				continue
			}
			for fi, field := range pkFields {
				vPkValue := vValueof.Index(i).FieldByName(field)
				if fi == 0 && vPkValue.Interface() == nil && vPkValue.IsNil() {
					break
				}
				vPkValues = append(vPkValues, vPkValue.Interface())
			}
			vPkValueArray = append(vPkValueArray, vPkValues)
		}
		return vPkValueArray
	}
	v1PkValueArray = vFunc(v1ValueOf)
	v2PkValueArray = vFunc(v2ValueOf)

	// 判断是否存在更新
	// TODO: 这里默认了第一个是主健，然后对比主健后面的其它字段
	for _, v1PkValues := range v1PkValueArray {
		diffField := &DiffField{
			FieldNum:  fieldNum,
			FieldTag:  FieldTag,
			After:     v1PkValues,
			Before:    nil,
			ActionTag: "update",
		}
		if len(v2PkValueArray) == 0 {
			diffField.ActionTag = "create"
			if len(diffField.After.([]interface{})) != 0 {
				diffFields = append(diffFields, diffField)
			}
			continue
		}

		// TODO: 这里处理记录每个主健后的字段，如果存在不一样的话
		var v1PkFieldExists bool
		for _, v2PkValues := range v2PkValueArray {
			// // 从第一个字段不相等就跳过
			if !reflect.DeepEqual(v1PkValues[0], v2PkValues[0]) {
				continue
			}
			v1PkFieldExists = true

			// 从第二个字段开始比较值是否相等，如果值是空的怎么比较？
			if !reflect.DeepEqual(v1PkValues, v2PkValues) {
				diffField.Before = v2PkValues
				diffFields = append(diffFields, diffField)
				break
			}
		}

		if !v1PkFieldExists {
			diffField.ActionTag = "create"
			diffFields = append(diffFields, diffField)
		}
	}

	// 判断是否存在删除
	for _, v2PkValues := range v2PkValueArray {
		diffField := &DiffField{
			FieldNum:  fieldNum,
			FieldTag:  FieldTag,
			After:     nil, // 删除删除
			Before:    v2PkValues,
			ActionTag: "delete",
		}
		if len(v1PkValueArray) == 0 {
			if len(diffField.Before.([]interface{})) != 0 {
				diffFields = append(diffFields, diffField)
			}
			continue
		}
		var v2PkFieldExists bool
		for _, v1PkValues := range v1PkValueArray {
			if !reflect.DeepEqual(v1PkValues[0], v2PkValues[0]) {
				continue
			}
			v2PkFieldExists = true
		}
		if !v2PkFieldExists {
			diffFields = append(diffFields, diffField)
		}
	}

	return diffFields

}

func (d *DiffStruct) diffMap(v1, v2 interface{}, PkField string, FieldTag string) []*DiffField {
	var (
		diffFields []*DiffField
		pkFields   []string
	)

	pkFields = strings.Split(PkField, ".")
	v1ValueOf := reflect.ValueOf(v1)
	v2ValueOf := reflect.ValueOf(v2)

	for i := 0; i < v1ValueOf.Len(); i++ {
		var (
			v1PkValues []interface{}
		)
		if v1ValueOf.Index(i).Kind() != reflect.Struct {
			continue
		}

		for fi, field := range pkFields {
			v1PkValue := v1ValueOf.Index(i).FieldByName(field)
			if fi == 0 && v1PkValue.IsZero() && v1PkValue.IsNil() {
				break
			}
			v1PkValues = append(v1PkValues, v1PkValue.Interface())
		}
		if len(v1PkValues) == 0 {
			continue
		}

		diffField := &DiffField{
			FieldNum:  i,
			FieldTag:  FieldTag,
			After:     nil,
			Before:    nil,
			ActionTag: "update",
		}
		if v2ValueOf.Len() == 0 {
			diffField.ActionTag = "create"
			diffField.After = v1PkValues

			if len(diffField.After.([]interface{})) != 0 {
				diffFields = append(diffFields, diffField)
			}
			continue
		}

		// TODO: 不如转换成map进行对比？
		// 必须遍历找到对应的切片元素？
		// 循环上次操作记录的数据列表，找出第一个字段值相等的再对比后面的字段
		var v1PkFieldExists bool
		for j := 0; j < v2ValueOf.Len(); j++ {
			var v2PkValues []interface{}
			var unEqual = true
			for fj, field := range pkFields {
				v2PkValue := v2ValueOf.Index(j).FieldByName(field).Interface()
				v2PkValues = append(v2PkValues, v2PkValue)
				// 第一个字段值不相等 直接退出当前循环
				if fj == 0 && !reflect.DeepEqual(v1PkValues[fj], v2PkValue) {
					break
				}
				v1PkFieldExists = true
				// 从第二个字段开始比较值是否相等，如果值是空的怎么比较？
				if !reflect.DeepEqual(v1PkValues[fj], v2PkValue) {
					unEqual = false
				}
			}
			if !unEqual {
				diffField.After = v1PkValues
				diffField.Before = v2PkValues
				break
			}
		}

		// 之前的操作记录中没有这个数据
		if !v1PkFieldExists {
			diffField.After = v1PkValues
			diffField.ActionTag = "create"
		}

		// TODO: 删除处理的记录？
		if diffField.After == nil {
			continue
		}

		diffFields = append(diffFields, diffField)
	}

	return diffFields
}

// ToCN 转中文
func (d *DiffStruct) ToCN() {

}
