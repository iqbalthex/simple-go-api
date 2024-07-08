package helper

import (
  "reflect"
)

func HideField(fieldName string, obj *any) any {
  tag := reflect.TypeOf(obj)

  field, found := tag.FieldByName(fieldName)

  if found {
    if field.Tag.Get("hide") {
      // delete obj.Id
      // delete obj[fieldName]

      obj[fieldName] = nil
    }
  }

  return obj
}
