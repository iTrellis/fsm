/*
Copyright © 2017 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package config

// DeepCopy 深度拷贝
func DeepCopy(value interface{}) interface{} {
	switch valueType := value.(type) {
	case map[string]interface{}, Options:
		newMap := make(map[string]interface{})
		valueMap := valueType.(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	case map[interface{}]interface{}:
		newMap := make(map[interface{}]interface{})
		for k, v := range valueType {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, len(valueType))
		for k, v := range valueType {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}

	return value
}
