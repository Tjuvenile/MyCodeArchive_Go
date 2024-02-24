package base_tool

import (
	"MyCodeArchive_Go/utils/logging"
	"hash/fnv"
	"reflect"
	"sync"
	"unsafe"
)

const LinkLockNum = 100

var Lock [LinkLockNum]sync.Mutex

func GetHashLock(linkId string) {
	data := []byte(linkId)
	index := strToHash64(data) % LinkLockNum
	Lock[index].Lock()
}

func ReleaseHashLock(linkId string) {
	data := []byte(linkId)
	index := strToHash64(data) % LinkLockNum
	if IsMutexLocked(&Lock[index]) {
		Lock[index].Unlock()
	} else {
		logging.Log.Warnf("Lock failed, the object is unlock. %s", linkId)
	}
}

func IsMutexLocked(m *sync.Mutex) bool {
	// 通过反射获取互斥锁的私有字段 state 的地址
	stateAddr := (*int32)(unsafe.Pointer(
		reflect.ValueOf(m).Elem().FieldByName("state").UnsafeAddr(),
	))

	// 读取 state 的值，如果值为 0 表示未被锁定，否则表示已被锁定
	return *stateAddr != 0
}

func strToHash64(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}
