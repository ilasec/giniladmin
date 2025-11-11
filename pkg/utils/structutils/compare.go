package structutils

import (
	"reflect"
)

//func CompareStructs(existingModel interface{}, updatedModel interface{}, ignoredFields []string) (map[string]interface{}, error) {
//	updates := make(map[string]interface{})
//
//	existingValue := reflect.ValueOf(existingModel).Elem()
//	updatedValue := reflect.ValueOf(updatedModel).Elem()
//
//	for i := 0; i < existingValue.NumField(); i++ {
//		fieldName := existingValue.Type().Field(i).Name
//
//		// 检查是否是忽略字段
//		if contains(ignoredFields, fieldName) {
//			continue
//		}
//
//		existingField := existingValue.Field(i)
//		updatedField := updatedValue.Field(i)
//
//		// 检查 updatedField 是否为零值，并判断是否需要强制更新零值字段
//		isUpdatedFieldZero := isZero(updatedField)
//		if !isUpdatedFieldZero || shouldUpdateZeroValueField(fieldName) {
//			existingFieldValue := existingField.Interface()
//			updatedFieldValue := updatedField.Interface()
//
//			if !reflect.DeepEqual(existingFieldValue, updatedFieldValue) {
//				// 使用字段名作为 key，无需转换为 snake_case
//				updates[fieldName] = updatedFieldValue
//			}
//		}
//	}
//
//	return updates, nil
//}

func CompareStructs(existingModel interface{}, updatedModel interface{}, ignoredFields []string, zeroCheck []string) (map[string]interface{}, error) {
	updates := make(map[string]interface{})

	existingValue := reflect.ValueOf(existingModel).Elem()
	updatedValue := reflect.ValueOf(updatedModel).Elem()

	for i := 0; i < existingValue.NumField(); i++ {
		fieldName := existingValue.Type().Field(i).Name

		// 检查是否是忽略字段
		if contains(ignoredFields, fieldName) {
			continue
		}

		existingField := existingValue.Field(i)
		updatedField := updatedValue.Field(i)

		// 检查 updatedField 是否为零值，并判断是否需要强制更新零值字段
		isUpdatedFieldZero := isZero(updatedField)
		if !isUpdatedFieldZero || shouldUpdateZeroValueField(fieldName, zeroCheck) {
			existingFieldValue := existingField.Interface()
			updatedFieldValue := updatedField.Interface()

			if !reflect.DeepEqual(existingFieldValue, updatedFieldValue) {
				// 记录变更的字段
				updates[fieldName] = updatedFieldValue

				// 更新 existingModel 的字段值
				if existingField.CanSet() {
					existingField.Set(updatedField)
				}
			}
		}
	}

	return updates, nil
}

// contains 检查字符串切片中是否包含指定字符串
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// isZero 和 shouldUpdateZeroValueField 函数与之前相同，这里省略...

// shouldUpdateZeroValueField 判断是否需要强制更新零值字段
func shouldUpdateZeroValueField(fieldName string, check []string) bool {
	for _, c := range check {
		if c == fieldName {
			return true
		}
	}
	return false
}

// isZero 判断一个 Value 是否为零值
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).Name == "Model" {
				continue // 跳过 gorm.Model
			}
			z = z && isZero(v.Field(i))
		}
		return z
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return isZero(reflect.Indirect(v))
	case reflect.Interface:
		if v.IsNil() {
			return true
		}
		return isZero(v.Elem())
	}
	// Compare other types directly.
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
