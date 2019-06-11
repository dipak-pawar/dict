package dict

import (
	"reflect"
	"fmt"
)

const (
	NestingLimit = 50
)

func merge(p, q map[string]interface{}) map[string]interface{} {
	return mergeDict(p, q, 0)
}

func mergeDict(p, q map[string]interface{}, limit int) map[string]interface{} {
	if limit > NestingLimit {
		panic(fmt.Sprintf("not supporting this limit at this time, nesting level should be <= %d", NestingLimit))
	}
	for k, v := range p {
		if qv, ok := q[k]; ok {
			qMap, qMapOk := mapify(qv)
			pMap, pMapOk := mapify(v)
			if qMapOk && pMapOk {
				v = mergeDict(pMap, qMap, limit+1)
			} else {
				continue
			}
		}
		q[k] = v

	}
	return q
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
